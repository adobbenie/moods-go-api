package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Artists  []string `json:"artists"`
	Duration int      `json:"duration"`
	Url      string   `json:"url"`
}

// omitempty = empty values not allowed

type User struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Mars_Id         int                `json:"mars_id"`
	Spotify_User_Id string             `json:"spotify_user_id"`
	Username        string             `json:"username"`
	Mood_list       []Song             `json:"mood_list"`
}
