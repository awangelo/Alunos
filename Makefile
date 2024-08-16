# Nome do arquivo de certificado e chave
DATABASE_PATH=./database.db

# Alvo para rodar o servidor
run:
	@echo "Iniciando variaveis de ambiente..."
	DATABASE_PATH=$(DATABASE_PATH) sudo go run cmd/alunos/main.go

.PHONY: run