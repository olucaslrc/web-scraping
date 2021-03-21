package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/home/").Handler(http.StripPrefix("/home/", http.FileServer(http.Dir("./app/"))))
	r.HandleFunc("/scrap", scrap).Methods("POST")
	http.Handle("/", r)
	log.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
