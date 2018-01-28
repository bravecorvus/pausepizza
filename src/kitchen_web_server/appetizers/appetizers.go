package appetizers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type AppetizersList struct {
	Appetizers []Appetizer `json:"list"`
}

type Appetizer struct {
	Title       string `json:"title"`
	Available   bool   `json:"available"`
	Deliverable bool   `json:"deliverable"`
	Image       Img    `json:"image"`
	Api         string `json:"api"`
}

type Img struct {
	Normal     string `json:"normal"`
	Monochrome string `json:"monochrome"`
}

type ImgStruct struct {
	Increment string `json:"increment"`
	Image     Img    `json:"image"`
}

type PriceStruct struct {
	Increment string  `json:"increment"`
	Price     float32 `json:"price"`
}

type PricePerExtraStruct struct {
	Increment     string  `json:"increment"`
	PricePerExtra float32 `json:"pricePerExtra"`
}

func (app_list *AppetizersList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/appetizers/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/appetizers/list.json")
	}
	err2 := json.Unmarshal(raw, &app_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the app_list")
	}
}

func (app_list *AppetizersList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(app_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal app_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/appetizers/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/appetizers/list.json file")
	}
}

func (elem *AppetizersList) Update(arg *AppetizersList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *AppetizersList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Appetizers {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
