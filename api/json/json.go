package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cristianoliveira/apitogo/api/store"
	"github.com/cristianoliveira/apitogo/common"
	"github.com/gorilla/mux"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func HandleGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	settings := common.Settings()
	vars := mux.Vars(r)

	path := settings.PathFile(vars["collection"])
	collection, err := CollectionLoad(path)
	if err != nil {
		handleNotFound(w, err)
		return
	}

	data, err := collection.AsBytes()
	if err != nil {
		handleNotFound(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func HandleGetById(repo *store.Store) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		vars := mux.Vars(r)

		docs, err := repo.Document(fmt.Sprint(vars["collection"], "-", vars["id"]))
		if err != nil {
			handleNotFound(w, err)
			return
		}
		if len(docs) == 0 {
			handleNotFound(w, errors.New("Data not found"))
			return
		}

		data, err := json.Marshal(docs)
		if err != nil {
			handleServerError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func handleNotFound(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	w.Write(NewError(http.StatusNotFound, err).AsBytes())
}

func handleBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(NewError(http.StatusBadRequest, err).AsBytes())
}

func handleServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(NewError(http.StatusInternalServerError, err).AsBytes())
}
