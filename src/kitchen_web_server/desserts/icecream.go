package desserts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type IceCreamList struct {
	IceCreams []IceCream `json:"list"`
}

type IceCream struct {
	Title              string                `json:"title"`
	FreeIngredients    []string              `json:"freeIngredients"`
	AddableIngredients []string              `json:"addableIngredients"`
	Price              []PriceStruct         `json:"price"`
	PricePerExtra      []PricePerExtraStruct `json:"pricePerExtra"`
	Image              []ImgStruct           `json:"image"`
	Api                string                `json:"api"`
}

func (icecream_list *IceCreamList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/desserts/icecream/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/desserts/icecream/list.json")
	}
	err2 := json.Unmarshal(raw, &icecream_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the icecream_list")
	}
}

func (icecream_list *IceCreamList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(icecream_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal icecream_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/desserts/icecream/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/desserts/icecream/list.json file")
	}
}

func (elem *IceCreamList) Update(arg *IceCreamList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *IceCreamList) FindFilenames(title, parameter string) (string, string) {
	for _, item := range elem.IceCreams {
		// fmt.Println(item.Title)
		for _, imgstruct := range item.Image {

			if item.Title == title && imgstruct.Increment == parameter {
				return utils.StripPath(imgstruct.Image.Normal), utils.StripPath(imgstruct.Image.Monochrome)
			}
		}

	}
	return "", ""
}
