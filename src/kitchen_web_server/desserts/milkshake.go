package desserts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type MilkshakeList struct {
	Milkshakes []Milkshake `json:"list"`
}

type Milkshake struct {
	Title                  string        `json:"title"`
	FreeIngredients        []string      `json:"freeIngredients"`
	AddableIngredients     []string      `json:"addableIngredients"`
	AddableIngredientTypes []string      `json:"addableIngredientTypes"`
	Price                  []PriceStruct `json:"price"`
	PricePerExtra          float32       `json:"pricePerExtra"`
	Image                  []ImgStruct   `json:"image"`
	Api                    string        `json:"api"`
}

func (milkshake_list *MilkshakeList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/desserts/milkshake/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/desserts/milkshake/list.json")
	}
	err2 := json.Unmarshal(raw, &milkshake_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the milkshake_list")
	}
}

func (milkshake_list *MilkshakeList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(milkshake_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal milkshake_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/desserts/milkshake/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/desserts/milkshake/list.json file")
	}
}

func (elem *MilkshakeList) Update(arg *MilkshakeList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *MilkshakeList) FindFilenames(title, parameter string) (string, string) {
	for _, item := range elem.Milkshakes {
		// fmt.Println(item.Title)
		for _, imgstruct := range item.Image {

			if item.Title == title && imgstruct.Increment == parameter {
				return utils.StripPath(imgstruct.Image.Normal), utils.StripPath(imgstruct.Image.Monochrome)
			}
		}

	}
	return "", ""
}
