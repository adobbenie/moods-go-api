package router

import (
	"github.com/gorilla/mux"
	c "main.go/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//USER ROUTES
	router.HandleFunc("/api/users", c.GetAllUsers).Methods("GET")

	//PUBLIC ROOM ROUTES

	//PRIVATE ROOM ROUTES
	router.HandleFunc("/api/rooms/public", c.GetPublicRoomById).Methods("GET")    //id added as string query param.
	router.HandleFunc("/api/rooms/public", c.CreateNewPublicRoom).Methods("POST") //room data as JSON in body (raspid, location, description)
	router.HandleFunc("/api/rooms/public", c.DeletePublicRoom).Methods("DELETE")  //id added as string query param.
	router.HandleFunc("/api/rooms/public", c.UpdatePublicRoom).Methods("PUT")

	return router
}
