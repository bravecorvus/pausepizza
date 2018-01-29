package v5

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type response_struct struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	OrderID string `json:"orderID"`
}

type kitchen_response_struct struct {
	status  bool   `json:"Status"`
	message string `json:"Message"`
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

func GetAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Landing
	if vars["slug1"] == "landing" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/landing/list.json")
		}

		// Ingredients
	} else if vars["slug1"] == "ingredients" {
		http.ServeFile(w, r, AssetsDir()+"v5/ingredients/list.json")

		// Pizza
	} else if vars["slug1"] == "pizza" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/pizza/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/pizza/"+vars["slug2"]+"/list.json")
		}

		// Deserts
	} else if vars["slug1"] == "desserts" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/desserts/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/desserts/"+vars["slug2"]+"/list.json")
		}

		// Appetizers
	} else if vars["slug1"] == "appetizers" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/appetizers/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/appetizers/"+vars["slug2"]+"/list.json")
		}

		// Drinks
	} else if vars["slug1"] == "drinks" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/drinks/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/drinks/"+vars["slug2"]+"/list.json")
		}

		// Sides
	} else if vars["slug1"] == "sides" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/sides/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/sides/"+vars["slug2"]+"/list.json")
		}

	}

}

func PostAPI(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	if vars["slug1"] == "checkout" {
		CheckOut(w, r)

	} else if vars["slug1"] == "ordercomplete" {
		orderid := vars["slug2"]
		fmt.Println("Order Complete for order " + orderid)
		json.NewEncoder(w).Encode(response_struct{Status: true, Message: "Order fulfilled POST request received"})
	}
}
