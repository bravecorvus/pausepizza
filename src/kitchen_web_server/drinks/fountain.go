package drinks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type FountainDrinksList struct {
	FountainDrinks []FountainDrink `json:"list"`
}

type FountainDrink struct {
	Title              string   `json:"title"`
	FreeIngredients    []string `json:"freeIngredients"`
	AddableIngredients []string `json:"addableIngredients"`
	Price              float32  `json:"price"`
	PricePerExtra      float32  `json:"pricePerExtra"`
	Image              Img      `json:"image"`
	Api                string   `json:"api"`
}

func (fountaindrink_list *FountainDrinksList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/drinks/fountain/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/drinks/fountain/list.json")
	}
	err2 := json.Unmarshal(raw, &fountaindrink_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the fountaindrink_list")
	}
}

func (fountaindrink_list *FountainDrinksList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(fountaindrink_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal fountaindrink_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/drinks/fountain/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/drinks/fountain/list.json file")
	}
}

func (elem *FountainDrinksList) Update(arg *FountainDrinksList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *FountainDrinksList) FindFilenames(title string) (string, string) {
	for _, item := range elem.FountainDrinks {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
