package migrations

import (
	"gorm.io/gorm"
	"upvote-crypto/database/models"
)

func RunMigrations(db *gorm.DB) {
	db.Debug().AutoMigrate(models.Currency{})
}