package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// Carregar arquivo .env se existir
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// Verificar variáveis obrigatórias
	requiredVars := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "JWT_SECRET"}
	
	for _, varName := range requiredVars {
		if os.Getenv(varName) == "" {
			log.Printf("Variável de ambiente %s não definida", varName)
		}
	}

	return nil
} 