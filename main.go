package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Data struct
type Data struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Init data
var datas []Data
var allowedHeaders string = "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello GO!"))
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	json.NewEncoder(w).Encode(datas)
}

func postData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Data
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1000)) // Test ID
	datas = append(datas, user)
	json.NewEncoder(w).Encode(user)

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

func putData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range datas {
		if item.ID == params["id"] {
			datas = append(datas[:index], datas[index+1:]...)
			break

		}
	}

	var user Data
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	datas = append(datas, user)
	json.NewEncoder(w).Encode(user)
	return

}

func deleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range datas {
		if item.ID == params["id"] {
			datas = append(datas[:index], datas[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(datas)
}

func main() {

	//Add mock data
	datas = append(datas, Data{"1", "Satoshi"})
	datas = append(datas, Data{"2", "Pikachu"})

	r := mux.NewRouter()

	//GET
	r.HandleFunc("/", getHome).Methods("GET")
	r.HandleFunc("/data", getData).Methods("GET")
	r.HandleFunc("/data/{id}", getDataById).Methods("GET")

	//POST
	r.HandleFunc("/data", postData).Methods("POST")

	//PUT
	r.HandleFunc("/data/update/{id}", putData).Methods("PUT")

	//DELETE
	r.HandleFunc("/data/delete/{id}", deleteData).Methods("DELETE")

	// f := func(w http.ResponseWriter, r *http.Request) {
	// 	if origin := r.Header.Get("Origin"); origin != "" {
	// 		w.Header().Set("Access-Control-Allow-Origin", "*")
	// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	// 		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	// 	}
	// 	return
	// }

	log.Println("Listening on :3000...")
	http.ListenAndServe(":3000", r)
}
