package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const dsn = "root:123456@tcp(localhost:4000)/gin_demo?charset=utf8mb4&parseTime=True&loc=Local"

func initializeDatabase() (err error) {

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return nil
}

func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}
	fmt.Println("connect to database success")
}
