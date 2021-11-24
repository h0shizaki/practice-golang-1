package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Data struct

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello GO!"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getHome).Methods("GET")

	http.ListenAndServe(":3000", r)
}
