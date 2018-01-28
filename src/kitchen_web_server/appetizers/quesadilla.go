package appetizers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type QuesadillaList struct {
	Quesadillas []Quesadilla `json:"list"`
}

type Quesadilla struct {
	Title                  string        `json:"title"`
	FreeIngredients        []string      `json:"freeIngredients"`
	AddableIngredients     []string      `json:"addableIngredients"`
	AddableIngredientTypes []string      `json:"addableIngredientTypes"`
	Price                  []PriceStruct `json:"price"`
	PricePerExtra          float32       `json:"pricePerExtra"`
	Image                  []ImgStruct   `json:"image"`
	Api                    string        `json:"api"`
}

func (ques_list *QuesadillaList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/appetizers/quesadilla/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/appetizers/quesadilla/list.json")
	}
	err2 := json.Unmarshal(raw, &ques_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the ques_list")
	}
}

func (ques_list *QuesadillaList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(ques_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal ques_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/appetizers/quesadilla/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/appetizers/quesadilla/list.json file")
	}
}

func (elem *QuesadillaList) Update(arg *QuesadillaList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *QuesadillaList) FindFilenames(title, parameter string) (string, string) {
	for _, item := range elem.Quesadillas {
		// fmt.Println(item.Title)
		for _, imgstruct := range item.Image {

			if item.Title == title && imgstruct.Increment == parameter {
				return utils.StripPath(imgstruct.Image.Normal), utils.StripPath(imgstruct.Image.Monochrome)
			}
		}

	}
	return "", ""
}
