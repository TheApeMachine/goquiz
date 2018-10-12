# goquiz

A simple quizzing application using Go and Vue.

## Setting up the environment

### Requirements

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

## Running the services

```
cd ./account-service
go run *.go

cd ./quiz-service
go run *.go
```

## Serving the frontend

```
cd ./frontend
npm install
npm run dev

visit http://localhost:8080 in the browser
```
