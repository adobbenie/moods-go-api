package model

import (
	"go.mongodb.org/mongo-driver/bson/primtive"
)

type song struct {
	SpotifyID string `json:"spotify_id"`
	TrackName string `json:"track_name"`
	Artist    string `json:"artist"`
	Duration  int    `json:"duration"`
}

// omitempty = empty values not allowed

type User struct {
	ID       primtive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MarsID   int               `json:"mars_id"`
	Username string            `json:"username"`
	Moodlist []song            `json:"mood_list"`
}
