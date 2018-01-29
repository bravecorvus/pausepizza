package v5

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gilgameshskytrooper/pausepizza/src/customer_web_server/orders"
	"github.com/gilgameshskytrooper/pausepizza/src/customer_web_server/utils"
	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/auth"
	"github.com/gorilla/mux"
)

type response_struct struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	OrderID string `json:"orderID"`
}

type kitchen_response_struct struct {
	status  bool   `json:"Status"`
	message string `json:"Message"`
}

func Pwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir + "/"
}

func FilesDir() string {
	return AssetsDir() + "files/"
}

func AssetsDir() string {
	return Pwd() + "assets/"
}

func GetAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Landing
	if vars["slug1"] == "landing" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/landing/list.json")
		}

		// Ingredients
	} else if vars["slug1"] == "ingredients" {
		http.ServeFile(w, r, AssetsDir()+"v5/ingredients/list.json")

		// Pizza
	} else if vars["slug1"] == "pizza" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/pizza/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/pizza/"+vars["slug2"]+"/list.json")
		}

		// Deserts
	} else if vars["slug1"] == "desserts" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/desserts/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/desserts/"+vars["slug2"]+"/list.json")
		}

		// Appetizers
	} else if vars["slug1"] == "appetizers" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/appetizers/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/appetizers/"+vars["slug2"]+"/list.json")
		}

		// Drinks
	} else if vars["slug1"] == "drinks" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/drinks/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/drinks/"+vars["slug2"]+"/list.json")
		}

		// Sides
	} else if vars["slug1"] == "sides" {
		if vars["slug2"] == "" {
			http.ServeFile(w, r, AssetsDir()+"v5/sides/list.json")
		} else {
			http.ServeFile(w, r, AssetsDir()+"v5/sides/"+vars["slug2"]+"/list.json")
		}

	}

}

func PostAPI(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	if vars["slug1"] == "checkout" {

		decoder := json.NewDecoder(r.Body)
		var o orders.OrderStruct
		err := decoder.Decode(&o)
		if err != nil {
			json.NewEncoder(w).Encode(response_struct{Status: false, Message: "checkout JSON API invalid", OrderID: ""})
			fmt.Println(err.Error())

		}
		o.OrderID = utils.GenerateRandomHash()
		marshaled, err2 := json.Marshal(o)
		if err2 != nil {
			fmt.Println("Couldn't marshall Order back to JSON []byte")
		}

		var super auth.Super
		raw, err3 := ioutil.ReadFile(Pwd() + "assets/v5/superadmin/list.json")
		if err3 != nil {
			fmt.Println("Could not open v5/superadmin/list.json")
		}
		err4 := json.Unmarshal(raw, &super)
		if err4 != nil {
			fmt.Println("Trouble unmarshalling the token list")
		}

		adminlogin, err5 := json.Marshal(super.SA)
		if err5 != nil {
			fmt.Println("Trouble marshalling just the admin login struct (i.e. just the username and password field)")
		}
		body := bytes.NewReader(adminlogin)
		req, err6 := http.NewRequest("POST", "http://localhost:7000/v5/login", body)
		if err6 != nil {
			fmt.Println("Couldn't construct the post to the login endpoint to get the superadmin token")
		}
		req.Header.Set("Content-Type", "application/json")
		resp1, err7 := http.DefaultClient.Do(req)
		if err7 != nil {
			fmt.Println("couldn't do POST request to kitchen management login endpoint")
		}
		defer resp1.Body.Close()
		var token orders.Token
		body1, err8 := ioutil.ReadAll(resp1.Body)
		if err8 != nil {
			fmt.Println("Couldn't read response from login endpoint")
		}
		err9 := json.Unmarshal(body1, &token)
		if err9 != nil {
			fmt.Println("Trouble unmarshaling the token received from the login endpoint")
		}

		resp2, err10 := http.Post("http://localhost:7000/v5/"+token.Token+"/neworder", "application/json", bytes.NewBuffer(marshaled))
		if err10 != nil {
			fmt.Println("Can't post to neworder endpoint")
		}
		defer r.Body.Close()
		var server_resp kitchen_response_struct
		body2, err11 := ioutil.ReadAll(resp2.Body)
		if err11 != nil {
			fmt.Println("Couldn't read kitchen server to client ordering app response to add new order")
		}
		err12 := json.Unmarshal(body2, &server_resp)
		if err12 != nil {
			fmt.Println("Couldn't unmarshal kitchen server to client response")
		}
		json.NewEncoder(w).Encode(response_struct{Status: true, Message: "Order sucessfully registered", OrderID: o.OrderID})

	} else if vars["slug1"] == "ordercomplete" {
		orderid := vars["slug2"]
		// fmt.Println(orderid)
		fmt.Println("Order Complete for order " + orderid)
		json.NewEncoder(w).Encode(response_struct{Status: true, Message: "Order fulfilled POST request received"})
	}
}
