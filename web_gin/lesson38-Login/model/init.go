package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitializeDatabase() (err error) {
	dsn := "root:123456@tcp(localhost:4000)/login?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect 2 database failed, err:%v\n", err)
		return
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	return nil
}
