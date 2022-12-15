package main

import (
	"fmt"
	"net/http"
	"os"

	r "main.go/router"
)

func main() {
	PORT := os.Getenv("PORT")
	router := r.Router()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	})

	http.ListenAndServe(PORT, router)
}
