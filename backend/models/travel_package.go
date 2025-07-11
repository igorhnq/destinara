package models

import (
	"time"

	"gorm.io/gorm"
)

type TravelPackage struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"type:text"`
	Destination string         `json:"destination" gorm:"not null"`
	Price       float64        `json:"price" gorm:"not null"`
	Duration    int            `json:"duration" gorm:"not null"` // duração em dias
	ImageURL    string         `json:"image_url"`
	Available   bool           `json:"available" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
} 