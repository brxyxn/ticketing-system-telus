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
	"github.com/brxyxn/ticketing-system-telus/backend/app/database"
	"github.com/brxyxn/ticketing-system-telus/backend/app/routes"
	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
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
	vars, err := config.Configure() // Configuring the app variables
	if err != nil {
		u.Log.Error("Environment variables weren't loaded correctly!", err)
		return
	}

	a.BindAddr = ":" + vars.Port

	d := database.NewDatabaseHandler()
	// Sql
	a.DB = d.InitializePostgresql(
		vars.Sql.Host,
		vars.Sql.Port,
		vars.Sql.User,
		vars.Sql.Password,
		vars.Sql.Name,
		vars.Sql.Sslmode,
	)

	c := database.NewCacheHandler()
	a.Cache = c.InitializeRedis(
		vars.Cache.Host+":"+vars.Cache.Port,
		vars.Cache.Password,
		vars.Cache.Name,
	)

	app := fiber.New()

	store := session.New()
	// use cors with fiber v2
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// app.Post("/login", Authenticate)

	// api := app.Group("/api/v1")

	// api.Post("/login", Authenticate)
	// JWT Middleware
	// api.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("secret"),
	// }))
	// api.Get("/restricted", restricted)

	// api := app.Group("/api/v1", middleware.Authenticate)
	// api.Get("/", restricted)

	// Frontend
	routes.ReactRoutes(app)
	// API
	routes.AccountRoutes(app, a.DB, a.Cache)
	routes.CustomerRoutes(app, a.DB)
	routes.AgentRoutes(app, a.DB)
	routes.TicketRoutes(app, a.DB)

	// middleware.Authenticate(&fiber.Ctx{})

	app.Listen(a.BindAddr)
}

// type Login struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func Authenticate(c *fiber.Ctx) error {
// 	var login Login
// 	c.BodyParser(&login)

// 	if login.Email == "" || login.Password == "" {
// 		return c.SendStatus(fiber.StatusBadGateway)
// 	}

// 	// Create the Claims
// 	claims := jwt.MapClaims{
// 		"email": login.Email,
// 		"admin": true,
// 		"exp":   time.Now().Add(time.Hour * 1).Unix(),
// 	}

// 	// Create token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Generate encoded token and send it as response.
// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	return c.JSON(fiber.Map{"token": t})
// }

// func restricted(c *fiber.Ctx) error {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	name := claims["email"].(string)
// 	return c.SendString("Welcome " + name)
// }

/*
Runs the new server.
*/
func (a *App) Run() {
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
