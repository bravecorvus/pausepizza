package websocket

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/orders"
)

type Hub struct {
	clients   map[string]*Client
	broadcast chan Message
	private   chan Message
	join      chan *Client
	exit      chan *Client
}

func NewHub() *Hub {
	hub := &Hub{
		clients:   make(map[string]*Client),
		broadcast: make(chan Message),
		private:   make(chan Message),
		join:      make(chan *Client),
		exit:      make(chan *Client)}

	return hub

}

func (h *Hub) Initialize() {
	h = &Hub{
		clients:   make(map[string]*Client),
		broadcast: make(chan Message),
		private:   make(chan Message),
		join:      make(chan *Client),
		exit:      make(chan *Client)}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.join:
			if _, ok := hub.clients[client.id]; ok {
				println("Duplicated id!")
			} else {
				hub.clients[client.id] = client
			}
		case client := <-hub.exit:
			println(client)
		case msg := <-hub.broadcast:
			// println("[BROADCASTING] message", msg.Data)
			for _, client := range hub.clients {
				client.send <- msg
			}
		case msg := <-hub.private:
			println("to", msg.To)
			hub.clients[msg.To].send <- msg
			// hub.clients[msg.From].send <- msg

		}
	}
}

func (hub *Hub) SendToUser(o orders.Order) {
	hub.clients[o.OrderID].send <- Message{To: o.OrderID, Type: "private", Payload: o}
}

func (hub *Hub) Broadcast(o orders.Order) {
	hub.clients[o.OrderID].send <- Message{To: o.OrderID, Type: "broadcast", Payload: o}
}

func (h *Hub) Runfunc() {
	time.Sleep(10 * time.Second)
	var o orders.Order
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
