package v5

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type response_struct struct {
	Id      string
	Status  string
	Message string
}

func Pwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir + "/"
}

func FilesDir() string {
	return AssetsDir() + "files/"
}

func AssetsDir() string {
	return Pwd() + "assets/"
}

func API(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Landing
	if vars["slug1"] == "landing" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/landing/list.json")
		}
	}

	// Ingredients
	if vars["slug1"] == "ingredients" {
		http.ServeFile(w, r, AssetsDir()+"v5/ingredients/list.json")
	}

	// Pizza
	if vars["slug1"] == "pizza" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/pizza/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/pizza/"+vars["slug2"]+"/list.json")
		}

	}

	// Deserts
	if vars["slug1"] == "desserts" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/desserts/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/desserts/"+vars["slug2"]+"/list.json")
		}

	}

	// Appetizers
	if vars["slug1"] == "appetizers" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/appetizers/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/appetizers/"+vars["slug2"]+"/list.json")
		}

	}

	// Drinks
	if vars["slug1"] == "drinks" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/drinks/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/drinks/"+vars["slug2"]+"/list.json")
		}

	}

	// Sides
	if vars["slug1"] == "sides" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/sides/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/sides/"+vars["slug2"]+"/list.json")
		}

	}

	if vars["slug1"] == "checkout" {
		json.NewEncoder(w).Encode(response_struct{Id: vars["slug2"], Status: "success", Message: "Order was placed"})

	}

}
