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
  apitogo run [-p <port>] [-d <dir>]
  apitogo -h | --help
  apitogo --version

Options:
  -d <dir>      Directory containing the json files.
  -p <port>     Server port (Default 8080).
  -h --help     Show this screen.
  --version     Show version.
`

func main() {
	arguments, _:= docopt.Parse(USAGE, nil, true, "apitogo "+VERSION, false)

	common.Settings().UpdateByArgs(arguments)

	api.Serve()
}
