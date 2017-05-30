package main

import (
	"github.com/cristianoliveira/apitogo/api"
	"github.com/cristianoliveira/apitogo/common"
	"github.com/docopt/docopt-go"
)

const VERSION = "0.1.0"

const USAGE string = `Api to go, please.

Usage:
  apitogo run
  apitogo run [-p <port>] [--dir <dir>]
  apitogo -h | --help
  apitogo --version

Options:
  --dir <dir>   Directory containing the json files.
  -p <port>     Server port (Default 8080).
  -h --help     Show this screen.
  --version     Show version.
`

func main() {
	arguments, _ := docopt.Parse(USAGE, nil, true, "apitogo "+VERSION, false)
	settings := common.Settings()

	portArg := arguments["--port"]
	if portArg != nil {
		settings.Port = portArg.(string)
	}

	dirArg := arguments["--dir"]
	if dirArg != nil {
		settings.Dir = dirArg.(string)
	}

	api.Serve()
}
