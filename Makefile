.PHONY: list-commands

list-commands:
	@bash -c "./list-commands.sh"

dependencies: 
	go mod download

build:
	- @echo "build....${BUILD_ARG}"

run-api:
	- cd api && go run main.go bootstrap.go

platform-db:
	@docker-compose -f docker-compose.yml --profile db up

platform-down:
	@docker-compose -f docker-compose.yml down 

migrate:
	@echo "migrate"	$(m-status)
	go run ./cmd/migrations/main.go $(m-status)




