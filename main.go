package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/docopt/docopt-go"

	"github.com/cristianoliveira/apitogo/api"
)

// Default values
var settings api.Settings = api.Settings{
	Port: "8080",
	Dir:  "./",
}

const USAGE string = `Api to go, please.

Usage:
  apitogo run
  apitogo run [--port=<port>|--dir=<dir>]
  apitogo -h | --help
  apitogo --version

Options:
  --dir=<dir>   Dir containing the json files.
  --port=<port> Server port (Default 8080).
  -h --help     Show this screen.
  --version     Show version.
`

func main() {
	arguments, _ := docopt.Parse(USAGE, nil, true, "Naval Fate 2.0", false)
	fmt.Println(arguments)

	portArg := arguments["--port"]
	if portArg != nil {
		settings.Port = ":" + portArg.(string)
	}

	dirArg := arguments["--dir"]
	if dirArg != nil {
		settings.Dir = dirArg.(string)
	}

	files, err := filepath.Glob(settings.Dir + "/*.json")
	if err != nil {
		log.Fatal(err)
	}

	api.Serve(files, settings)
}
