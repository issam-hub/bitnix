include .env

.PHONY: run
run: build
	@./bin/app

.PHONY: build
build:
	@echo "-> building..."
	go build -o ./bin/app ./cmd/app

.PHONY: migrate
migrate:
	@echo "-> Creating migration files..."
	migrate create -seq -ext=.sql -dir=./migrations $$name


MIGRATE_PATH := migrate -path=./migrations -database=${DSN}

.PHONY: migrateRun
migrateRun:
	@echo "-> Running up migrations..."
	@$(MIGRATE_PATH) up

.PHONY: migrateReset
migrateReset:
	@echo "-> Running down migrations..."
	@$(MIGRATE_PATH) down

.PHONY: migrateGoto
migrateGoto:
	@echo "-> Running goto migrations..."
	@$(MIGRATE_PATH) goto $$version

.PHONY: migrateForce
migrateForce:
	@echo "-> Running force migrations..."
	@$(MIGRATE_PATH) force $$version

.PHONY: test
test:
	go test -v ./internal/application/services \
	./internal/interface/api/rest_test


.PHONY: docs-gen
docs-gen:
	swag init -g cmd/app/main.go