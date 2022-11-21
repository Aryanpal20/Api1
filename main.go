package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Info struct {
	id      string   `json:"id"`
	subject string   `json:"subject"`
	mark    string   `json:"mark`
	student *student `json:"student"`
}

type student struct {
	Firstname string `json:"firstname`
	Lastname  string `json:"lastname"`
}

var infos []Info

// func getinfos(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(infos)

// }

// this function will delete the info

func deleteInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range infos {
		if item.id == params["id"] {
			infos = append(infos[:index], infos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(infos)

}

//thid function will

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range infos {
		if item.id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Info Info
	_ = json.NewDecoder(r.Body).Decode(&Info)
	Info.id = strconv.Itoa(rand.Intn(100000000))
	infos = append(infos, Info)
	json.NewEncoder(w).Encode(Info)

}

func updateInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range infos {
		if item.id == params["id"] {
			infos = append(infos[:index], infos[index+1:]...)
			var Info Info
			_ = json.NewDecoder(r.Body).Decode(&Info)
			Info.id = params["id"]
			infos = append(infos, Info)
			json.NewEncoder(w).Encode(Info)
		}
	}
}

func main() {
	r := mux.NewRouter()

	infos = append(infos, Info{id: "1", subject: "math", mark: "80", student: &student{Firstname: "ram", Lastname: "das"}})
	infos = append(infos, Info{id: "2", subject: "math", mark: "85", student: &student{Firstname: "dev", Lastname: "das"}})
	infos = append(infos, Info{id: "3", subject: "math", mark: "70", student: &student{Firstname: "ravi", Lastname: "das"}})

	r.HandleFunc("/infos", getinfos).Methods("GET")
	r.HandleFunc("/infos/{id}", getInfo).Methods("GET")
	r.HandleFunc("/infos/{id}", createInfo).Methods("POST")
	r.HandleFunc("/infos/{id}", updateInfo).Methods("PUT")
	r.HandleFunc("/infos/{id}", deleteInfo).Methods("DELETE")

	fmt.Println("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
