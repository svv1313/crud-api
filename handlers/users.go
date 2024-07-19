package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/svv1313/crud-api/config"
	"github.com/svv1313/crud-api/constants"
	"github.com/svv1313/crud-api/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	defer client.Disconnect(context.Background())

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database(constants.DATABASE_NAME).Collection(constants.COLLECTION_USERS)
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    json.NewEncoder(w).Encode(result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET USER!")
}