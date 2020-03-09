package config

import (
	"log"
	"strconv"
	"strings"
)

// Version => Model of App Version
type Version struct {
	Major int64
	Minor int64
	Patch int64
}

// App => Model of App
type App struct {
	Name    string
	Version Version
}

// Init => initializes the app config
func (app *App) Init(name string, version string) {
	// Split version into Semantic Versioning
	sv := strings.Split(version, ".")

	major, _ := strconv.Atoi(sv[0])
	minor, _ := strconv.Atoi(sv[1])
	patch, _ := strconv.Atoi(sv[2])

	app.Version.Major = int64(major)
	app.Version.Minor = int64(minor)
	app.Version.Patch = int64(patch)

	app.Name = name

	log.Printf("Starting %v v%v.%v.%v\n", app.Name, app.Version.Major, app.Version.Minor, app.Version.Patch)
}
