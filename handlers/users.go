package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/svv1313/crud-api/config"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	defer client.Disconnect(context.Background())
    fmt.Fprintf(w, "USER Created!")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET USER!")
}