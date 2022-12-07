package main

import (
	"fmt"
	"log"
	"net/http"

	r "main.go/router"
)

func main() {
	var PORT = ":1234"
	router := r.Router()

	fmt.Println("Starting up server")
	err := http.ListenAndServe(PORT, router)
	if err != nil {
		log.Fatal(err)
	}
}
