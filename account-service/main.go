package main

import (
	"fmt"

	"github.com/theapemachine/goquiz/account-service/dbclient"
	"github.com/theapemachine/goquiz/account-service/service"
)

var appName = "account-service"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
