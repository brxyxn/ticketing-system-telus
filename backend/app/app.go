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
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
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

func (a *App) initRoutes() {
	app := fiber.New()

	app.Static("/", "./build", fiber.Static{
		Index: "index.html",
	})
	// app.Get("/api", handlerApi)

	app.Listen(":5000")
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
