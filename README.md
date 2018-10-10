# goquiz

A simple quizzing application using Go and Vue.

## Setting up the environment

### Requirements

VirtualBox to host the Docker swarm.

```
// To get an embedded key/value store up and running.
go get -u github.com/boltdb/bolt

// Install test framework
go get -u github.com/smartystreets/goconvey/convey

// Install mocking library
go get -u github.com/stretchr/testify/mock
```

## Running the tests

```
go test ./...
```

## Building the services

```
cd ./account-service
CGO_ENABLED=0 GOOS=linux go build -o account-service-linux-amd64
```

## Deploying to local Docker swarm

```
// From root folder of the app.
docker build -t theapemachine/account-service account-service

// Create the swarm manager
docker-machine create --driver virtualbox --virtualbox-cpu-count 2 --virtualbox-memory 2048 --virtualbox-disk-size 20000 swarm-manager-1
```

## Serving the backend

```
// From root folder of the app.
docker run --rm theapemachine/account-service
```

## Serving the frontend

```
cd ./frontend
npm install
npm run dev
```
