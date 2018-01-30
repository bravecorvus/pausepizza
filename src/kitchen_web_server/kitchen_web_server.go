package main

import (
	"net/http"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/v5"

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
	// This one in particular will both check if tokens need to be removed from the list, and also change the values of the landing/set endpoint if the current time goes within or outside of the Open-Close parameters set.
	c.AddFunc("0 * * * * *", func() { cache.Times.EveryMinute(cache.Landing_List) })
	// This one changes the superadmin username and password every midnight, and emails out the new combination the email list.
	c.AddFunc("@midnight", func() { cache.SuperAdmin.EveryDay() })

	// Run every 10 seconds
	// c.AddFunc("@every 10s", func() { times.EveryMinute(&landing_list) })
	c.Start()

	cache.WebSocketHub.Run()

	r := mux.NewRouter()

	// API v5
	// WebSocket
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
