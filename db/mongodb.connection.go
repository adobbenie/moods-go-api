package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString_ = "mongodb+srv://user:user@moodsmusicdb.teextgp.mongodb.net/?retryWrites=true&w=majority"
const DbName = "moodsdb"

var UserCollection *mongo.Collection
var PubRoomCollection *mongo.Collection
var PriRoomCollection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString_)
	client_, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("succesfully connected to mongodb")

	UserCollection = client_.Database(DbName).Collection("users")
	PubRoomCollection = client_.Database(DbName).Collection("public_rooms")
	PriRoomCollection = client_.Database(DbName).Collection("private_rooms")
	fmt.Println("All collection found")
}
