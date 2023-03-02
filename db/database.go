package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "p7b7agz83lbrmi3j:xazn26d1tr5mimqe@tcp(u3r5w4ayhxzdrw87.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/tchvl6tc2odrr8yv")
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
