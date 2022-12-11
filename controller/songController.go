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

func UpdateMoodList(w http.ResponseWriter, r *http.Request) {
	givenId := r.URL.Query().Get("mars_id")
	filter := bson.M{"mars_id": givenId}
	var updatedList []model.Song
	json.NewDecoder(r.Body).Decode(&updatedList)

	result, err := db.UserCollection.UpdateOne(
		context.Background(),
		filter,
		bson.D{
			{"$set", bson.D{{"mood_list", updatedList}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Succesfully updated a moodlist with count:", result.ModifiedCount)
}
