package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	databaseUrl := os.Getenv("DATABASE_URL")

	//Using MySQL
	// const MYSQL = "indra:Indra19@tcp(127.0.0.1:3306)/gofiber?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := MYSQL
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//Using PostgreSQL
	// const POSTGRES = "postgresql://postgres:Indra19@localhost:5432/gofiber?sslmode=disable&TimeZone=Asia/Jakarta"
	// dsn := POSTGRES
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//Using PostgreSQL Heroku
	dsn := databaseUrl
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}

	fmt.Println("Connected to database")
}
