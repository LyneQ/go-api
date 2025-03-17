package sql

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func Main() {

	config := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "fakedb",
	}

	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	fmt.Println("[SQL] Connected to database")
}
