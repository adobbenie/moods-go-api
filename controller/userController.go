package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"main.go/db"
	"main.go/model"
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
	givenId := r.URL.Query().Get("mars_id")
	//id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"mars_id": givenId}

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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	givenId := r.URL.Query().Get("mars_id")
	//id, _ := primitive.ObjectIDFromHex(givenId)
	filter := bson.M{"mars_id": givenId}

	var updatedUser model.User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	result, err := db.UserCollection.UpdateOne(
		context.Background(),
		filter,
		bson.D{
			{"$set", bson.D{
				{"username", updatedUser.Username},
				{"spotify_user_id", updatedUser.SpotifyID},
			},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Succesfully updated a user with count: ", result.ModifiedCount)
}
