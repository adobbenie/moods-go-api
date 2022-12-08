package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main.go/model"
)

func GetPublicRoomById(w http.ResponseWriter, r *http.Request) {
	RoomId := r.URL.Query().Get("id")
	id, _ := primitive.ObjectIDFromHex(RoomId)
	filter := bson.M{"_id": id}

	cursor, err := PubRoomCollection.Find(context.Background(), filter)
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

	succes, err := PubRoomCollection.InsertOne(context.Background(), newPubRoom)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted public room with id", succes.InsertedID)
}

func DeletePublicRoom(w http.ResponseWriter, r *http.Request) {
	RoomId := r.URL.Query().Get("id")
	id, _ := primitive.ObjectIDFromHex(RoomId)
	filter := bson.M{"_id": id}

	deleteCount, err := PubRoomCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Public room deleted with count: ", deleteCount)

}
