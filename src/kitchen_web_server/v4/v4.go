package v4

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../auth"
	"../utils"

	"github.com/gorilla/mux"
)

type response_struct struct {
	Id      string
	Status  string
	Message string
}

func GetAPI(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Landing
	if vars["slug1"] == "landing" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/landing/list.json")
		} else if vars["slug2"] == "set" {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/landing/set/list.json")
		}
	}

	// Ingredients
	if vars["slug1"] == "ingredients" {
		http.ServeFile(w, r, utils.AssetsDir()+"v4/ingredients/list.json")
	}

	// Pizza
	if vars["slug1"] == "pizza" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/pizza/list.json")
		} else {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/pizza/"+vars["slug2"]+"/list.json")
		}

	}

	// Deserts
	if vars["slug1"] == "desserts" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/desserts/list.json")
		} else {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/desserts/"+vars["slug2"]+"/list.json")
		}

	}

	// Appetizers
	if vars["slug1"] == "appetizers" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/appetizers/list.json")
		} else {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/appetizers/"+vars["slug2"]+"/list.json")
		}

	}

	// Drinks
	if vars["slug1"] == "drinks" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/drinks/list.json")
		} else {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/drinks/"+vars["slug2"]+"/list.json")
		}

	}

	// Sides
	if vars["slug1"] == "sides" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/sides/list.json")
		} else {
			http.ServeFile(w, r, utils.AssetsDir()+"v4/sides/"+vars["slug2"]+"/list.json")
		}

	}

	if vars["slug1"] == "checkout" {
		json.NewEncoder(w).Encode(response_struct{Id: vars["slug2"], Status: "success", Message: "Order was placed"})

	}
}

func PostAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Login attempt
	if vars["slug1"] == "login" {

		authorized_users := auth.ValidUsersList{}
		authorized_users.Initialize()

		decoder := json.NewDecoder(r.Body)
		var query auth.Login
		err := decoder.Decode(&query)
		if err != nil {
			fmt.Println("Could not decode query for login")
		}

		if authorized_users.Validate(query.Username, query.Password) {
			tokens := auth.GenerateNewToken(query.Username)
			json.NewEncoder(w).Encode(tokens.Tokens[len(tokens.Tokens)-1])
		}
	}

	if vars["slug1"] == "landing" {
		if vars["slug2"] == "set" {
			json.NewEncoder(w).Encode(response_struct{Id: vars["slug2"], Status: "success", Message: "Order was placed"})
		}
	}

}
