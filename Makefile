GOOSE_DBSTRING= "root:123456@tcp(127.0.0.1:3306)/goshop"
GOOSE_MIGRATION_DIR ?= sql/schema


# Name APP
APP_NAME := server

# Run app
dev:
	go run ./cmd/$(APP_NAME)

run:
	docker compose up -d && go run ./cmd/$(APP_NAME)

kill:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down

upse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=${GOOSE_MIGRATION_DIR} up

downse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=${GOOSE_MIGRATION_DIR} down

resetse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=${GOOSE_MIGRATION_DIR} reset


.PHONY: run downse upse resetse

.PHONY: air