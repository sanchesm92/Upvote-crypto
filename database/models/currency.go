package models

import (
	"time"
	"gorm.io/gorm"
)

type Currency struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string  `json:"name"`
	Vote uint `json:"vote"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}