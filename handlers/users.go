package handlers

import (
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "USER Created!")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET USER!")
}