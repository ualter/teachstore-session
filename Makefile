.DEFAULT_GOAL:=help
SHELL:=/bin/bash
.PHONY: help build test upgrade run

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[33m\n\nTargets:\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-18s\033[33m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

build: ## Build the project
	go build -o teachstore-session

test: ## Test Service Ah! Perfectousing Mock
	go test

list-dep: ## List available minor and patch upgrades for all direct and indirect dependencies
	go list -u -m all	

upgrade: ## Upgrade to the lastest or minor patch release	
	go get -u ./...

#AQUI!!!! LOAD IMAGE TO Kind!
kind-load: show-env ## Load the Docker Image Application to Kind Cluster Named k8s3nodes
	@echo -e "\033[0;33m---------> Loading to Kind Cluster Kubernetes\033[0;0m"
	kind load docker-image teachstore-session:${VERSION} teachstore-session:${VERSION} --name k8s3nodes

docker-rm: ## Remove the teachstore-session image created
	@echo -e "\033[0;33m---------> Removing teachstore-session\033[0;0m"
	docker rmi $$(docker images teachstore-session -qa)

docker-build: show-env docker-rm ## Build the docker image of this microservice with name: teachstore-session:${VERSION}
	@echo -e "\033[0;33m---------> Building teachstore-session\033[0;0m"
	docker build . -t teachstore-session:${VERSION}

docker-run: show-env ## Run a Container with this microservice
	docker run -d -p 8383:9393 \
	       -e IP_DOCKER_HOST=${IP_DOCKER_HOST} \
	       --name teachstore-session teachstore-session:${VERSION}

show-env: check-env # Display the needed environment variables for some targets here
	@echo -e "\033[0;33mVersion for images..: \033[0;94m$(VERSION)\033[0;0m"
	@echo -e "\033[0;33mDocker Host IP......: \033[0;94m$(IP_DOCKER_HOST)\033[0;0m"

check-env: # Check environment variables necessary for some targets here
ifndef VERSION
	$(error VERSION environment variable is undefined, please set it: "export VERSION=1.0.0")
endif
ifndef IP_DOCKER_HOST
	$(error IP_DOCKER_HOST environment variable is undefined, please set it: "export IP_DOCKER_HOST=192.168.1.45")
endif

run: ## Run Go application
	go run .

helm-dry: show-env ## Test Helm running in dry-run mode
	helm install teachstore-session \
      --set version=${VERSION} \
	  --set ipDockerHost=${IP_DOCKER_HOST} \
      --set image=teachstore-session:${VERSION} \
      --set-file 'configValues=config/config.yaml' \
      k8s/helm/teachstore-session \
      --dry-run

helm-build-install: docker-build kind-load helm-install ## Perform Docker Build, Docker Load at Kind K8s Cluster and then Helm Install
	@echo -e "\033[0;33mComplete workflow performed: \033[0;94mDocker Build\033[0;33m --> \033[0;94mDocker Load to K8s Kind Cluster \033[0;33m--> \033[0;94mHelm Install Microservices\033[0;0m"

helm-install: show-env ## Install the teachstore-session Helm chart (Deployment, Services, Ingress, ConfigMap, Role, RoleBinding, etc.)
	helm install teachstore-session \
      --set version=${VERSION} \
	  --set ipDockerHost=${IP_DOCKER_HOST} \
      --set image=teachstore-session:${VERSION} \
      --set-file 'configValues=config/config.yaml' \
      k8s/helm/teachstore-session

helm-ls: ## List the helm charts installed with "teachstore" in name
	helm ls --filter teachstore

helm-del: ## Erase the teachstore-session chart (Consequently all its Kubernetes workloads creates)
	helm delete teachstore-session