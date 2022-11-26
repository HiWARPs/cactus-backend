# cactus-backend
The backend for Project Cactus

## Description:
1. Use the terminal command `docker-compose up` in the base directory of the repository to build the images and to start the containers. This should start the server.
2. To test the connection to the server, open http://localhost:3001/ in your web browser.
3. Open a second terminal window and use the command `docker ps` to see a list of all running containers. This allows you to see the container id.
4. To enter a container use the command `docker exec -it <container id> bash`
5. To write mongo queries, open the mongo shell within the mongo container with the command `mongosh`.


## Debugging the server

Having the code in a running container makes debugging tedious. To debug the server easier, run the DB in a container 
and the server in your local machine. Here are the steps:

1. Run the DB in a container with `docker run -p 27017:27017 --name cactus-mongo -d mongo`