package controller

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://user:user@moodsmusicdb.teextgp.mongodb.net/?retryWrites=true&w=majority"
const dbName = "MoodsMusicDb"
const colName = "users"

var collection *mongo.Collection

// connecten
// TODO kan ook vervangen worden met een background context maar dit zou steeds een open connectie houden
func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		console.log(err)
	}
	fmt.Print.ln()
}
