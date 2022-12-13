package controller

import (
	"context"
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

var Ctx = context.Background()

func GetAllPrivateRooms(w http.ResponseWriter, r *http.Request) {
	cursor, err := db.PriRoomCollection.Find(Ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var rooms []bson.M
	if err = cursor.All(Ctx, &rooms); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(rooms)
}

func GetPrivateRoomById(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)["_id"]
	ObjId, err := primitive.ObjectIDFromHex(givenId)
	if err != nil {
		log.Fatal(err)
	}

	cursor, err := db.PriRoomCollection.Find(Ctx, bson.M{"_id": ObjId})
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

func CreateNewPrivateRoom(w http.ResponseWriter, r *http.Request) {
	var newPriRoom model.PrivateRoom
	json.NewDecoder(r.Body).Decode(&newPriRoom)

	succes, err := db.PriRoomCollection.InsertOne(Ctx, newPriRoom)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert a new private room with id:", succes.InsertedID)
}

func UpdatePrivateRoom(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)["_id"]
	ObjId, _ := primitive.ObjectIDFromHex(givenId)
	var updatedPrivRoom model.PrivateRoom
	json.NewDecoder(r.Body).Decode(&updatedPrivRoom)

	result, err := db.PriRoomCollection.ReplaceOne(
		Ctx,
		bson.M{"_id": ObjId},
		updatedPrivRoom,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Succesfully updated a private room with count: ", result.ModifiedCount)
}

func DeletePrivateRoom(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)["_id"]
	ObjId, _ := primitive.ObjectIDFromHex(givenId)

	deleteCount, err := db.PriRoomCollection.DeleteOne(Ctx, bson.M{"_id": ObjId})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Private room deleted with count: ", deleteCount)
}
