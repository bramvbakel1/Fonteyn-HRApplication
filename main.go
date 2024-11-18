package main

import (
	"log"
	"net/http"

	"hrapplication/handler"
)

func main() {
	certFile := "./certs/server.crt"
	keyFile := "./certs/server.key"

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/users", handler.UserManagementHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Starting HTTPS server on port 443...")
	err := http.ListenAndServeTLS(":443", certFile, keyFile, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
