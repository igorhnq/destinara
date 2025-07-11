# 🏖️ Destinara - Sistema de Pacotes de Viagem

Sistema completo de pacotes de viagem com frontend React/TypeScript e backend Go.

## 🚀 Visão Geral

O Destinara é uma aplicação web moderna para gerenciamento e reserva de pacotes de viagem. O projeto consiste em:

- **Frontend**: React + TypeScript + Vite + Tailwind CSS
- **Backend**: Go + Gin + GORM + PostgreSQL
- **Autenticação**: JWT
- **Containerização**: Docker + Docker Compose

## 📋 Funcionalidades

### Para Usuários
- ✅ Registro e login de usuários
- ✅ Visualização de pacotes de viagem
- ✅ Filtros por destino e disponibilidade
- ✅ Sistema de reservas
- ✅ Gestão de reservas pessoais

### Para Administradores
- ✅ CRUD completo de pacotes de viagem
- ✅ Gestão de disponibilidade
- ✅ Visualização de reservas

## 🛠️ Tecnologias Utilizadas

### Frontend
- **React 19** - Biblioteca JavaScript para interfaces
- **TypeScript** - Tipagem estática
- **Vite** - Build tool e dev server
- **Tailwind CSS** - Framework CSS utilitário
- **React Router** - Roteamento
- **Phosphor Icons** - Ícones

### Backend
- **Go 1.21** - Linguagem de programação
- **Gin** - Framework web
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **JWT** - Autenticação
- **CORS** - Cross-origin resource sharing

### DevOps
- **Docker** - Containerização
- **Docker Compose** - Orquestração de containers
- **Make** - Automação de tarefas

## 📁 Estrutura do Projeto

```
destinara/
├── frontend/                 # Aplicação React
│   ├── src/
│   │   ├── components/      # Componentes reutilizáveis
│   │   ├── pages/          # Páginas da aplicação
│   │   ├── services/       # Serviços de API
│   │   ├── layouts/        # Layouts da aplicação
│   │   └── assets/         # Recursos estáticos
│   ├── package.json
│   └── vite.config.ts
├── backend/                 # API Go
│   ├── config/             # Configurações
│   ├── controllers/        # Controladores da API
│   ├── database/           # Configuração do banco
│   ├── middleware/         # Middlewares
│   ├── models/             # Modelos de dados
│   ├── routes/             # Definição de rotas
│   ├── scripts/            # Scripts utilitários
│   ├── docs/               # Documentação
│   ├── main.go
│   └── go.mod
├── docker-compose.yml      # Orquestração Docker
└── README.md
```

## 🚀 Como Executar

### Pré-requisitos

- **Node.js 18+** e **npm/yarn**
- **Go 1.21+**
- **PostgreSQL 12+** (ou Docker)
- **Git**

### Opção 1: Execução Local

#### 1. Clone o repositório
```bash
git clone <url-do-repositorio>
cd destinara
```

#### 2. Configure o Backend
```bash
cd backend

# Instale as dependências
go mod tidy

# Configure as variáveis de ambiente
cp env.example .env
# Edite o arquivo .env com suas configurações

# Execute o servidor
go run main.go
```

#### 3. Configure o Frontend
```bash
cd ../frontend

# Instale as dependências
npm install

# Configure as variáveis de ambiente
cp env.example .env

# Execute o servidor de desenvolvimento
npm run dev
```

### Opção 2: Docker Compose (Recomendado)

```bash
# Clone o repositório
git clone <url-do-repositorio>
cd destinara

# Execute todos os serviços
docker-compose up -d

# Para parar os serviços
docker-compose down
```

### Opção 3: Scripts Automatizados

#### Backend
```bash
cd backend

# Setup completo
make dev-setup

# Executar
make run

# Com hot reload (requer air)
make dev
```

#### Frontend
```bash
cd frontend

# Instalar dependências
npm install

# Executar em desenvolvimento
npm run dev

# Build para produção
npm run build
```

## 🔧 Configuração

### Variáveis de Ambiente

#### Backend (.env)
```env
# Servidor
PORT=8080
ENV=development

# Banco de Dados
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=destinara_db

# JWT
JWT_SECRET=seu_jwt_secret_aqui

# CORS
ALLOWED_ORIGINS=http://localhost:5173,http://localhost:3000
```

#### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080/api
VITE_APP_TITLE=Destinara
```

## 📚 API Endpoints

### Autenticação
- `POST /api/auth/register` - Registrar usuário
- `POST /api/auth/login` - Fazer login

### Pacotes de Viagem
- `GET /api/packages` - Listar pacotes
- `GET /api/packages/:id` - Buscar pacote
- `POST /api/admin/packages` - Criar pacote (admin)
- `PUT /api/admin/packages/:id` - Atualizar pacote (admin)
- `DELETE /api/admin/packages/:id` - Deletar pacote (admin)

### Reservas
- `GET /api/bookings` - Listar reservas do usuário
- `POST /api/bookings` - Criar reserva
- `PUT /api/bookings/:id` - Atualizar reserva
- `DELETE /api/bookings/:id` - Cancelar reserva

## 🧪 Testes

### Backend
```bash
cd backend

# Executar todos os testes
make test

# Testes com cobertura
make test-coverage
```

### Frontend
```bash
cd frontend

# Executar testes
npm test

# Testes em modo watch
npm run test:watch
```

## 📊 Banco de Dados

### Migrações
As migrações são executadas automaticamente quando o servidor inicia.

### Dados Iniciais
```bash
cd backend
go run scripts/seed.go
```

### Estrutura das Tabelas

#### Users
- `id` (PK)
- `name`
- `email` (unique)
- `password` (hashed)
- `created_at`
- `updated_at`

#### TravelPackages
- `id` (PK)
- `name`
- `description`
- `destination`
- `price`
- `duration`
- `image_url`
- `available`
- `created_at`
- `updated_at`

#### Bookings
- `id` (PK)
- `user_id` (FK)
- `travel_package_id` (FK)
- `start_date`
- `end_date`
- `total_price`
- `status`
- `created_at`
- `updated_at`

## 🐳 Docker

### Construir Imagens
```bash
# Backend
docker build -t destinara-backend ./backend

# Frontend
docker build -t destinara-frontend ./frontend
```

### Executar com Docker Compose
```bash
# Desenvolvimento
docker-compose up -d

# Produção
docker-compose -f docker-compose.prod.yml up -d
```

## 📝 Desenvolvimento

### Comandos Úteis

#### Backend
```bash
make help          # Ver todos os comandos
make run           # Executar servidor
make dev           # Executar com hot reload
make test          # Executar testes
make seed          # Popular banco de dados
make fmt           # Formatar código
make lint          # Executar linter
```

#### Frontend
```bash
npm run dev        # Servidor de desenvolvimento
npm run build      # Build para produção
npm run preview    # Preview do build
npm run lint       # Executar linter
```

### Padrões de Código

- **Backend**: Seguir convenções Go
- **Frontend**: ESLint + Prettier
- **Commits**: Conventional Commits
- **Branches**: Git Flow

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 🆘 Suporte

Se você encontrar algum problema ou tiver dúvidas:

1. Verifique a documentação
2. Procure por issues existentes
3. Crie uma nova issue com detalhes do problema

## 🎯 Roadmap

- [ ] Sistema de pagamentos
- [ ] Avaliações e comentários
- [ ] Notificações por email
- [ ] Dashboard administrativo
- [ ] Sistema de cupons
- [ ] Integração com APIs de hotéis
- [ ] App mobile (React Native)

---

**Desenvolvido com ❤️ pela equipe Destinara** 