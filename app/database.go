package app

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/komblog/restful-api-go/helper"
)

func Connection() *sql.DB {
	connection, errConnection := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_restful_api")
	helper.PanicIfError(errConnection)

	connection.SetMaxIdleConns(5)
	connection.SetMaxOpenConns(10)
	connection.SetConnMaxLifetime(60 * time.Minute)
	connection.SetConnMaxIdleTime(10 * time.Minute)

	return connection
}
