.DEFAULT_GOAL:=help
SHELL:=/bin/bash

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

build: ## Build the project
	go build -o teachstore-session

test: ## Test Service Ah! Perfectousing Mock
	go test

list-dep: ## List available minor and patch upgrades for all direct and indirect dependencies
	go list -u -m all	

upgrade: ## Upgrade to the lastest or minor patch release	
	go get -u ./...

docker-image: ## Build the docker image of this microservice with name: ualter/teachstore-session
	docker build . -t ualter/teachstore-session

docker-run: ## Run a Container with this microservice
	docker run -d -p 8383:9393 --name teachstore-session ualter/teachstore-session

run: ## Run main
	go run .


