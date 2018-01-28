package main

import (
	"net/http"

	"./v4"
	"./v5"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v4/{slug1}/{slug2}/{slug3}/{slug4}", v4.API)
	r.HandleFunc("/v4/{slug1}/{slug2}/{slug3}", v4.API)
	r.HandleFunc("/v4/{slug1}/{slug2}", v4.API)
	r.HandleFunc("/v4/{slug1}", v4.API)

	r.HandleFunc("/v5/{slug1}/{slug2}/{slug3}", v5.API)
	r.HandleFunc("/v5/{slug1}/{slug2}", v5.API)
	r.HandleFunc("/v5/{slug1}", v5.API)

	// fs := http.FileServer(http.Dir(v5.AssetsDir() + "files/"))
	// r.Handle("/files", http.StripPrefix("/files", fs))
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(v5.FilesDir()))))

	http.ListenAndServe(":8000", r)
}
