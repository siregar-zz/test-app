package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"myapp/api"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/counter", api.CheckCounter).Methods("GET")
	r.HandleFunc("/api/users", api.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", api.DeleteUser).Methods("DELETE")
	r.HandleFunc("/api/users/{iduser}/{idbarang}", api.DelBarang).Methods("DELETE")
	r.HandleFunc("/api/users/{id}", api.AddBarang).Methods("POST")
	r.HandleFunc("/api/users", api.GetAllUser).Methods("GET")
	if err := http.ListenAndServe(":8088", r); err != nil {
		log.Fatal(err)
	}
}	

