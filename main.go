package main

import (
  "fmt"
  "net/http"
  "log"
  "path/filepath"
  "io/ioutil"
  "encoding/json"
  "strconv"

  "github.com/gorilla/mux"
  "github.com/docopt/docopt-go"
)

type Settings struct {
  port string
  directory string
}

// Default values
var settings Settings = Settings {
  port: "8080",
  directory: "./",
}

const USAGE string = `Api to go, please.

Usage:
  apitogo run
  apitogo run [--port=<port>|--dir=<dir>]
  apitogo -h | --help
  apitogo --version

Options:
  --dir=<dir>   Directory containing the json files.
  --port=<port> Server port (Default 8080).
  -h --help     Show this screen.
  --version     Show version.
`

func main() {
	arguments, _ := docopt.Parse(USAGE, nil, true, "Naval Fate 2.0", false)
	fmt.Println(arguments)

  portArg := arguments["--port"]
  if portArg != nil { settings.port = ":" + portArg.(string) }

  dirArg := arguments["--dir"]
  if dirArg != nil { settings.directory = dirArg.(string) }

  files, err := filepath.Glob(settings.directory + "/*.json")
	if err != nil { log.Fatal(err) }

  fmt.Println("Server listening on: http://0.0.0.0:" + settings.port)
  log.Fatal(http.ListenAndServe("0.0.0.0:"+settings.port, route(files)))
}

func route(files []string) *mux.Router {
  router := mux.NewRouter().StrictSlash(true)

  for _, file := range files {
    endpoint := "/" + file[:len(file) - 5]
    fmt.Println(endpoint)
    fmt.Println(endpoint + "/:id")
	}

  router.HandleFunc("/{collection}", getAll)
  router.HandleFunc("/{collection}/{id}", getById)

  return router
}

func getAll(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  path, err := filepath.Abs(settings.directory)
  if err != nil {
    handleBadRequest(w, err)
    return
  }

  data, err := ioutil.ReadFile(path + "/" + vars["collection"] +".json")
  if err != nil {
    handleBadRequest(w, err)
    return
  }

  if err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintln(w, "Error: ", err)
  } else {
    w.WriteHeader(http.StatusOK)
    w.Write(data)
  }
}
func handleBadRequest(w http.ResponseWriter, err error) {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Println(w, err)
}

func getById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  path, err := filepath.Abs(settings.directory)
  if err != nil {
    handleBadRequest(w, err)
    return
  }

  data, err := ioutil.ReadFile(path + "/" + vars["collection"] +".json")
  if err != nil {
    handleBadRequest(w, err)
    return
  }

  pId, err := strconv.ParseFloat(vars["id"], 64)
  if err != nil {
    handleBadRequest(w, err)
    return
  }

  var jsonData interface{}
  err = json.Unmarshal(data, &jsonData)
  if err != nil {
    handleBadRequest(w, err)
    return
  }

  jsonMap := jsonData.(map[string]interface{})
  items := jsonMap[vars["collection"]].([]interface{})

  matchById := func(i map[string]interface{}) bool { return i["id"] == pId }
  selected := filter(items, matchById)

  if selected != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, selected)
  } else {
    w.WriteHeader(http.StatusNotFound)
    fmt.Println(w, nil)
  }
}

func filter(items []interface{}, isMatch func(map[string]interface{}) bool) map[string]interface{} {
  for i := range items {
    if isMatch(items[i].(map[string]interface{})) {
      return items[i].(map[string]interface{})
    }
  }

  return nil
}
