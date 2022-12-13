package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"main.go/db"
	"main.go/model"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	cursor, err := db.UserCollection.Find(Ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []bson.M
	if err = cursor.All(Ctx, &users); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)["mars_id"]
	filter := bson.M{"mars_id": givenId}

	cursor, err := db.UserCollection.Find(Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var foundUser []bson.M
	if err = cursor.All(Ctx, &foundUser); err != nil {
		log.Fatal(err)
	}
	fmt.Println(foundUser)
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(foundUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)["mars_id"]
	filter := bson.M{"mars_id": givenId}

	var updatedUser model.User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	result, err := db.UserCollection.UpdateOne(
		Ctx,
		filter,
		bson.D{
			{"$set", bson.D{
				{"username", updatedUser.Username},
				{"spotify_user_id", updatedUser.Spotify_User_Id},
			},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Succesfully updated a user with count: ", result.ModifiedCount)
}
