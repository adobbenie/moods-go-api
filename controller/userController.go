package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": id}

	cursor, err := db.UserCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var foundUser []bson.M
	if err = cursor.All(context.Background(), &foundUser); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(foundUser)
}

func updateUser(w http.ResponseWriter, r *http.Request) {

}
