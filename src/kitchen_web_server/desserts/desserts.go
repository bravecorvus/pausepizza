package desserts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type DessertsList struct {
	Desserts []Desserts `json:"list"`
}

type Desserts struct {
	Title       string `json:"title"`
	Available   bool   `json:"available"`
	Deliverable bool   `json:"deliverable"`
	Image       Img    `json:"image"`
	Api         string `json:"api"`
}

type ImgStruct struct {
	Increment string `json:"increment"`
	Image     Img    `json:"image"`
}

type Img struct {
	Normal     string `json:"normal"`
	Monochrome string `json:"monochrome"`
}

func (des_list *DessertsList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/desserts/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/desserts/list.json")
	}
	err2 := json.Unmarshal(raw, &des_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the des_list")
	}
}

func (desserts_list *DessertsList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(desserts_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal desserts_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/desserts/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/desserts/list.json file")
	}
}

type PriceStruct struct {
	Increment string  `json:"increment"`
	Price     float32 `json:"price"`
}

type PricePerExtraStruct struct {
	Increment     string  `json:"increment"`
	PricePerExtra float32 `json:"pricePerExtra"`
}

func (elem *DessertsList) Update(arg *DessertsList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *DessertsList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Desserts {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
