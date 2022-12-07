package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type privateRoom struct {
	ID          primitive.ObjectID
	Owner       string
	Description string
}

type publicRoom struct {
	ID          primitive.ObjectID
	RaspID      int
	Description string
}
