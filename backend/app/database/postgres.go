package database

import (
	"database/sql"
	"fmt"

	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DatabaseHandler interface {
	InitializePostgresql(host, port, user, password, dbname, sslmode string) *sql.DB
}

type DatabaseHdl struct {
	db *sql.DB
}

func NewDatabaseHandler() DatabaseHandler {
	return &DatabaseHdl{}
}

func (d *DatabaseHdl) InitializePostgresql(host, port, user, password, dbname, sslmode string) *sql.DB {
	u.Log.Debug("Initializing Postgresql")
	connectionStr := fmt.Sprintf(
		"host=%s port=%v user=%s "+
			"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	var err error
	d.db, err = sql.Open("pgx", connectionStr)
	if err != nil {
		u.Log.Error("Error opening a new connection to the DB.", err)
		return &sql.DB{}
	}

	err = d.db.Ping()
	if err != nil {
		d.db.Close()
		u.Log.Error(err)
		return &sql.DB{}
	}
	u.Log.Info("Connected to database", dbname, "with user", user, "at", host+":"+port)
	return d.db
}
