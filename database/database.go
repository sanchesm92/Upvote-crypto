package database

import (
	"gorm.io/gorm"
	"time"
	"log"
	"gorm.io/driver/postgres"
	"upvote-crypto/database/migrations"
)

var db *gorm.DB

func StartDB() {
	db_config := "host=localhost port=5432 user=postgres dbname=upvote sslmode=disable password=postgres"
	database, err := gorm.Open(postgres.Open(db_config), &gorm.Config{})
	if err != nil {
		log.Fatal("error: ", err)
	}
	db = database

	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigrations(db)

}

func GetDatabase() *gorm.DB {
	return db
}