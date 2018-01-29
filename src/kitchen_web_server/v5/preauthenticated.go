package v5

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/auth"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/response"

	"github.com/gorilla/mux"
)

// This is the first point that all GET requests go through.
// If a valid token is not provided as slug1, the API will not pass on the GET request to any authenticated endpoints.
func (obj *ObjectStore) PreAuthenticatedGetAPI(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	token := vars["slug1"]

	if obj.Tokens.Validate(token) {
		obj.postAuthenticatedGetAPI(w, r)
	} else {
		json.NewEncoder(w).Encode(response.Response{Status: false, Message: "User did not access the link with a valid token. Please log in"})
	}

}

// This is the first point that all POST requests go through.
// If a valid token is not provided as slug1, the API will not pass on the request to any authenticated endpoints.
func (obj *ObjectStore) PreAuthenticatedPostAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Login attempt
	if vars["slug1"] == "login" {

		decoder := json.NewDecoder(r.Body)
		var query auth.Login
		err := decoder.Decode(&query)
		if err != nil {
			fmt.Println("Could not decode query for login")
		}

		// AUTH
		if obj.Admins.Validate(query.Username, query.Password) {
			tokens := auth.GenerateNewToken(query.Username, 24*time.Hour)
			json.NewEncoder(w).Encode(tokens.Tokens[len(tokens.Tokens)-1])
		} else if obj.SuperAdmin.Validate(query.Username, query.Password) {
			tokens := auth.GenerateNewToken(query.Username, 24*time.Hour)
			json.NewEncoder(w).Encode(tokens.Tokens[len(tokens.Tokens)-1])
		} else {
			json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Username and Password combination does not exist. If you forgot your password, please talk to another manager for the app "})

		}

		// Other than logging in, everything else done on this API including the GET and POST can only be done if they user accomplishes this via the secure token link that was generated when they logged in
	} else {
		token := vars["slug1"]
		if obj.Tokens.Validate(token) {
			// if auth.Validate(token) {
			obj.postAuthenticatedPostAPI(w, r)
		} else {
			json.NewEncoder(w).Encode(response.Response{Status: false, Message: "User did not access link with valid token. Please log in"})

		}
	}

}
