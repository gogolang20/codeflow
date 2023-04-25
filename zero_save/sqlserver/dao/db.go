package dao

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

const (
	username = "gorm"
	password = "LoremIpsum86"
	hostname = "localhost"
	port     = 9930
	database = "gorm"
)

type IDB interface {
	Get() (interface{}, error)
	List() (interface{}, error)
}

type DB struct {
	db *gorm.DB
}

func NewDB() IDB {
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

	return &DB{
		db: db,
	}
}

func (d *DB) Get() (interface{}, error) {
	// d.db.Get()
	return nil, nil
}

func (d *DB) List() (interface{}, error) {

	return nil, nil
}
