package dao

import (
	"codeflow/zero_save/sqlserver/model"
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

	// db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	// d.db.Get()
	return nil, nil
}

/*
temp
*/
var dbGorm *gorm.DB

func init() {
	dsn := fmt.Sprintf(`sqlserver://%s:%s@%s:%d?database=%s`, username, password, hostname, port, database)
	dbGorm, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Sqlserver][init] sqlserver error: ", err)
	}

	sqlDB, err := dbGorm.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

func CreateJob(job model.Job) (*model.Job, error) {
	dbGorm.Exec("")

	return nil, nil
}

func GetJob(jobID string) (*model.Job, error) {
	dbGorm.Exec("")

	return nil, nil
}

func ListJob() ([]*model.Job, error) {
	dbGorm.Exec("")

	return nil, nil
}

func UpdateJob() (*model.Job, error) {
	dbGorm.Exec("")

	return nil, nil
}

func DeleteJob() error {
	dbGorm.Exec("")

	return nil
}
