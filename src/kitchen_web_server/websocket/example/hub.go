package main

type Hub struct {
	clients   map[string]*Client
	broadcast chan Message
	private   chan Message
	join      chan *Client
	exit      chan *Client
}

func newHub() *Hub {
	hub := &Hub{
		clients:   make(map[string]*Client),
		broadcast: make(chan Message),
		private:   make(chan Message),
		join:      make(chan *Client),
		exit:      make(chan *Client)}

	return hub

}

func (hub *Hub) run() {
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
		// case msg := <-hub.broadcast:
		// println("[BROADCASTING] message", msg.Data)
		// for _, client := range hub.clients {
		// client.send <- msg
		// }
		case msg := <-hub.private:
			println("to", msg.To)
			hub.clients[msg.To].send <- msg
			// hub.clients[msg.From].send <- msg

		}
	}
}

type OrderStruct struct {
	Dorm    string  `json:"dorm"`
	Items   []Item  `json:"itemsOrdered"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Price   float32 `json:"price"`
	OrderID string
}

type Item struct {
	Category       string   `json:"category"`
	ExtraIncrement []string `json:"extraIncrement"`
	Increment      string   `json:"increment"`
	Item           string   `json:"item"`
}

type Token struct {
	AssociatedUser string `json:"associatedUser"`
	Token          string `json:"value"`
	Timestamp      string `json:"timestamp"`
}

func (hub *Hub) SendToUser(o OrderStruct) {
	hub.clients[o.OrderID].send <- Message{To: o.OrderID, Payload: o}
}

// func (hub *Hub) Broadcast(o OrderStruct) {
// hub.clients[o.OrderID].send <- Message{To: o.OrderID, Payload: o}
// }
