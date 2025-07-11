package database

import "destinara-backend/models"

// Importar modelos para que sejam registrados no GORM
var (
	User          = models.User{}
	TravelPackage = models.TravelPackage{}
	Booking       = models.Booking{}
) 