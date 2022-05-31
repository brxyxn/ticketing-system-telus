package main

import (
	a "github.com/brxyxn/ticketing-system-telus/backend/app"
	u "github.com/brxyxn/ticketing-system-telus/backend/app/utils"
)

func main() {
	a := a.App{}

	a.L = u.InitLogs("telus-api")

	a.Setup() // Setup database and cache

	a.Run() // Run the app
}
