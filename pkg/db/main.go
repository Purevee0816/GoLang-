package db

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

type GORMDB struct {
	Err error
	DB  *gorm.DB
}

var db *GORMDB

func try() {
	s := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ulaanbaatar", os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))


	db.DB, db.Err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  s,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	
	sqlDB, _ := db.DB.DB()
	val, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONN"))
	if err != nil {
		val = 10
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(val)

	val, err = strconv.Atoi(os.Getenv("MAX_OPEN_CONN"))
	if err != nil {
		val = 100
	}
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(val)

	val, err = strconv.Atoi(os.Getenv("MAX_CONN_LIFETIME"))
	if err != nil {
		val = 120
	}
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(val) * time.Second)
}

func Init() {
	db = new(GORMDB)
	try()
}

func Tx(ctx context.Context) (g *gorm.DB, err error) {
	sqlDB, err := db.DB.DB()

	if sqlDB.PingContext(ctx); err != nil {
		try()
		return db.DB, err
	}

	return db.DB.WithContext(ctx), nil
}