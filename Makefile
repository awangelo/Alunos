# Caminho para a db
DATABASE_PATH=$(shell pwd)/database.db

# Alvo para rodar o servidor
run:
	@echo "Iniciando variaveis de ambiente..."
	DATABASE_PATH=$(DATABASE_PATH) sudo -E go run cmd/alunos/main.go

.PHONY: run