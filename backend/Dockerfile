# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o destinara-backend .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário compilado
COPY --from=builder /app/destinara-backend .

# Expor porta
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./destinara-backend"] 