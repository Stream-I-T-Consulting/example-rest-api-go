# example-rest-api-go
Example REST API by Go Fiber Framework with Repository Pattern

(Route - Handler - Repository - Database)

___

## Developement

### Install packages
```
go install
```

### Run the development server
`go run .` or `go run main.go`

The server will listens on `http://localhost:3000`

___

## Build
### Build the executable file

`go build` 

or `go build -o main` for rename the executable file.

### Build the Docker image
```
docker build -t {DOCKER_HUB_USERNAME}/{REPOSITORY_NAME} .
```

Example: `docker build -t streaminnovation/example-go-rest-api .`

### Push the Docker image to registry host (hub.docker.io)
```
docker push {DOCKER_HUB_USERNAME}/{REPOSITORY_NAME}
```

Example: `docker push streaminnovation/example-go-rest-api`

___

## How to use this image
### Start a Go Fiber instance
```
docker run --name example-go-rest-api -e DATABASE_DSN="host=localhost user=innovation password=111222333 dbname=innovation_db port=5432 sslmode=disable TimeZone=Asia/Bangkok" -p 3000:3000 -d streaminnovation/example-go-rest-api:latest
```
### or via docker-compose
create the `.env` file in same `docker-compose.yml` directory
```yml
version: '3.8'

services:
  app:
    image: streaminnovation/example-go-rest-api:latest
    container_name: example-go-rest-api
    restart: always
    env_file:
      - .env
```

___

## Environment Variables
`DATABASE_DSN`

A DSN is a string with a specific data structure that describes a connection to a particular data source.