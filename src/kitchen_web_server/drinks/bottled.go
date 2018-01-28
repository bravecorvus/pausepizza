package drinks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type BottledDrinksList struct {
	BottledDrinks []BottledDrink `json:"list"`
}

type BottledDrink struct {
	Title              string   `json:"title"`
	FreeIngredients    []string `json:"freeIngredients"`
	AddableIngredients []string `json:"addableIngredients"`
	Price              float32  `json:"price"`
	PricePerExtra      float32  `json:"pricePerExtra"`
	Image              Img      `json:"image"`
	Api                string   `json:"api"`
}

func (bottleddrink_list *BottledDrinksList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/drinks/bottled/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/drinks/bottled/list.json")
	}
	err2 := json.Unmarshal(raw, &bottleddrink_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the bottleddrink_list")
	}
}

func (bottleddrink_list *BottledDrinksList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(bottleddrink_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal bottleddrink_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/drinks/bottled/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/drinks/bottled/list.json file")
	}
}

func (elem *BottledDrinksList) Update(arg *BottledDrinksList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *BottledDrinksList) FindFilenames(title string) (string, string) {
	for _, item := range elem.BottledDrinks {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
