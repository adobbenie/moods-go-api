package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"main.go/db"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	cursor, err := db.UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []bson.M
	if err = cursor.All(context.Background(), &users); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(users)
}
