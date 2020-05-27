package main

import (
	"fmt"
	"net/http"
	"time"
)

func init() {
	fmt.Println("Starting server http://localhost:8080/")
}

func main() {
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", indexHandler)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v", err.Error())
		return
	}
}
