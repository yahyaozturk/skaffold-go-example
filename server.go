package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func init() {
	hostname := os.Getenv("MONGOHOST")
	username := os.Getenv("MONGOUSER")
	password := os.Getenv("MONGOPASSWORD")
	log.Println("hostname", hostname)
	log.Println("username", username)
	log.Println("password", password)
	err := ConfigDB(hostname, username, password)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	log.Println("happy-birthday api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", homeLink)
	api.HandleFunc("/health", health)
	api.HandleFunc("/hello/{name}", saveORupdateRecord).Methods(http.MethodPut)
	api.HandleFunc("/hello/{name}", searchByName).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
