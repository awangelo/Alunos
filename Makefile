# Nome do arquivo de certificado e chave
CERT_FILE=cert.pem
KEY_FILE=key.pem
DATABASE_PATH=./database.db

# Alvo para rodar o servidor
run:
	@echo "Iniciando variaveis de ambiente..."
	CERT_FILE=$(CERT_FILE) KEY_FILE=$(KEY_FILE) go run cmd/alunos/main.go

.PHONY: run