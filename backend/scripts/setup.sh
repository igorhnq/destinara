#!/bin/bash

echo "🚀 Configurando o backend Destinara..."

# Verificar se Go está instalado
if ! command -v go &> /dev/null; then
    echo "❌ Go não está instalado. Por favor, instale o Go 1.21 ou superior."
    exit 1
fi

echo "✅ Go encontrado: $(go version)"

# Verificar se PostgreSQL está instalado
if ! command -v psql &> /dev/null; then
    echo "⚠️  PostgreSQL não encontrado. Você pode usar Docker para executar o banco de dados."
    echo "   Execute: docker-compose up postgres -d"
else
    echo "✅ PostgreSQL encontrado"
fi

# Instalar dependências
echo "📦 Instalando dependências..."
go mod tidy

# Criar arquivo .env se não existir
if [ ! -f .env ]; then
    echo "📝 Criando arquivo .env..."
    cp env.example .env
    echo "✅ Arquivo .env criado. Configure as variáveis de ambiente conforme necessário."
else
    echo "✅ Arquivo .env já existe"
fi

echo ""
echo "🎉 Setup concluído!"
echo ""
echo "📋 Próximos passos:"
echo "1. Configure as variáveis de ambiente no arquivo .env"
echo "2. Certifique-se de que o PostgreSQL está rodando"
echo "3. Execute: go run main.go"
echo "4. (Opcional) Execute: go run scripts/seed.go para popular o banco"
echo ""
echo "🐳 Ou use Docker: docker-compose up" 