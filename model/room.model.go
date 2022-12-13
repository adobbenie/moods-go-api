package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PrivateRoom struct {
	_Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Owner       string             `json:"owner"`
	Description string             `json:"description"`
	Passcode    string             `json:"passcode"`
}

type PublicRoom struct {
	_Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Rpi_ID      string             `json:"rpi_id"`
	Location    string             `json:"location"`
	Description string             `json:"description"`
	Played_Song []string           `json:"played_songs"`
}
