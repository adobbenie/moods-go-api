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

func UpdateMoodList(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)["mars_id"]
	filter := bson.M{"mars_id": givenId}

	var updatedList []model.Song
	json.NewDecoder(r.Body).Decode(&updatedList)

	result, err := db.UserCollection.UpdateOne(
		Ctx,
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
