package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main.go/db"
	"main.go/model"
)

func GetPublicRoomById(w http.ResponseWriter, r *http.Request) {
	RoomId := r.URL.Query().Get("id")
	id, _ := primitive.ObjectIDFromHex(RoomId)
	filter := bson.M{"_id": id}

	cursor, err := db.PubRoomCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var foundRoom []bson.M
	if err = cursor.All(context.Background(), &foundRoom); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(foundRoom)
}

func CreateNewPublicRoom(w http.ResponseWriter, r *http.Request) {
	var newPubRoom model.PublicRoom

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewDecoder(r.Body).Decode(&newPubRoom)

	succes, err := db.PubRoomCollection.InsertOne(context.Background(), newPubRoom)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a new public room with id:", succes.InsertedID)
}

func UpdatePublicRoom(w http.ResponseWriter, r *http.Request) {
	RoomId := r.URL.Query().Get("id")
	id, _ := primitive.ObjectIDFromHex(RoomId)
	var updatedPubRoom model.PublicRoom
	json.NewDecoder(r.Body).Decode(&updatedPubRoom)

	result, err := db.PubRoomCollection.ReplaceOne(
		context.Background(),
		bson.M{"_id": id},
		updatedPubRoom,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Succesfully updated a public room with count: ", result.ModifiedCount)
}

func DeletePublicRoom(w http.ResponseWriter, r *http.Request) {
	RoomId := r.URL.Query().Get("id")
	id, _ := primitive.ObjectIDFromHex(RoomId)
	filter := bson.M{"_id": id}

	deleteCount, err := db.PubRoomCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Public room deleted with count: ", deleteCount)

}
