package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"main.go/db"
)

var ctx = context.Background()

func GetAllPrivateRooms(w http.ResponseWriter, r *http.Request) {
	cursor, err := db.PriRoomCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var rooms []bson.M
	if err = cursor.All(ctx, &rooms); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(rooms)
}

func GetPrivateRoomById(w http.ResponseWriter, r *http.Request) {

}

func CreateNewPrivateRoom(w http.ResponseWriter, r *http.Request) {

}

func UpdatePrivateRoom(w http.ResponseWriter, r *http.Request) {

}

func DeletePrivateRoom(w http.ResponseWriter, r *http.Request) {

}
