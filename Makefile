help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

db_clean: ## Stop and remove the mongo container and the image
	-docker stop cactus_mongodb
	-docker rm cactus_mongodb
	-docker rmi cactus/mongodb

db_build: ## Build mongo image
	docker build -f Dockerfile -t cactus/mongodb .
	echo "`docker images | grep cactus`"

db_run: db_clean db_build ## rebuild and run the mongo image
	docker run -d -p 27017:27017 --name cactus_mongodb cactus/mongodb
	echo "Point your DB client to localhost:27017 to connect to this DB"

build: ## Build the service in a temp directory
	echo "building the hiwarp service"
	go build -o /tmp/hiwarp

run: build ## build and run service
	echo "running the service"
	/tmp/hiwarp

compose: db_build ## Build and run the service and db in a docker container
	docker-compose up --build

clean: ## Stop the service and remove the containers and the images
	-docker stop cactus_mongodb
	-docker stop cactus-backend-backend-1
	-docker stop cactus-backend-mongodb-1
	-docker rm cactus_mongodb
	-docker rm cactus-backend-backend-1
	-docker rm cactus-backend-mongodb-1
	-docker rmi cactus/mongodb
	-docker rmi cactus-backend-backend


.PHONY: docker_run docker_build
