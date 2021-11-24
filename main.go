package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Data struct
type Data struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Init data
var data []Data

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello GO!"))
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {

	//Add mock data
	data = append(data, Data{"1", "Satoshi"})
	data = append(data, Data{"2", "Pikachu"})

	r := mux.NewRouter()

	r.HandleFunc("/", getHome).Methods("GET")
	r.HandleFunc("/data", getData).Methods("GET")

	http.ListenAndServe(":3000", r)
}
