// https://godoc.org/github.com/codegangsta/negroni

// https://dzone.com/articles/go-microservices-part-2-building-our-first-service
package main

import (
	"log"

	"github.com/abhishek-inspopindia/gomicro/webserver/src/dbclient"
	"github.com/abhishek-inspopindia/gomicro/webserver/src/service"
)

type version struct {
	major int64
	minor int64
	patch int64
}
type app struct {
	name    string
	version version
}

func main() {
	MyAppVersion := version{1, 0, 0}
	MyApp := app{"Go Microservices", MyAppVersion}
	log.Printf("Starting %v v%v.%v.%v\n", MyApp.name, MyApp.version.major, MyApp.version.minor, MyApp.version.patch)

	initializeBoltClient() // NEW

	service.StartWebServer("6767") // NEW
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
