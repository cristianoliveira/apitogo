package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Settings struct {
	Port string
	Dir string
}

var settings Settings

func Router() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/authorize", HandleAuthorization)
	router.HandleFunc("/token", HandleToken)

	router.HandleFunc("/{collection}", getAll)
	router.HandleFunc("/{collection}/{id}", getById)

  return router
}

func Serve(s Settings) {
  settings = s

	files, err := filepath.Glob(settings.Dir + "/*.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server listening on: http://0.0.0.0:" + settings.Port, "\n")
  fmt.Println("Endpoints for this folder:")

	for _, file := range files {
		endpoint := "/" + file[:len(file)-5]
		fmt.Println(endpoint)
		fmt.Println(endpoint + "/:id")
	}

  fmt.Println("")
  fmt.Println("Endpoints for oauth2:")
	fmt.Println("/authorize")
	fmt.Println("/token")
  fmt.Println("use client_id: 1234 and client_secret: apitogo1234")

	log.Fatal(http.ListenAndServe("0.0.0.0:" + settings.Port, Router()))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

  path := PathFile(settings.Dir, vars["collection"])
	collection, err := CollectionLoad(path)
	if err != nil {
		handleNotFound(w, err)
		return
	}

  data, err := collection.AsBytes()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error: ", err)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

  path := PathFile(settings.Dir, vars["collection"])
	collection, err := CollectionLoad(path)

	pId, err := strconv.ParseFloat(vars["id"], 64)
	if err != nil {
		handleNotFound(w, err)
		return
	}

	selected, err := collection.GetById(pId).AsBytes()
	if selected != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(selected)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(w, nil)
	}
}

func handleNotFound(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Println(w, err)
}

func handleBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Println(w, err)
}
