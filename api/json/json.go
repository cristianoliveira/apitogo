package json

import (
	"net/http"
	"strconv"

	"github.com/cristianoliveira/apitogo/common"
	"github.com/gorilla/mux"
)

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

func HandleGetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	settings := common.Settings()
	vars := mux.Vars(r)

	path := settings.PathFile(vars["collection"])
	collection, err := CollectionLoad(path)

	pId, err := strconv.ParseFloat(vars["id"], 64)
	if err != nil {
		handleBadRequest(w, err)
		return
	}

	selected, err := collection.GetById(pId).AsBytes()
	if err != nil {
		handleNotFound(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(selected)
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
