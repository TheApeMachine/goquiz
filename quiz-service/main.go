package main

import (
	"fmt"

	"github.com/theapemachine/goquiz/quiz-service/dbclient"
	"github.com/theapemachine/goquiz/quiz-service/service"
)

var appName = "quiz-service"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6768")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
