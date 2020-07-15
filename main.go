package main

import (
	"github.com/kingsman20/fintech_app/api"
	"github.com/kingsman20/fintech_app/database"
)

func main() {
	// Do migration
	// migrations.MigrateTransactions()

	// Init database
	database.InitDatabase()
	api.StartApi()
}
