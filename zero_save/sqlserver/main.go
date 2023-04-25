package main

import (
	"codeflow/zero_save/sqlserver/logic"
	"codeflow/zero_save/sqlserver/middleware"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

/*
sqlserver gorm
*/

const (
	username = "gorm"
	password = "LoremIpsum86"
	hostname = "localhost"
	port     = 9930
	database = "gorm"
)

var (
	db *gorm.DB
)

func init2() {
	// github.com/denisenkom/go-mssqldb
	dsn := fmt.Sprintf(`sqlserver://%s:%s@%s:%d?database=%s`, username, password, hostname, port, database)
	// dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Sqlserver][init] sqlserver error: ", err)
	}

	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

type User struct {
	ID string `json:"id"`
}

func main() {
	router := gin.Default()
	// router.Use(Auth())

	go func() {
		logrus.Info("start prometheus")
		middleware.Start()
	}()

	gp := router.Group("app/v1")
	{
		gp.Use(middleware.Metric())
		gp.GET("/log", logic.Login)
	}

	router.Run("127.0.0.1:9000")
}
