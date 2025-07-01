include $(wildcard .env)

ifeq ($(OS),Windows_NT)
	PLATFORM := Windows
else
	PLATFORM := $(shell uname)
endif

CGO_ENABLED=1

.PHONY: gen-db
gen-db:
	dev gen:db

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: migrate
migrate:
	dev migrate

.PHONY: migrate-down
migrate-down:
	dev migrate down

.PHONY: migrate-status
migrate-status:
	dev migrate status

.PHONY: migrate-version
migrate-version:
	dev migrate version

.PHONY: watch
watch:
	@go install github.com/air-verse/air@latest
ifeq ($(PLATFORM),Windows)
	@air
else
	@air -build.cmd "go build -o ./bin/app ." -build.full_bin "./bin/app" -build.args_bin "api"
endif

.PHONY: build
build:
	go build -o ./bin/app .

.PHONY: swagger
swagger:
	dev gen:swagger

.PHONY: mocks
mocks:
	@go install github.com/vektra/mockery/v3@v3.3.6
	@mockery --config internal/services/.mockery.yml
	@mockery --config storage/.mockery.yml

.PHONY: version
version:
	@app version

.PHONY: monitor
monitor:
	docker compose up -d victoriametrics victorialogs vmagent vmauth grafana

.PHONY: monitor-down
monitor-down:
	docker compose down victoriametrics victorialogs vmagent vmauth grafana

.PHONY: services
services:
	docker compose up -d db redis

.PHONY: services-down
services-down:
	docker compose down db redis

.PHONY: run-server
run-server:
	docker compose up -d server

.PHONY: stop-server
stop-server:
	docker compose down server
