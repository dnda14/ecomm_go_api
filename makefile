# Variables
CMD_DIR := cmd


# Ejecutar todo el directorio 
.PHONY: run
run:
	go run ./$(CMD_DIR)