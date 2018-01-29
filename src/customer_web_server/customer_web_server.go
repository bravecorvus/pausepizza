package main

import (
	"net/http"

	"github.com/gilgameshskytrooper/pausepizza/src/customer_web_server/v5"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v5/{slug1}", v5.PostAPI).Methods("POST")
	r.HandleFunc("/v5/{slug1}/{slug2}/{slug3}", v5.GetAPI).Methods("GET")
	r.HandleFunc("/v5/{slug1}/{slug2}", v5.GetAPI).Methods("GET")
	r.HandleFunc("/v5/{slug1}", v5.GetAPI).Methods("GET")

	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(v5.FilesDir()))))

	http.ListenAndServe(":8000", r)
}
