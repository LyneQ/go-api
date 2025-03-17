package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

var DB *sql.DB

func Main() {

	config := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "fakedb",
	}

	var err error
	DB, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	fmt.Println("[SQL] Connected to database")
}

func verifyConnectionState() (bool, error) {
	if DB == nil {
		return false, fmt.Errorf("[SQL] Database connection is not established")
	}
	return true, nil
}
