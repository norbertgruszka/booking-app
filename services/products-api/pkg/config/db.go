package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

type dbDetails struct {
	host     string
	user     string
	password string
	dbName   string
	port     string
}

func ConnectToDb() {
	dbDetails := &dbDetails{
		host:     os.Getenv("POSTGRES_HOST"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASS"),
		dbName:   os.Getenv("POSTGRES_DB_NAME"),
		port:     os.Getenv("POSTGRES_PORT"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbDetails.host, dbDetails.user, dbDetails.password, dbDetails.dbName, dbDetails.port)
	dd, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = dd
}
