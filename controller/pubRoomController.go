package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main.go/db"
	"main.go/model"
)

func GetPublicRoomById(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)["_id"]
	roomId, _ := primitive.ObjectIDFromHex(givenId)
	filter := bson.M{"_id": roomId}

	cursor, err := db.PubRoomCollection.Find(Ctx, filter)
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

func CreateNewPublicRoom(w http.ResponseWriter, r *http.Request) {
	var newPubRoom model.PublicRoom

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewDecoder(r.Body).Decode(&newPubRoom)

	succes, err := db.PubRoomCollection.InsertOne(Ctx, newPubRoom)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a new public room with id:", succes.InsertedID)
}

func UpdatePublicRoom(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)["_id"]
	roomId, _ := primitive.ObjectIDFromHex(givenId)
	filter := bson.M{"_id": roomId}

	var updatedPubRoom model.PublicRoom
	json.NewDecoder(r.Body).Decode(&updatedPubRoom)

	result, err := db.PubRoomCollection.ReplaceOne(
		Ctx,
		filter,
		updatedPubRoom,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Succesfully updated a public room with count: ", result.ModifiedCount)
}

func DeletePublicRoom(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)["_id"]
	roomId, _ := primitive.ObjectIDFromHex(givenId)
	filter := bson.M{"_id": roomId}

	deleteCount, err := db.PubRoomCollection.DeleteOne(Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Public room deleted with count: ", deleteCount)
}
