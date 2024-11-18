package handler

import (
	"html/template"
	"log"
	"net/http"
	"net/mail"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Println("Error loading template: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func UserManagementHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Get the form data
		username := r.FormValue("username")
		email := r.FormValue("email")

		// input validation
		if username == "" || email == "" {
			http.Error(w, "Both username and email are required.", http.StatusBadRequest)
			return
		}

		// Email validation using net/mail package
		_, err := mail.ParseAddress(email)
		if err != nil {
			http.Error(w, "Invalid email address.", http.StatusBadRequest)
			return
		}

		log.Printf("User created: Username: %s, Email: %s\n", username, email)

		tmpl, err := template.ParseFiles("templates/user_management.html")
		if err != nil {
			log.Println("Error loading template: ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Success bool
			Error   string
		}{
			Success: true,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println("Error executing template: ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	tmpl, err := template.ParseFiles("templates/user_management.html")
	if err != nil {
		log.Println("Error loading template: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
