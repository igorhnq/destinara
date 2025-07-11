package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("erro ao conectar com banco de dados: %v", err)
	}

	DB = db

	// Auto-migrar tabelas
	if err := AutoMigrate(); err != nil {
		return fmt.Errorf("erro ao migrar tabelas: %v", err)
	}

	log.Println("Banco de dados conectado e migrado com sucesso")
	return nil
}

func AutoMigrate() error {
	return DB.AutoMigrate(
		&User{},
		&TravelPackage{},
		&Booking{},
	)
} 