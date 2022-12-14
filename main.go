package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	r "main.go/router"
)

func main() {
	PORT := os.Getenv("PORT")
	router := r.Router()

	fmt.Println("Starting up server")
	err := http.ListenAndServe(PORT, router)
	if err != nil {
		log.Fatal(err)
	}
}
