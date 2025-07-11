# Destinara Backend

Backend da aplicação Destinara desenvolvido em Go com Gin framework e PostgreSQL.

## 🚀 Funcionalidades

- **Autenticação JWT**: Registro e login de usuários
- **Gestão de Pacotes de Viagem**: CRUD completo para pacotes turísticos
- **Sistema de Reservas**: Criação e gestão de reservas de viagem
- **API RESTful**: Endpoints bem estruturados e documentados
- **Banco de Dados PostgreSQL**: Com GORM para ORM
- **CORS Configurado**: Para comunicação com frontend React

## 📋 Pré-requisitos

- Go 1.21 ou superior
- PostgreSQL 12 ou superior
- Git

## 🛠️ Instalação

1. **Clone o repositório**
   ```bash
   git clone <url-do-repositorio>
   cd destinara/backend
   ```

2. **Instale as dependências**
   ```bash
   go mod tidy
   ```

3. **Configure o banco de dados**
   - Crie um banco PostgreSQL chamado `destinara_db`
   - Copie o arquivo `env.example` para `.env`
   - Configure as variáveis de ambiente no arquivo `.env`

4. **Execute as migrações**
   ```bash
   go run main.go
   ```

5. **Popule o banco com dados iniciais (opcional)**
   ```bash
   go run scripts/seed.go
   ```

## 🔧 Configuração

### Variáveis de Ambiente

Crie um arquivo `.env` baseado no `env.example`:

```env
# Configurações do Servidor
PORT=8080
ENV=development

# Configurações do Banco de Dados
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=destinara_db

# JWT Secret
JWT_SECRET=seu_jwt_secret_aqui

# CORS
ALLOWED_ORIGINS=http://localhost:5173,http://localhost:3000
```

## 🚀 Executando o Projeto

```bash
# Desenvolvimento
go run main.go

# Produção
go build -o destinara-backend
./destinara-backend
```

## 📚 API Endpoints

### Autenticação

- `POST /api/auth/register` - Registrar novo usuário
- `POST /api/auth/login` - Fazer login

### Pacotes de Viagem (Públicos)

- `GET /api/packages` - Listar todos os pacotes
- `GET /api/packages/:id` - Buscar pacote por ID

### Pacotes de Viagem (Admin - Requer Autenticação)

- `POST /api/admin/packages` - Criar novo pacote
- `PUT /api/admin/packages/:id` - Atualizar pacote
- `DELETE /api/admin/packages/:id` - Deletar pacote

### Reservas (Requer Autenticação)

- `GET /api/bookings` - Listar reservas do usuário
- `GET /api/bookings/:id` - Buscar reserva por ID
- `POST /api/bookings` - Criar nova reserva
- `PUT /api/bookings/:id` - Atualizar reserva
- `DELETE /api/bookings/:id` - Cancelar reserva

### Health Check

- `GET /health` - Verificar status da API

## 🔐 Autenticação

A API usa JWT (JSON Web Tokens) para autenticação. Para acessar rotas protegidas, inclua o header:

```
Authorization: Bearer <seu_token_jwt>
```

## 📊 Estrutura do Banco de Dados

### Tabelas

- **users**: Informações dos usuários
- **travel_packages**: Pacotes de viagem disponíveis
- **bookings**: Reservas dos usuários

### Relacionamentos

- Um usuário pode ter múltiplas reservas
- Uma reserva pertence a um usuário e um pacote de viagem
- Um pacote de viagem pode ter múltiplas reservas

## 🧪 Testando a API

### Exemplo de Registro

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@email.com",
    "password": "123456"
  }'
```

### Exemplo de Login

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "joao@email.com",
    "password": "123456"
  }'
```

### Exemplo de Listagem de Pacotes

```bash
curl -X GET http://localhost:8080/api/packages
```

## 📁 Estrutura do Projeto

```
backend/
├── config/          # Configurações
├── controllers/     # Controladores da API
├── database/        # Configuração do banco de dados
├── middleware/      # Middlewares (auth, cors)
├── models/          # Modelos de dados
├── routes/          # Definição de rotas
├── scripts/         # Scripts utilitários
├── main.go          # Arquivo principal
├── go.mod           # Dependências Go
└── README.md        # Documentação
```

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes. 