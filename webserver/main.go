// https://godoc.org/github.com/codegangsta/negroni

// https://dzone.com/articles/go-microservices-part-2-building-our-first-service
package main

import (
	"fmt"

	"github.com/abhishek-inspopindia/gomicro/webserver/src/service" // NEW
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
	fmt.Printf("Starting %v v%v.%v.%v\n", MyApp.name, MyApp.version.major, MyApp.version.minor, MyApp.version.patch)

	service.StartWebServer("6767") // NEW
}
