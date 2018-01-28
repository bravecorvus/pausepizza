package main

import (
	"net/http"

	"./v4"
	"./v5"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

var (
	cache v5.ObjectStore
)

func init() {

	cache.Initialize()

}

func main() {

	// Cron Jobs (Things that it makes sense for time sensitive API endpoints (i.e. the Pause Kitchen is open or closed) will be done through a globally running cron job)
	c := cron.New()

	// Run at the beginning of every minute
	c.AddFunc("0 * * * * *", func() { cache.Times.EveryMinute(cache.Landing_List) })

	// Run every 10 seconds
	// c.AddFunc("@every 10s", func() { times.EveryMinute(&landing_list) })
	c.Start()

	r := mux.NewRouter()

	// API v4
	// HTTP GET Methods
	r.HandleFunc("/v4/{slug1}/{slug2}/{slug3}", v4.GetAPI).Methods("GET")
	r.HandleFunc("/v4/{slug1}/{slug2}", v4.GetAPI).Methods("GET")
	r.HandleFunc("/v4/{slug1}", v4.GetAPI).Methods("GET")

	// HTTP POST Methods
	r.HandleFunc("/v4/{slug1}/{slug2}/{slug3}", v4.PostAPI).Methods("POST")
	r.HandleFunc("/v4/{slug1}/{slug2}", v4.PostAPI).Methods("POST")
	r.HandleFunc("/v4/{slug1}", v4.PostAPI).Methods("POST")

	// API v5
	// HTTP GET Methods
	r.HandleFunc("/v5/{slug1}/{slug2}/{slug3}", cache.PreAuthenticatedGetAPI).Methods("GET")
	r.HandleFunc("/v5/{slug1}/{slug2}", cache.PreAuthenticatedGetAPI).Methods("GET")
	r.HandleFunc("/v5/{slug1}", cache.PreAuthenticatedGetAPI).Methods("GET")

	// HTTP POST Methods
	r.HandleFunc("/v5/{slug1}/{slug2}/{slug3}/{slug4}", cache.PreAuthenticatedPostAPI).Methods("POST")
	r.HandleFunc("/v5/{slug1}/{slug2}/{slug3}", cache.PreAuthenticatedPostAPI).Methods("POST")
	r.HandleFunc("/v5/{slug1}/{slug2}", cache.PreAuthenticatedPostAPI).Methods("POST")
	r.HandleFunc("/v5/{slug1}", cache.PreAuthenticatedPostAPI).Methods("POST")

	err := http.ListenAndServe(":7000", r)
	if err != nil {
		panic(err)
	}
}
