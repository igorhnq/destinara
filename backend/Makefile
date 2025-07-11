.PHONY: help build run test clean deps seed docker-build docker-run

# Variáveis
BINARY_NAME=destinara-backend
MAIN_FILE=main.go

# Comandos principais
help: ## Mostra esta ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

deps: ## Instala as dependências
	go mod tidy
	go mod download

build: ## Compila a aplicação
	go build -o $(BINARY_NAME) $(MAIN_FILE)

run: ## Executa a aplicação em modo desenvolvimento
	go run $(MAIN_FILE)

dev: ## Executa com hot reload (requer air)
	air

test: ## Executa os testes
	go test ./...

test-coverage: ## Executa os testes com cobertura
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean: ## Remove arquivos compilados
	rm -f $(BINARY_NAME)
	rm -f coverage.out

seed: ## Popula o banco com dados iniciais
	go run scripts/seed.go

setup: ## Configura o ambiente inicial
	@echo "Configurando ambiente..."
	@if [ ! -f .env ]; then cp env.example .env; fi
	go mod tidy
	@echo "✅ Ambiente configurado! Configure o arquivo .env conforme necessário."

# Docker
docker-build: ## Constrói a imagem Docker
	docker build -t $(BINARY_NAME) .

docker-run: ## Executa o container Docker
	docker run -p 8080:8080 --env-file .env $(BINARY_NAME)

docker-compose-up: ## Inicia os serviços com Docker Compose
	docker-compose up -d

docker-compose-down: ## Para os serviços do Docker Compose
	docker-compose down

# Linting e formatação
fmt: ## Formata o código
	go fmt ./...

lint: ## Executa o linter
	golangci-lint run

# Banco de dados
db-migrate: ## Executa as migrações do banco
	go run $(MAIN_FILE)

# Utilitários
health: ## Verifica se a API está funcionando
	curl -f http://localhost:8080/health || echo "API não está respondendo"

docs: ## Gera documentação (se disponível)
	@echo "Documentação disponível em: docs/swagger.json"
	@echo "Visualize em: https://editor.swagger.io/"

# Desenvolvimento completo
dev-setup: deps setup ## Setup completo para desenvolvimento
	@echo "🎉 Ambiente de desenvolvimento configurado!"
	@echo "📋 Próximos passos:"
	@echo "1. Configure o arquivo .env"
	@echo "2. Certifique-se de que o PostgreSQL está rodando"
	@echo "3. Execute: make run"
	@echo "4. (Opcional) Execute: make seed" 