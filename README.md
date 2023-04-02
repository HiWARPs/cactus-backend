# cactus-backend

The backend for Project Cactus

## Running the app while developing

Use the following to run this app with auto restarts on changes:
npm start

This is possible due to importing the `nodemon` dependency and adding this in package.json

```json
{
  "scripts": {
    "start": "nodemon server.js"
  }
}
```

## Sample calls to the /project/:pid/form endpoint

### Create a form

```bash
curl --location 'http://localhost:3000/project/6425b800328e9c670b4b27b5/form' \
--header 'Content-Type: application/json' \
--data '{ 
    "name": "form name",
    "description": "form description", 
    "references" : "form references"
}'
```

### Get a single form

```bash
curl --location 'http://localhost:3000/project/6425b800328e9c670b4b27b5/form/6425bacdbd1388308db5a1bf'
```

## Description:

1. Use the terminal command `docker-compose up` in the base directory of the repository to build the images and to start
   the containers. This should start the server.
2. To test the connection to the server, open http://localhost:3001/ in your web browser.
3. Open a second terminal window and use the command `docker ps` to see a list of all running containers. This allows
   you to see the container id.
4. To enter a container use the command `docker exec -it <container id> bash`
5. To write mongo queries, open the mongo shell within the mongo container with the command `mongosh`.
