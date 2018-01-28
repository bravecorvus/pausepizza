package sides

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type SaucesList struct {
	Sauces []Sauce `json:"list"`
}

type Sauce struct {
	Title                  string   `json:"title"`
	FreeIngredients        []string `json:"freeIngredients"`
	AddableIngredients     []string `json:"addableIngredients"`
	AddableIngredientTypes []string `json:"addableIngredientTypes"`
	Price                  float32  `json:"price"`
	PricePerExtra          float32  `json:"pricePerExtra"`
	Image                  Img      `json:"image"`
	Api                    string   `json:"api"`
}

func (sauces_list *SaucesList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/sides/sauces/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/sides/sauces/list.json")
	}
	err2 := json.Unmarshal(raw, &sauces_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the sauces_list")
	}
}

func (sauces_list *SaucesList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(sauces_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal sauces_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/sides/sauces/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/sides/sauces/list.json file")
	}
}

func (elem *SaucesList) Update(arg *SaucesList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *SaucesList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Sauces {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
