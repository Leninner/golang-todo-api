# Todo API with Golang

This is a simple Todo API written in Golang. It uses a simple postgres database to store the todos.

## Requirements

- Golang 1.22 >=
- Docker
- Air
- Docker
- Linux or MacOS

## Running the API

To run the API, you need to have a postgres database running. You can use the following command to run a postgres database using docker:

```bash
make up
```

This will start a postgres database on port 5432. You can then run the API using the following command:

```bash
air
```

This will start the API on port [8080](http://localhost:8080).