package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//const connectionString_ = "mongodb+srv://user:user@moodsmusicdb.teextgp.mongodb.net/?retryWrites=true&w=majority"

var UserCollection *mongo.Collection
var PubRoomCollection *mongo.Collection
var PriRoomCollection *mongo.Collection

func init() {
	//connect to MongoDB Atlas
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("Your 'MONGODB_URI' environmental variable must be set by giving your connection string as it's value in the .env file.")
	}

	clientOption := options.Client().ApplyURI(uri)
	client_, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("succesfully connected to mongodb")

	//connect to database itself
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("Your 'DB_NAME' environmental variable must be set by giving setting the name of your database as it's value in the .env file.")
	}

	UserCollection = client_.Database(dbName).Collection("users")
	PubRoomCollection = client_.Database(dbName).Collection("public_rooms")
	PriRoomCollection = client_.Database(dbName).Collection("private_rooms")
	fmt.Println("All collections found")
}
