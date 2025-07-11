package main

import (
	"log"

	"destinara-backend/config"
	"destinara-backend/database"
	"destinara-backend/models"
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

	// Criar pacotes de viagem iniciais
	packages := []models.TravelPackage{
		{
			Name:        "Aventura na Amazônia",
			Description: "Explore a selva amazônica com guias experientes. Inclui passeios de barco, observação de animais selvagens e hospedagem em lodge na floresta.",
			Destination: "Manaus, Amazonas",
			Price:       2500.00,
			Duration:    5,
			ImageURL:    "/images/manaus-card.svg",
			Available:   true,
		},
		{
			Name:        "Cultura e História de Salvador",
			Description: "Descubra a rica cultura afro-brasileira de Salvador. Visite o Pelourinho, experimente a culinária local e participe de apresentações de capoeira.",
			Destination: "Salvador, Bahia",
			Price:       1800.00,
			Duration:    4,
			ImageURL:    "/images/salvador-card.svg",
			Available:   true,
		},
		{
			Name:        "Romance em Gramado",
			Description: "O destino perfeito para casais. Conheça a charmosa cidade de Gramado, visite o Mini Mundo, prove chocolates artesanais e aproveite o clima europeu.",
			Destination: "Gramado, Rio Grande do Sul",
			Price:       2200.00,
			Duration:    3,
			ImageURL:    "/images/gramado-card.svg",
			Available:   true,
		},
		{
			Name:        "Maravilhas de Foz do Iguaçu",
			Description: "Conheça uma das Sete Maravilhas Naturais do Mundo. Visite as Cataratas do Iguaçu, o Parque das Aves e a Usina de Itaipu.",
			Destination: "Foz do Iguaçu, Paraná",
			Price:       1600.00,
			Duration:    3,
			ImageURL:    "/images/foz-do-iguacu-card.svg",
			Available:   true,
		},
	}

	// Inserir pacotes no banco
	for _, pkg := range packages {
		var existingPackage models.TravelPackage
		if err := database.DB.Where("name = ?", pkg.Name).First(&existingPackage).Error; err == nil {
			log.Printf("Pacote '%s' já existe, pulando...", pkg.Name)
			continue
		}

		if err := database.DB.Create(&pkg).Error; err != nil {
			log.Printf("Erro ao criar pacote '%s': %v", pkg.Name, err)
		} else {
			log.Printf("Pacote '%s' criado com sucesso", pkg.Name)
		}
	}

	log.Println("Seed concluído com sucesso!")
} 