# Variáveis
APP_NAME=auth-api
# PORT=8080
# DB_URL=postgres://seu_usuario:sua_senha@localhost:5432/auth_db?sslmode=disable

# Rodar a aplicação localmente
run:
	go run main.go

# Build do binário
build:
	go build -o $(APP_NAME) main.go

# Rodar testes
test:
	go test ./... -v

# Rodar a aplicação com Docker
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Criar banco e aplicar migrations (usando Goose, Gorm ou outra lib)
migrate:
	go run cmd/migrate.go

# Limpar binário
clean:
	rm -f $(APP_NAME)

# Rodar linter
lint:
	golangci-lint run
