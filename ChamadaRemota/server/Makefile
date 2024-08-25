# Variáveis
BINARY_NAME = server
MAIN_PATH = ./cmd/server/main.go
BUILD_DIR = ./bin

.PHONY: all clean build run

all: clean build run

clean:
	@echo "Limpando arquivos binários anteriores..."
	@rm -rf $(BUILD_DIR)

build:
	@echo "Compilando o binário..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

run: build
	@echo "Executando o programa..."
	@$(BUILD_DIR)/$(BINARY_NAME)

test:
	@echo "Executando testes..."
	@go test ./...

lint:
	@echo "Verificando lint..."
	@golangci-lint run

deps:
	@echo "Atualizando dependências..."
	@go mod tidy

fmt:
	@echo "Formatando o código..."
	@go fmt ./...

install-tools:
	@echo "Instalando ferramentas..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest


