package ingredients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type IngredientsList struct {
	List []Ingredient `json:"list"`
}

type Ingredient struct {
	Title     string `json:"title"`
	Type      string `json:"type"`
	Available bool   `json:"available"`
	Image     Img    `json:"image"`
}

type Img struct {
	Normal     string `json:"normal"`
	Monochrome string `json:"monochrome"`
}

func (ing_list *IngredientsList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/ingredients/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/ingredients/list.json")
	}
	err2 := json.Unmarshal(raw, &ing_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the ing_list")
	}
}

func (ing_list *IngredientsList) writeFile() {
	writeFile, err1 := json.MarshalIndent(ing_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal ing_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/ingredients/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/ingredients/list.json file")
	}
}

func (ing_list *IngredientsList) Update(arg *IngredientsList) {
	*ing_list = *arg
	ing_list.writeFile()
}

func (ing_list *IngredientsList) FindFilenames(title string) (string, string) {
	for _, item := range ing_list.List {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
