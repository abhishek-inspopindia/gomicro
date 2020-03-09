// https://godoc.org/github.com/codegangsta/negroni
// https://dzone.com/articles/go-microservices-part-2-building-our-first-service
package main

import (
	config "github.com/abhishek-inspopindia/gomicro/webserver/src/config"
	"github.com/abhishek-inspopindia/gomicro/webserver/src/dbclient"
	service "github.com/abhishek-inspopindia/gomicro/webserver/src/services"
)

func init() {
	// Initializing App Configuration
	appConfig := new(config.App)

	appName := "My Go Microservice"
	appVersion := "1.0.0"

	appConfig.Init(appName, appVersion)

	// Initializing BoltClient
	initializeBoltClient()
}

func main() {
	// Defining Port
	port := "6767"
	// Starting WebServer
	service.StartWebServer(port) // NEW
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
