package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetAllPlayedSongs(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)["_id"]
	roomId, _ := primitive.ObjectIDFromHex(givenId)
	filter := bson.M{"_id": roomId}
	opts := options.Find().SetProjection(bson.D{{"played_songs", 1}, {"_id", 0}})

	cursor, err := db.PubRoomCollection.Find(Ctx, filter, opts)
	if err != nil {
		log.Fatal(err)
	}

	var foundRoom []bson.M
	if err = cursor.All(Ctx, &foundRoom); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(foundRoom)
}

func UpdatePlayedSongs(w http.ResponseWriter, r *http.Request) {

}
