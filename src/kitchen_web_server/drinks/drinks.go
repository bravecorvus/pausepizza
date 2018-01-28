package drinks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type DrinksList struct {
	Drinks []Drink `json:"list"`
}

type Drink struct {
	Title       string `json:"title"`
	Available   bool   `json:"available"`
	Deliverable bool   `json:"deliverable"`
	Image       Img    `json:"image"`
	Api         string `json:"api"`
}

func (drink_list *DrinksList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/drinks/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/drinks/list.json")
	}
	err2 := json.Unmarshal(raw, &drink_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the drink_list")
	}
}

func (drink_list *DrinksList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(drink_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal drink_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/drinks/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/drinks/list.json file")
	}
}

type Img struct {
	Normal     string `json:"normal"`
	Monochrome string `json:"monochrome"`
}

func (elem *DrinksList) Update(arg *DrinksList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *DrinksList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Drinks {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
