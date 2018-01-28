package sides

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type ChipsList struct {
	Chips []Chip `json:"list"`
}

type Chip struct {
	Title                  string   `json:"title"`
	FreeIngredients        []string `json:"freeIngredients"`
	AddableIngredients     []string `json:"addableIngredients"`
	AddableIngredientTypes []string `json:"addableIngredientTypes"`
	Price                  float32  `json:"price"`
	PricePerExtra          float32  `json:"pricePerExtra"`
	Image                  Img      `json:"image"`
	Api                    string   `json:"api"`
}

func (chips_list *ChipsList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/sides/chips/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/sides/chips/list.json")
	}
	err2 := json.Unmarshal(raw, &chips_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the chips_list")
	}
}

func (chips_list *ChipsList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(chips_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal chips_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/sides/chips/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/sides/chips/list.json file")
	}
}

func (elem *ChipsList) Update(arg *ChipsList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *ChipsList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Chips {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
