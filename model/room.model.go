package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PrivateRoom struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Owner       string	`json:"owner"`
	Description string	`json:"description"`
	Passcode	string	`json:"passcode"`
}

type PublicRoom struct {
	ID          primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	RaspID      int	`json:"rasp-id"`
	Location 	string	`json:"location"`
	Description string	`json:"description"`
}
