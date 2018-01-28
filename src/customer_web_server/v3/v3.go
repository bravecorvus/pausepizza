package v3

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func Pwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir + "/"
}

func AssetsDir() string {
	return Pwd() + "assets/"
}

func API(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Landing
	if vars["slug1"] == "landing" {
		http.ServeFile(w, r, AssetsDir()+"v3/landing/list.json")
	}

	// Pizza, Desserts, Appetizers, Drinks, Sides General List

	// Pizza
	if vars["slug1"] == "pizza" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v3/pizza/list.json")
		} else if vars["slug2"] == "specialty" {
			fmt.Println(vars["slug3"])
			if vars["slug3"] == "" {
				http.ServeFile(w, r, AssetsDir()+"v3/pizza/specialty/list.json")
			} else {
				http.ServeFile(w, r, AssetsDir()+"v3/pizza/specialty/"+vars["slug3"]+"/list.json")
			}
		} else if vars["slug2"] == "build" {
			if vars["slug3"] == "" {
				http.ServeFile(w, r, AssetsDir()+"v3/pizza/build/list.json")
			} else {
				http.ServeFile(w, r, AssetsDir()+"v3/pizza/build/"+vars["slug3"]+"list.json")
			}
		}
	}

	// Deserts
	if vars["slug1"] == "desserts" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v3/desserts/list.json")
		}
	}

	// Appetizers
	if vars["slug1"] == "appetizers" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v3/appetizers/list.json")
		}
	}

	// Drinks
	if vars["slug1"] == "drinks" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v3/drinks/list.json")
		}
	}

	// Sides
	if vars["slug1"] == "sides" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v3/sides/list.json")
		}
	}

}
