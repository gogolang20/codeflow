package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root:123456@tcp(localhost:4000)/gin_demo?charset=utf8mb4&parseTime=True&loc=Local"

// "root:123456@tcp(localhost:3060)/atguigudb"

func main() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 尝试建立链接
	if err := db.Ping(); err != nil {
		fmt.Println("connect to database failed...")
		panic(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// 这才是真正的链接成功
	fmt.Println("connect to database success")
}
