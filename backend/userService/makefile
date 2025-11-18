include .env

.PHONY: run
run: build
	@./bin/app

.PHONY: build
build:
	@echo "-> building..."
	go build -o ./bin/app ./cmd/app

.PHONY: test
test:
	go test -v ./internal/application/services \
	./internal/interface/api/rest_test


.PHONY: docs-gen
docs-gen:
	swag init -g cmd/app/main.go