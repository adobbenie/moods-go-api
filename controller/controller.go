package controller

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://user:user@moodsmusicdb.teextgp.mongodb.net/?retryWrites=true&w=majority"
const dbName = "moodsdb"
const colName = "users"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("succesfully connected to mongo")

	//maak instantie van de collectie die we zoeken
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []bson.M
	if err = cursor.All(context.Background(), &users); err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}
