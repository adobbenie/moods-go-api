package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type song struct {
	SpotifyID     string   `json:"id"`
	TrackName     string   `json:"name"`
	Artist        []string `json:"artists"`
	Duration      int      `json:"duration"`
	AlbumCoverUrl string   `json:"url"`
}

// omitempty = empty values not allowed

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MarsID    int                `json:"mars_id"`
	SpotifyID string             `json:"spotify_user_id"`
	Username  string             `json:"username"`
	Moodlist  []song             `json:"mood_list"`
}
