package server

import (
	"log"
	"net/http"
)

// StartServer starts the HTTPS server with the given certificate and key
func StartServer(addr, certFile, keyFile string) error {
	err := http.ListenAndServeTLS(addr, certFile, keyFile, nil)
	if err != nil {
		log.Printf("Error starting HTTPS server: %v", err)
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

		return err
	}
	return nil
}
