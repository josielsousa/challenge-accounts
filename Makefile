export TERM=xterm-256color
export CLICOLOR_FORCE=true
export RICHGO_FORCE_COLOR=1

default: test

# Faz o Build
build: install
	@echo "Fazendo build..."
	env GOOS=linux go build -o bin/service cmd/main.go

test: install test-coverage

setup-local: install
	@go get -u golang.org/x/tools/...
	@go get -u golang.org/x/lint/golint
	@go get -u github.com/haya14busa/goverage
	@go get -u github.com/kyoh86/richgo
	@go get github.com/joho/godotenv/cmd/godotenv

# Instala dependencias
install:
	@echo "Baixando depedencias..."
	@go mod verify

# Roda teste unitarios e gera coverage
test-coverage: 
	@echo "Rodando testes"
	@richgo test -tags 'nopie' -failfast -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@rm -rf DBTest*

clean:
	@rm -rf ./bin
	@rm -rf go.sum
	@rm -rf coverage*
