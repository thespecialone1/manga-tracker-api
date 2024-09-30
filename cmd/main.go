package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"manga-tracker-api/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", api.GetRoot)
	mux.HandleFunc("/manga/{$}",api.GetManga)

	log.Println("Starting Server to listen and serve on Port :3333")
	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server Closed\n")
	} else if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
		os.Exit(1)
	}
}
