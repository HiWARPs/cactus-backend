# cactus-backend
The backend for Project Cactus

# Docker 

Dockerfile
---
Run the following commands to build the Docker image and run it:
```
docker build . -t kaktus/backend 
```
```
docker run -p 3001:3001 -d kaktus/backend
```
To call app:
```
curl -i localhost:3001
```

Docker-compose
---
```
docker-compose -f compose.yml up
```
