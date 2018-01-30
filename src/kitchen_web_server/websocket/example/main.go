package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (h *Hub) runfunc() {
	time.Sleep(30 * time.Second)
	var o OrderStruct
	bytes := []byte(`{"dorm":"Thorson","itemsOrdered":[{"category":"Pizza","extraIncrement":["Chicken","Bacon"],"increment":"Large","item":"Build Your Own Pizza"}],"name":"Deepak","phone":"55566677777","price":11.5,"OrderID":"5PSW9QjylBRDjIEdJlKkrOrYFJmxQ2nF2BHASr3x"}`)
	err1 := json.Unmarshal(bytes, &o)
	if err1 != nil {
		fmt.Println("Cant decode o struct")
	}
	fmt.Println(o)
	h.SendToUser(o)
	time.Sleep(3 * time.Second)
	h.SendToUser(o)
	time.Sleep(3 * time.Second)
	h.SendToUser(o)
}

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveClient(hub, w, r)
	})

	if err := http.ListenAndServe(":2002", nil); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
