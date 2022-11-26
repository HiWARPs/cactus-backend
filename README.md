# cactus-backend
The backend for Project Cactus

## Description:
1. Use the terminal command `docker-compose up` in the base directory of the repository to build the images and to start the containers. This should start the server.
2. To test the connection to the server, open http://localhost:3001/ in your web browser.
3. Open a second terminal window and use the command `docker ps` to see a list of all running containers. This allows you to see the container id.
4. To enter a container use the command `docker exec -it <container id> bash`
5. To write mongo queries, open the mongo shell within the mongo container with the command `mongosh`.


## Notes on upload a CSV file 

Dennis, feel free to remove this section once you are done with the upload mechanism.

I wrote a simple endpint which takes a CSV file, parses it and returns it as a response. To test this, you can follow 
these steps:

1. Start server.js
2. Make a POST request to http://localhost:3001/functions with a CSV file as body. You can use Postman for this.
   Or, you can use this curl call:
```curl
   curl --location --request POST 'localhost:3001/functions' --form 'uploaded_file=@"sample_functions.csv"'
```

Make sure you are in the same directory as the sample_functions.csv file.  You should get a response the looks like this:

```json
[
    {
        "a": "12",
        "b": "2",
        "c": "3"
    },
    {
        "a": "1",
        "b": "2",
        "c": "3"
    },
    {
        "a": "34",
        "b": "34",
        "c": "23"
    }
]
```