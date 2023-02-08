package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:snmsung1.@tcp(localhost:3306)/writedream?parseTime=true")
	// defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Connect DB Success")
	Database = db
}
