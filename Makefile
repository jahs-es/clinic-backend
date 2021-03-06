.PHONY: all
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-api

build-api: 
	go build -tags $(LIBRARY_ENV) -o ./bin/api src/application/main.go

build-api-prod:
	go build -tags prod -o ./bin/api src/application/main.go

centos-api-prod:
	CGO_ENABLED=0 GOOS=linux go build -tags prod -o ./bin/api src/application/main.go

linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(LIBRARY_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/api api/main.go

ci: dependencies test	

test:
	go test -tags testing ./...

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

# First set TAG to create
generate-docker:
	docker-compose build

# First set TAG to launch
launch-docker:
	cd docker && docker-compose up -d

stop-docker:
	cd docker && docker-compose down --remove-orphans --volumes
