# goquiz

A simple quizzing application using Go and Vue.

## Setting up the environment

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
```

## Serving the backend

```
// From root folder of the app.
docker run --rm theapemachine/account-service
```

## Serving the frontend
