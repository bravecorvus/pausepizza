package pizza

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type SpecialtyList struct {
	List []SpecialtyItem `json:"list"`
}

type SpecialtyItem struct {
	Title                  string        `json:"title"`
	FreeIngredients        []string      `json:"freeIngredients"`
	AddableIngredients     []string      `json:"addableIngredients"`
	AddableIngredientTypes []string      `json:"addableIngredientTypes"`
	Prices                 []PriceStruct `json:"price"`
	PricePerExtra          []PriceStruct `json:"pricePerExtra"`
	Image                  []ImgStruct   `json:"image"`
}

func (specialty_list *SpecialtyList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/pizza/specialty/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/pizza/specialty/list.json")
	}
	err2 := json.Unmarshal(raw, &specialty_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the specialty_list")
	}
}

func (specialty_list *SpecialtyList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(specialty_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal specialty_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/pizza/specialty/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/pizza/specialty/list.json file")
	}
}

func (p *SpecialtyList) Update(arg *SpecialtyList) {

	*p = *arg
	p.WriteFile()

}

func (p *SpecialtyList) FindFilenames(title, parameter string) (string, string) {
	for _, item := range p.List {
		// fmt.Println(item.Title)
		for _, imgstruct := range item.Image {

			if item.Title == title && imgstruct.Increment == parameter {
				return utils.StripPath(imgstruct.Image.Normal), utils.StripPath(imgstruct.Image.Monochrome)
			}
		}

	}
	return "", ""
}
