package datsource

import (
	"database/sql"
	"fmt"

	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func (h *Handlers) InitializePostgresql(host, port, user, password, dbname, sslmode string) {
	u.Log.Debug("Initializing Postgresql")
	connectionStr := fmt.Sprintf(
		"host=%s port=%v user=%s "+
			"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	var err error
	h.db, err = sql.Open("pgx", connectionStr)
	if err != nil {
		u.Log.Error("Error opening a new connection to the DB.", err)
		return
	}

	err = h.db.Ping()
	if err != nil {
		h.db.Close()
		u.Log.Error(err)
		return
	}
	u.Log.Info("Connected to database", dbname, "with user", user, "at", host+":"+port)
}
