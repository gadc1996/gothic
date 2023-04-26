package mysql

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var database *sql.DB

func init() {
	config := getConfig()

	database = open(config)
	ping(database)
}

func getConfig() string {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "i7QQWfRmkNA7MBWJZC0wMPCaT",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}
	return config.FormatDSN()
}

func open(config string) *sql.DB {
	var error error
	database, error = sql.Open("mysql", config)
	if error != nil {
		log.Fatal(error)
	}
	return database
}

func ping(database *sql.DB) {
	error := database.Ping()
	if error != nil {
		log.Fatal(error)
	}
}
