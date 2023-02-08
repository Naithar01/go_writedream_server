package main

import (
	"github.com/Naithar01/go_write_dream/db"
	"github.com/Naithar01/go_write_dream/router"
)

func main() {
	app := router.InitRouter()

	db.ConnectDB()
	defer db.Database.Close()

	app.Run()
}
