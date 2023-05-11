package db

import (
	"database/sql"
	"e-commerce/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USER + ":" + conf.DB_PASS + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString failed")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN invalid connection")
	}
}

func CreateCon() *sql.DB {
	return db
}
