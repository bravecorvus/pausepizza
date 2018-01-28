package appetizers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type CheesybreadList struct {
	Cheesybreads []Cheesybread `json:"list"`
}

type Cheesybread struct {
	Title                  string   `json:"title"`
	FreeIngredients        []string `json:"freeIngredients"`
	AddableIngredients     []string `json:"addableIngredients"`
	AddableIngredientTypes []string `json:"addableIngredientTypes"`
	Price                  float32  `json:"price"`
	PricePerExtra          float32  `json:"pricePerExtra"`
	Image                  Img      `json:"image"`
	Api                    string   `json:"api"`
}

func (cheesy_list *CheesybreadList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/appetizers/cheesybread/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/appetizers/cheesybread/list.json")
	}
	err2 := json.Unmarshal(raw, &cheesy_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the cheesy_list")
	}
}

func (cheesy_list *CheesybreadList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(cheesy_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal cheesy_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/appetizers/cheesybread/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/appetizers/cheesybread/list.json file")
	}

}

func (elem *CheesybreadList) Update(arg *CheesybreadList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *CheesybreadList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Cheesybreads {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
