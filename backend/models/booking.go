package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	UserID        uint           `json:"user_id" gorm:"not null"`
	User          User           `json:"user" gorm:"foreignKey:UserID"`
	TravelPackageID uint         `json:"travel_package_id" gorm:"not null"`
	TravelPackage TravelPackage  `json:"travel_package" gorm:"foreignKey:TravelPackageID"`
	StartDate     time.Time      `json:"start_date" gorm:"not null"`
	EndDate       time.Time      `json:"end_date" gorm:"not null"`
	TotalPrice    float64        `json:"total_price" gorm:"not null"`
	Status        string         `json:"status" gorm:"default:'pending'"` // pending, confirmed, cancelled
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
} 