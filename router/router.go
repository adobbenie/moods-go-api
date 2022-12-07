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
	router.HandleFunc("/api/rooms/public", c.CreateNewPublicRoom).Methods("POST")
	//router.HandleFunc("/api/rooms/public/{id}", c.DeletePublicRoom).Methods("DELETE")

	return router
}
