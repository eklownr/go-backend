package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ReadBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    title := vars["title"]
    fmt.Fprintf(w, "Du läser boken: %s\n", title)
}   