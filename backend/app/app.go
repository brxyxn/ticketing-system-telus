package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/brxyxn/ticketing-system-telus/backend/app/config"
	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	p "github.com/brxyxn/ticketing-system-telus/backend/internal/datasource"
	"github.com/brxyxn/ticketing-system-telus/backend/internal/datasource/postgres"
	"github.com/brxyxn/ticketing-system-telus/backend/internal/users"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type App struct {
	Server   *http.Server
	Router   *fiber.App
	DB       *sql.DB
	Cache    *redis.Client
	Ctx      context.Context
	L        *log.Logger
	BindAddr string
	handler  http.Handler
}

func (a *App) Setup() {
	db := p.NewHandlers(a.DB, a.Cache)

	vars, err := config.Configure() // Configuring the app variables
	if err != nil {
		u.Log.Error("Environment variables weren't loaded correctly!", err)
		return
	}

	a.BindAddr = ":" + vars.Port

	// Sql
	db.InitializePostgresql(
		vars.Sql.Host,
		vars.Sql.Port,
		vars.Sql.User,
		vars.Sql.Password,
		vars.Sql.Name,
		vars.Sql.Sslmode,
	)

	// // Cache
	db.InitializeCache(
		vars.Cache.Host+":"+vars.Cache.Port,
		vars.Cache.Password,
		vars.Cache.Name,
	)
}

func (a *App) initRoutes() {
	app := fiber.New()

	// use cors with fiber v2
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	app.Static("/", "./build", fiber.Static{
		Index: "index.html",
	})

	api := app.Group("/api/v1")

	// Users
	userRepo := postgres.NewPostgresUserRepository(a.DB)
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserFiberHandler(userService)

	// Accounts
	api.Get("/register", userHandler.RegisterAccount) // register profile, user and [customer + company | agent + team]
	api.Get("/login", userHandler.RegisterAccount)    // login user
	api.Get("/user", userHandler.RegisterAccount)     // returns user profile
	api.Get("/logout", userHandler.RegisterAccount)   // logout user

	// customer routes
	api.Get("/customers", userHandler.RegisterAccount)
	api.Get("/customers/:user_id", userHandler.Authenticate)

	// agent routes
	api.Get("/teams/", userHandler.RegisterAccount)
	api.Get("/teams/:team_id", userHandler.RegisterAccount)
	api.Get("/tiers", userHandler.RegisterAccount)
	api.Get("/tiers/:tier_id", userHandler.RegisterAccount)

	// ticket routes
	api.Get("/tickets", userHandler.RegisterAccount)
	api.Get("/tickets/:ticket_id", userHandler.RegisterAccount)
	api.Get("/tickets/:ticket_id/comments", userHandler.RegisterAccount)

	// route for tracking of history of tickets
	api.Get("/tickets/:ticket_id/history", userHandler.RegisterAccount)

	app.Listen(a.BindAddr)
}

/*
Runs the new server.
*/
func (a *App) Run() {
	// Initializing routes
	a.initRoutes()

	// Creating a new server
	a.Server = &http.Server{
		Addr:         a.BindAddr,        // configure the bind address
		Handler:      a.handler,         // set the default handler
		ErrorLog:     a.L,               // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// Starting the server
	go func() {
		u.Log.Info("Running server on port", a.BindAddr)

		err := a.Server.ListenAndServe()
		if err != nil {
			u.Log.Info("Server Status: ", err)
			os.Exit(1)
		}
	}()

	// Creating channel
	cs := make(chan os.Signal, 1)

	signal.Notify(cs, os.Interrupt, os.Kill)
	// signal.Notify(sigchan, os.Kill) // If running on Windows

	sigchan := <-cs
	u.Log.Debug("Signal received:", sigchan)

	ctx, fn := context.WithTimeout(context.Background(), 30*time.Second)
	defer fn()
	a.Server.Shutdown(ctx)
}
