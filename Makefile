BINARY_NAME=ecli
WINDOWS=$(BINARY_NAME).exe
LINUX=$(BINARAY_NAME)
DARWIN=$(BINARY_NAME) 
VERSION=$(shell git describe --tags --always --long --dirty)
SHELL=/bin/bash

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

build: windows linux  # darwin has to be called explicitly

$(WINDOWS):
	# env GOOS=windows GOARCH=amd64 go build -i -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GOFLAGS="-ldflags=-s -ldflags=-w" go build -o ${WINDOWS}

$(LINUX):
	# env GOOS=linux GOARCH=amd64 go build -i -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go
	CGO_ENABLED=0 GOOS=linux GOFLAGS="-ldflags=-s -ldflags=-w" go build -o ${LINUX}

$(DARWIN):
	# env GOOS=darwin GOARCH=amd64 go build -i -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/service/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 GOFLAGS="-ldflags=-s -ldflags=-w" go build -o ${DARWIN}

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-SA4006 ./...
	go test -race -buildvcs -vet=off ./...

## clean: clean go artefacts (binary included)
clean:
	go clean
	rm ${WINDOWS} ${LINUX} ${DARWIN}

## release: test build and audit current code
release: test build audit security

## security: perform security check
security:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: launch quick tests
test: 
	./run_tests.sh

## cover: launch coverage
cover:
	go tool cover -html=./coverage.out

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## help: display this usage
.PHONY: help
help:
	@echo 'Usage:'
	@echo ${MAKEFILE_LIST}
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]
