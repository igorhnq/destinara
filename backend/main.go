package main

import (
	"log"
	"os"

	"destinara-backend/config"
	"destinara-backend/routes"
	"destinara-backend/database"
)

func main() {
	// Carregar variáveis de ambiente
	if err := config.LoadEnv(); err != nil {
		log.Fatal("Erro ao carregar variáveis de ambiente:", err)
	}

	// Inicializar banco de dados
	if err := database.InitDB(); err != nil {
		log.Fatal("Erro ao inicializar banco de dados:", err)
	}

	// Configurar e iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := routes.SetupRoutes()
	
	log.Printf("Servidor iniciado na porta %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
} 