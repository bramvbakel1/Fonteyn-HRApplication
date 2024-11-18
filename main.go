package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

var oauth2Config *oauth2.Config
var oauth2Token *oauth2.Token

// Initialize OAuth2 configuration using GitHub secrets or environment variables
func initOAuth2Config() {
	oauth2Config = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),                            // GitHub secret for client ID
		ClientSecret: os.Getenv("CLIENT_SECRET"),                        // GitHub secret for client secret
		RedirectURL:  os.Getenv("REDIRECT_URL"),                         // GitHub secret for redirect URL
		Scopes:       []string{"User.Read"},                             // You can add more Microsoft Graph scopes if needed
		Endpoint:     microsoft.AzureADEndpoint(os.Getenv("TENANT_ID")), // GitHub secret for tenant ID
	}
}

func main() {
	// Initialize OAuth2 config
	initOAuth2Config()

	// Define HTTPS routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/callback", callbackHandler)

	// Start HTTPS server
	log.Println("Server started on https://localhost:8080")
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// Home route that shows the login link
func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html") // Serving a simple index.html file with login link
}

// Login route to start the OAuth2 flow
func loginHandler(w http.ResponseWriter, r *http.Request) {
	authURL := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, authURL, http.StatusFound)
}

// Callback route to handle the OAuth2 callback and exchange the code for a token
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := oauth2Config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	oauth2Token = token

	// Use the token to access Microsoft Graph API
	client := resty.New()
	resp, err := client.R().
		SetAuthToken(oauth2Token.AccessToken).
		Get("https://graph.microsoft.com/v1.0/me")

	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Properly use the 'resp' variable (for example, log or display the response)
	if resp.StatusCode() != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error from Microsoft Graph API: %s", resp.Status()), http.StatusInternalServerError)
		return
	}

	// Example: Display user info from the response
	userInfo := resp.String() // Get the response body as string
	fmt.Fprintf(w, "Authenticated User Info: %s\n", userInfo)
}
