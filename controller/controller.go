package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
	"main.go/model"
)

const connectionString = "mongodb+srv://user:user@moodsmusicdb.teextgp.mongodb.net/?retryWrites=true&w=majority"
const dbName = "MoodsMusicDb"
const colName = "users"

var collection *mongo.Collection

// connecten
// TODO kan ook vervangen worden met een background context maar dit zou steeds een open connectie houden
func init() {
	//maak options instantie met de con.string
	clientOption := options.Client().ApplyURI(connectionString)
	//gebruik de options om een connectie te maken met Atlas
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("succesfully connected to mongo")

	//maak instantie van de collectie die we zoeken
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

func insertNewuser(user model.User) {
	succes, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted user with id", succes.InsertedID)
}
