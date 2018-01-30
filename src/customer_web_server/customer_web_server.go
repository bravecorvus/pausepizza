package main

import (
	"net/http"

	"github.com/gilgameshskytrooper/pausepizza/src/customer_web_server/v5"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/websocket"

	"github.com/gorilla/mux"
)

var (
	cache v5.ObjectStore
)

func init() {
	cache.Initialize()
}

func main() {
	r := mux.NewRouter()
	go cache.WebSocketHub.Runfunc()

	r.HandleFunc("/v5/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeClient(cache.WebSocketHub, w, r)
	})
	r.HandleFunc("/v5/{slug1}/{slug2}", cache.PostAPI).Methods("POST")
	r.HandleFunc("/v5/{slug1}", cache.PostAPI).Methods("POST")
	r.HandleFunc("/v5/{slug1}/{slug2}/{slug3}", cache.GetAPI).Methods("GET")
	r.HandleFunc("/v5/{slug1}/{slug2}", cache.GetAPI).Methods("GET")
	r.HandleFunc("/v5/{slug1}", cache.GetAPI).Methods("GET")

	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(v5.FilesDir()))))

	http.ListenAndServe(":8000", r)
}
