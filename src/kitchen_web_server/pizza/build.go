package pizza

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type BuildList struct {
	List []BuildItem `json:"list"`
}

type BuildItem struct {
	Title                  string        `json:"title"`
	FreeIngredients        []string      `json:"freeIngredients"`
	AddableIngredients     []string      `json:"addableIngredients"`
	AddableIngredientTypes []string      `json:"addableIngredientTypes"`
	Prices                 []PriceStruct `json:"price"`
	PricePerExtra          []PriceStruct `json:"pricePerExtra"`
	Image                  []ImgStruct   `json:"image"`
}

func (build_list *BuildList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/pizza/build/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/pizza/build/list.json")
	}
	err2 := json.Unmarshal(raw, &build_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the build_list")
	}
}

func (build_list *BuildList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(build_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal build_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/pizza/build/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/pizza/build/list.json file")
	}
}

func (p *BuildList) Update(arg *BuildList) {

	*p = *arg
	p.WriteFile()

}

func (p *BuildList) FindFilenames(title, parameter string) (string, string) {
	for _, item := range p.List {
		for _, imgstruct := range item.Image {

			if item.Title == title && imgstruct.Increment == parameter {
				return utils.StripPath(imgstruct.Image.Normal), utils.StripPath(imgstruct.Image.Monochrome)
			}
		}

	}
	return "", ""
}
