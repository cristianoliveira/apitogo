package api

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
  "github.com/cristianoliveira/apitogo/common"
  "github.com/cristianoliveira/apitogo/api/auth"
  "github.com/cristianoliveira/apitogo/api/json"
)


func Router() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/authorize", auth.HandleAuthorization)
	router.HandleFunc("/token", auth.HandleToken)

	router.HandleFunc("/{collection}", json.HandleGetAll)
	router.HandleFunc("/{collection}/{id}", json.HandleGetById)

  return router
}

func Serve() {
  settings := common.Settings()
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

