package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/svv1313/crud-api/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user", handlers.GetUser).Methods("GET")
	http.Handle("/", r)

    fmt.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}