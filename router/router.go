package router

import (
	"github.com/gorilla/mux"
	c "main.go/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//USER ROUTES
	router.HandleFunc("/api/users", c.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{mars_id}", c.GetUserById).Methods("GET")
	router.HandleFunc("/api/users/{mars_id}/data", c.UpdateUser).Methods("PUT")

	//SONG ROUTES
	router.HandleFunc("/api/users/{mars_id}/list", c.UpdateMoodList).Methods("PUT")
	router.HandleFunc("/api/rooms/public/{_id}/played", c.GetAllPlayedSongs).Methods("GET")
	router.HandleFunc("/api/rooms/public/{_id}/played", c.UpdatePlayedSongs).Methods("PUT")

	//PUBLIC ROOM ROUTES
	router.HandleFunc("/api/rooms/public/{_id}", c.GetPublicRoomById).Methods("GET")
	router.HandleFunc("/api/rooms/public", c.CreateNewPublicRoom).Methods("POST")
	router.HandleFunc("/api/rooms/public/{_id}", c.DeletePublicRoom).Methods("DELETE")
	router.HandleFunc("/api/rooms/public/{_id}", c.UpdatePublicRoom).Methods("PUT")

	//PRIVATE ROOM ROUTES
	router.HandleFunc("/api/rooms/private", c.GetAllPrivateRooms).Methods("GET")
	router.HandleFunc("/api/rooms/private/{_id}", c.GetPrivateRoomById).Methods("GET")
	router.HandleFunc("/api/rooms/private", c.CreateNewPrivateRoom).Methods("POST")
	router.HandleFunc("/api/rooms/private/{_id}", c.UpdatePrivateRoom).Methods("PUT")
	router.HandleFunc("/api/rooms/private/{_id}", c.DeletePrivateRoom).Methods("DELETE")

	return router
}
