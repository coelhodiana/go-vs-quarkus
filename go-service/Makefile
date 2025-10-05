# Makefile para automação de tarefas do go-service

# Declara que 'test' e 'coverage' não são nomes de ficheiros. É uma boa prática.
.PHONY: test coverage

## test: Executa os testes do serviço Go dentro de um contêiner Docker.
test:
	@echo "==> Executando testes para o go-service..."
	@docker run --rm -v "$(PWD):/app" -w /app golang:1.22-alpine go test ./...

## coverage: Gera o relatório de cobertura de testes em HTML.
coverage:
	@echo "==> Gerando ficheiro de cobertura (coverage.out)..."
	@docker run --rm -v "$(PWD):/app" -w /app golang:1.22-alpine go test -coverprofile=coverage.out ./...
	@echo "==> Gerando relatório HTML (coverage.html)..."
	@docker run --rm -v "$(PWD):/app" -w /app golang:1.22-alpine go tool cover -html=coverage.out -o coverage.html
	@echo "==> Relatório gerado! Abra o ficheiro coverage.html no seu navegador."