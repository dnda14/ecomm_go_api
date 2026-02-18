# Variables
CMD_DIR := cmd
API_MAIN := $(CMD_DIR)/api/main.go

# Ejecutar API (archivo específico)
.PHONY: run-api
run-api:
	go run $(API_MAIN)

# Ejecutar migraciones
.PHONY: run-migrate
run-migrate:
	go run $(MIGRATE_MAIN)

# Ejecutar todo el directorio cmd/api (recomendado)
.PHONY: run
run:
	go run ./$(CMD_DIR)