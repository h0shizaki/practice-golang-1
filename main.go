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
var datas []Data

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello GO!"))
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datas)
}

func postData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func getDataById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get parameter
	//Loop search
	for _, item := range datas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Data{})
}

func main() {

	//Add mock data
	datas = append(datas, Data{"1", "Satoshi"})
	datas = append(datas, Data{"2", "Pikachu"})

	r := mux.NewRouter()

	r.HandleFunc("/", getHome).Methods("GET")
	r.HandleFunc("/data", getData).Methods("GET")
	r.HandleFunc("/data/{id}", getDataById).Methods("GET")
	r.HandleFunc("/data", postData).Methods("POST")

	http.ListenAndServe(":3000", r)
}
