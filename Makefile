# ==================================== #
# HELPERS
# ==================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==================================== #
# DEVELOPMENT
# ==================================== #

## run: run the main application
.PHONY: run
run:
	go run ./main.go

# ==================================== #
# QUALITY CHECK
# ==================================== #

## tidy: format all .go files and tidy module dependencies
.PHONY: tidy
tidy:
	@echo 'Formatting .go files...'
	go fmt ./...
	@echo 'Tidying module dependencies...'
	go mod tidy
	@echo 'Verifying module dependencies'
	go mod verify

## audit: run quality control checks
.PHONY: audit
audit:
	@echo 'Checking module dependencies...'
	go mod tidy -diff
	go mod verify
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	go test -race -vet=off ./...

## test: run unit tests
.PHONY: test
test:
	@echo 'Running tests...'
	go test -race ./...

# ==================================== #
# BUILD
# ==================================== #

# build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api
