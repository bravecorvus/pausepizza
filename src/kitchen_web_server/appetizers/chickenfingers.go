package appetizers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type ChickenfingersList struct {
	Chickenfingers []Chickenfinger `json:"list"`
}

type Chickenfinger struct {
	Title       string `json:"title"`
	Available   bool   `json:"available"`
	Deliverable bool   `json:"deliverable"`
	Image       Img    `json:"image"`
	Api         string `json:"api"`
}

func (chickenfingers_list *ChickenfingersList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/appetizers/chickenfingers/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/appetizers/chickenfingers/list.json")
	}
	err2 := json.Unmarshal(raw, &chickenfingers_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the app_list")
	}
}

func (chickenfingers_list *ChickenfingersList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(chickenfingers_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal chickenfingers_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/appetizers/chickenfingers/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/appetizers/chickenfingers/list.json file")
	}
}

func (elem *ChickenfingersList) Update(arg *ChickenfingersList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *ChickenfingersList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Chickenfingers {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
