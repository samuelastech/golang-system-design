package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "3333"

type Config struct {
}

func main() {
	app := Config{}

	log.Println("Starting broker service on port", port)

	// Define the server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	// Starts the server
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
