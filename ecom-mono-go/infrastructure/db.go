package infrastructure

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	migrate "github.com/rubenv/sql-migrate"
)

func NewDb(env *Env) *gorm.DB{
	connStr := fmt.
	Sprintf("user=%s password=%s dbname=%s sslmode=disable host=localhost port=%s",
	env.DB_USER,
	env.DB_PASSWORD,
	env.DB_NAME,
	env.DB_SERVER_PORT,
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

func Migrate(db *gorm.DB,){
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/",
	}	

	sqlDB, err := db.DB()

	if err!=nil {
		log.Fatal(err.Error())
	}

	log.Default().Print("running migrations")

	_,err = migrate.Exec(sqlDB,"postgres", migrations, migrate.Up)

	if err!=nil {
		log.Fatal(err.Error())
	}

	log.Default().Print("migration successful")
}