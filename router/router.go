package router

import (
	"github.com/gorilla/mux"
	c "main.go/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//USER ROUTES
	router.HandleFunc("/api/users", c.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user", c.GetUserById).Methods("GET")     //marsId as query string parameters
	router.HandleFunc("/api/user/data", c.UpdateUser).Methods("PUT") //complete user object in body (changeable field: username, spotify_user_id)

	//MOODLIST ROUTES
	router.HandleFunc("/api/user/list", c.UpdateMoodList).Methods("PUT") //only moodlist in body => array of songs

	//PRIVATE ROOM ROUTES
	router.HandleFunc("/api/rooms/private", c.GetAllPrivateRooms).Methods("GET")
	router.HandleFunc("/api/rooms/private/{id}", c.GetPrivateRoomById).Methods("GET")
	router.HandleFunc("/api/rooms/private", c.CreateNewPrivateRoom).Methods("POST")
	router.HandleFunc("/api/rooms/private", c.UpdatePrivateRoom).Methods("PUT")
	router.HandleFunc("/api/rooms/private/{id}", c.DeletePrivateRoom).Methods("DELETE")

	//PUBLIC ROOM ROUTES
	router.HandleFunc("/api/rooms/public", c.GetPublicRoomById).Methods("GET")    //id added as string query param.
	router.HandleFunc("/api/rooms/public", c.CreateNewPublicRoom).Methods("POST") //room data as JSON in body (raspid, location, description)
	router.HandleFunc("/api/rooms/public", c.DeletePublicRoom).Methods("DELETE")  //id added as string query param.
	router.HandleFunc("/api/rooms/public", c.UpdatePublicRoom).Methods("PUT")     ////id added as string query param. + complete room info should be in request body

	return router
}
