# Caminho para a db
DATABASE_PATH=$(shell pwd)/database.db

# Rodar com go run para testes
test:
	@echo "Iniciando variaveis de ambiente..."
	DATABASE_PATH=$(DATABASE_PATH) sudo -E go run cmd/alunos/main.go

# Builda para a arch do sistema atual
build:
	@echo "Iniciando build..."
	go build -o alunos cmd/alunos/main.go
	@echo ok

# Alvo para rodar o servidor
run:
	DATABASE_PATH=$(DATABASE_PATH) sudo -E ./alunos

# Builda e inicia o servidor
all: build run

.PHONY: test run all
