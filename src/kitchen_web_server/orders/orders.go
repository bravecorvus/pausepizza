package orders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/utils"
)

type OrderList struct {
	Orders []Order
}

type Order struct {
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
	AssociatedUser string    `json:"associatedUser"`
	Token          string    `json:"value"`
	Timestamp      time.Time `json:"timestamp"`
}

func (o *OrderList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/orders/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/pizza/orders/list.json")
	}
	err2 := json.Unmarshal(raw, &o)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the order_list")
	}
}

func (o *OrderList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(o, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal order_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/orders/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/orders/list.json file")
	}
}

func (o *OrderList) AddNewOrder(arg *Order) {
	o.Orders = append(o.Orders, *arg)
}

func (o *OrderList) RemoveOrderFromList(orderid string) {
	for i, order := range o.Orders {
		if order.OrderID == orderid {
			o.Orders = append(o.Orders[:i], o.Orders[i+1:]...)
			return
		}
	}
}

func (o *OrderList) CheckValidOrder(checkid string) bool {
	for _, order := range o.Orders {
		if order.OrderID == checkid {
			return true
		}
	}
	return false
}
