package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main.go/model"
	"net/http"
)

func CreateNewPublicRoom(w http.ResponseWriter, r *http.Request)  {
	var newPubRoom model.PublicRoom

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewDecoder(r.Body).Decode(&newPubRoom)

	succes, err := PubRoomCollection.InsertOne(context.Background(), newPubRoom)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted user with id", succes.InsertedID)
}

