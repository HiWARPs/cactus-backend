docker_clean:
	-docker stop cactus_mongodb
	-docker rm cactus_mongodb
	-docker rmi cactus/mongodb

docker_build:
	docker build -f cactus_mongo_db/Dockerfile -t cactus/mongodb .
	echo "`docker images | grep cactus`"

docker_run: docker_clean docker_build
	docker run -d -p 27017:27017 --name cactus_mongodb cactus/mongodb
	echo "Point your DB client to localhost:27017 to connect to this DB"

.PHONY: docker_run docker_build