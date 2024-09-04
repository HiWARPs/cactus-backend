# hiwarp-backend
API for HiWarp

## Using Taskfile

See the following for installation: https://taskfile.dev/

`task docker_clean` <-- destroys local mongo
`task docker_run` <-- starts local mongo
`task run` <-- runs the service 

## Running the mongo db

TLDR: `make docker_run`

That will run MongoDB in a container.  You can then run `make run` to run the app which will connect to this DB.

## Usage

1. To start the hi-warp API server, run the following command: 

        go run main.go

Using the Makefile, you can run:

        make run

2. The API server should now be running and accessible at http://localhost:3000 (where 3000 is the configured port number from the .env file).

## Testing The Health Point

To test the health endpoint of the API, you can use a tool like cURL or a REST client such as Postman. The health endpoint provides basic information about the status of the API server.

1. Make a GET request to the following endpoint: GET http://localhost:3000/health

You can type the following in a command prompt:


        curl http://localhost:3000/health

2. If the API server is running successfully, you should receive a response with a status code of 200 and a JSON object containing the health information.


# Making queries

Import a CSV file like this: 

```csv
x1,x2,y1,y2
0.1,1.1,2,3
0.2,1.1,2,3
0.3,1.2,2,3
0.4,1.2,2,3
0.5,1.1,2,3
```

Then you can query the API like this:

```bash

curl --location 'localhost:3000/query_electrons' \
--header 'Content-Type: application/json' \
--data '{
    "id": "6539433c7618dd3f7c014e72",
    "range": {
        "name": "x1",
        "lower_bound": 0.1,
        "upper_bound": 0.4,
        "increment": 0.1
    },
    "x":[
        {
            "name": "x2",
            "value": 1.2
        }
    ],
    "exclude": ["y1"]
}'
```

You should get a result like below with filters as well as the increments.

```json 

{
    "id": "6539433c7618dd3f7c014e72",
    "range": {
        "name": "x1",
        "lower_bound": 0.1,
        "upper_bound": 0.4,
        "increment": 0.1
    },
    "x":[
        {
            "name": "x2",
            "value": 1.2
        }
    ],
    "exclude": ["y1"]
}

```

## Another Example

Using this csv:

```csv
x1,x2,y1,y2
0.1,1.1,2,3
0.2,1.1,2,3
0.3,1.2,2,3
0.4,1.2,2,3
0.5,1.1,2,3
```

Query the API

```bash

curl --location 'localhost:3000/query_electrons' \
--header 'Content-Type: application/json' \
--data '{
    "id": "6539433c7618dd3f7c014e72",
    "range": {
        "name": "x1",
        "lower_bound": 0.1,
        "upper_bound": 0.4,
        "increment": 0.1
    },
    "x":[],
    "exclude": []
}'

```

It will return this result

```json 

{
    "_id": "6539433c7618dd3f7c014e72",
    "data": [
        {
            "x1": 0.1,
            "x2": 1.1,
            "y1": 2,
            "y2": 3
        },
        {
            "x1": 0.2,
            "x2": 1.1,
            "y1": 2,
            "y2": 3
        },
        {
            "x1": 0.3,
            "x2": 1.2,
            "y1": 2,
            "y2": 3
        },
        {
            "x1": 0.4,
            "x2": 1.2,
            "y1": 2,
            "y2": 3
        }
    ]
}

```
