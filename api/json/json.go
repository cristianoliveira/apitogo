package json

import (
  "fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
  "github.com/cristianoliveira/apitogo/common"
)

func HandleGetAll(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error: ", err)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func HandleGetById(w http.ResponseWriter, r *http.Request) {
  settings := common.Settings()
	vars := mux.Vars(r)

  path := settings.PathFile(vars["collection"])
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
