# cactus-backend

The backend for Project Cactus

## Running the app while developing


*Step 1*: Start the database

To start the DB you can follow the [instructions in Confluence](https://oregonstate-innovationlab.atlassian.net/wiki/spaces/VOVA/pages/61145089/Docker+and+MongoDB).

Or you can use the Makefile. 

In a terminal run `make docker_run`. That will rebuild and run the mongodb docker image.

Running `docker ps` will show a running mongodb container.

```bash
CONTAINER ID   IMAGE            COMMAND                  CREATED         STATUS         PORTS                      NAMES
c7e14c84ed84   cactus/mongodb   "docker-entrypoint.s…"   3 minutes ago   Up 3 minutes   0.0.0.0:27017->27017/tcp   cactus_mongodb
```

*Step 2*: Start the backend


```bash

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

If you don't have your `.env` file setup to connect to the database, you can use the following command to run the app with a local database:

```bash
export DATABASE_URL=mongodb://localhost:27017 && npm start
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

## Uploading a file

When running locally, you can upload a file with the following curl command. 
The file sample.csv is located in the test_data directory.


```bash
curl --location 'http://localhost:3000/file' \
--form 'file=@"sample.csv"'

```


## Working with the download endpoint

/download is the endpoint

curl -o /tmp/what_i_downloaded -v http://localhost:3000/download

This will download a file in the /tmp directory.


## Description:

1. Use the terminal command `docker-compose up` in the base directory of the repository to build the images and to start
   the containers. This should start the server.
2. To test the connection to the server, open http://localhost:3001/ in your web browser.
3. Open a second terminal window and use the command `docker ps` to see a list of all running containers. This allows
   you to see the container id.
4. To enter a container use the command `docker exec -it <container id> bash`
5. To write mongo queries, open the mongo shell within the mongo container with the command `mongosh`.
