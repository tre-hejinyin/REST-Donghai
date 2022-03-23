package infra

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	timezone := os.Getenv("TZ")
	dsn := fmt.Sprintf("user=%s password=%s database=%s port=%s host=%s TimeZone=%s", user, password, dbname, port, host, timezone)
	return newPostgres(dsn)
}

func newPostgres(dsn string) (conn *gorm.DB, err error) {
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	db, err := conn.DB()
	if err != nil {
		return
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	db.SetConnMaxLifetime(time.Hour)
	return
}
