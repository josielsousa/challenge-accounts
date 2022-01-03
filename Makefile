OS ?= linux
APP_NAME=challange-accounts
PROJECT_PATH ?= github.com/josielsousa/challenge-accounts
PKG ?= $(PROJECT_PATH)
ENVIRONMENT_STAGE=dev

# docker files
DOCKER_FILE=api.dockerfile
DOCKER_COMPOSE_FILE=labs/docker-compose.yml

# git tags to add into release file
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_BUILD_TIME=$(shell date '+%Y-%m-%d__%I:%M:%S%p')
VCS_REF = $(if $(GITHUB_SHA),$(GITHUB_SHA),$(shell git rev-parse HEAD))

define goBuild
	@echo "==> Go Building $2"
	@env GOOS=${OS} GOARCH=amd64 go build -v -o  build/$1 \
	-ldflags "-X main.BuildGitCommit=$(GIT_COMMIT) -X main.BuildTime=$(GIT_BUILD_TIME)" \
	${PKG}/$2
endef

default: test-coverage

.PHONY: install-dependencies
install-dependencies:
	@echo "Installing dependencies"
	go mod tidy


.PHONY: clean
clean:
	@echo "==> Cleaning releases"
	@GOOS=${OS} go clean -i -x ./...
	@rm -f build/${APP_NAME}

.PHONY: compile
compile: clean install-dependencies gofmt
	@echo "==> Compiling releases"
	$(call goBuild,${APP_NAME},cmd)
	zip -r build/${APP_NAME}.zip  build/${APP_NAME}

.PHONY: setup-down
setup-down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) -p ${APP_NAME} down

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

.PHONY: metalint
metalint:
ifeq (, $(shell which $$(go env GOPATH)/bin/golangci-lint))
	@echo "==> installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin
	go install ./...
endif
	$$(go env GOPATH)/bin/golangci-lint run --fix --allow-parallel-runners -c ./.golangci.yml ./...

.PHONY: gofmt
gofmt: 
	@echo "==> Formating code"
	go mod tidy
	go install github.com/daixiang0/gci@latest
	gci -local ${PROJECT_PATH} -w .
	go install mvdan.cc/gofumpt@latest
	gofumpt -w -extra .
	go fmt ./...