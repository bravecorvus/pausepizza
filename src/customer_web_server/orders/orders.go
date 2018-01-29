package orders

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
