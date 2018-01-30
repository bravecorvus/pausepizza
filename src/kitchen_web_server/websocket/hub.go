package websocket

import "github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/orders"

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
