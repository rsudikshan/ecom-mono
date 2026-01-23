package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb(env *Env) *gorm.DB{
	connStr := fmt.
	Sprintf("user=%s password=%s dbname=%s sslmode=disable host=localhost port=%s",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_SERVER_PORT"),
	)
	db,err := gorm.Open(postgres.Open(connStr),&gorm.Config{})

	if err!=nil {
		log.Fatal(err.Error())
	}

	sqlDB,err := db.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(1*time.Hour)

	if err!=nil {
		log.Fatal(err.Error())
	}

	return db
}