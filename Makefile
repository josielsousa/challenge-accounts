OS ?= linux
APP_NAME=challange-accounts
PROJECT_PATH ?= github.com/josielsousa/challenge-accounts
PKG ?= $(PROJECT_PATH)
ENVIRONMENT_STAGE=dev
VERSION ?= local

# docker files
DOCKER_FILE=api.dockerfile
DOCKER_COMPOSE_FILE=labs/docker-compose.yml

# git tags to add into release file
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_BUILD_TIME=$(shell date '+%Y-%m-%d__%I:%M:%S%p')
VCS_REF = $(if $(GITHUB_SHA),$(GITHUB_SHA),$(shell git rev-parse HEAD))

GOLANGCI_LINT_PATH=$$(go env GOPATH)/bin/golangci-lint
GOLANGCI_LINT_VERSION=1.62.2

define goBuild
	@echo "==> Go Building $2"
	@env GOOS=${OS} GOARCH=amd64 go build -v -o  build/$1 \
	-ldflags "-X main.BuildGitCommit=$(GIT_COMMIT) -X main.BuildTime=$(GIT_BUILD_TIME)" \
	${PKG}/$2
endef

define build
	@echo "==> Building Docker image: $1"
	@@DOCKER_BUILDKIT=1 docker build \
		--build-arg BUILD_TIME=$(GIT_BUILD_TIME) --build-arg GIT_COMMIT=$(GIT_COMMIT) --build-arg GIT_TAG=$(GIT_TAG) --build-arg GO_CMD=$1 \
		--pull --ssh default -f $(DOCKER_FILE) -t $(PROJECT_PATH)-$1:$(VERSION) .
endef

default: test-coverage

.PHONY: tools
tools:
	@echo "Installing dependencies"
	go mod tidy

.PHONY: clean
clean:
	@echo "==> Cleaning releases"
	@GOOS=${OS} go clean -i -x ./...
	@rm -f build/${APP_NAME}

.PHONY: compile
compile: clean tools gofmt
	@echo "==> Compiling releases"
	$(call goBuild,${APP_NAME},cmd)

.PHONY: build
build: compile
	$(call build,"api")

.PHONY: setup-down
setup-down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p ${APP_NAME} down --remove-orphans

.PHONY: setup-dev
setup-dev: setup-down
	@echo "==> Starting dev docker-compose"
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p ${APP_NAME} up --build -d

.PHONY: test
test:
	@echo "==> Running Tests"
	@go install github.com/rakyll/gotest@latest
	gotest -race -failfast -timeout 5m -count=1 -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "==> Running tests"
	go install github.com/rakyll/gotest@latest
	@gotest -race -failfast -timeout 5m -count=1 -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

.PHONY: docker-build
docker-build: compile
	@echo "==> Compiling docker images"
	docker image build --label "challange.accounts.vcs-ref=$(VCS_REF)" -t josielsousa/${APP_NAME}:${ENVIRONMENT_STAGE} build -f api.dockerfile

## Execute 'codelint'
lint: codelint

## Lint code using golangci-lint
codelint:
	@echo "==> Installing golangci-lint"
ifeq (,$(findstring $(GOLANGCI_LINT_VERSION),$(shell which $(GOLANGCI_LINT_PATH) && eval $(GOLANGCI_LINT_PATH) version)))
	@echo "installing golangci-lint v$(GOLANGCI_LINT_VERSION)"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v$(GOLANGCI_LINT_VERSION)
else
	@echo "already installed: $(shell eval $(GOLANGCI_LINT_PATH) version)"
endif
	@echo "==> Running golangci-lint"
	@$(GOLANGCI_LINT_PATH) run -c ./.golangci.yml --fix


## Execute 'mocks' + 'docs' + 'lint'
format: mocks gofmt lint

gofmt:
	@echo "==> Tidy modules"
	@go mod tidy
	@echo "==> GCI ${PROJECT_PATH}"
	@gci write --skip-generated -s standard -s default -s "prefix(github.com/josielsousa)" -s "prefix(${PROJECT_PATH})" -s blank -s dot .
	@echo "==> gofmt+"
	@gofumpt -w -extra .
	@go fmt ./...

## Generate mocks files
mocks:
	@echo "==> Generating mocks files"
	@git submodule update --init --remote
	@go generate ./...

## Check for vuln. using govulncheck
vuln:
	@echo "==> Installing go vuln check"
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@echo "==> Running go vuln check"
	@govulncheck ./...
