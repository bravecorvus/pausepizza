package sides

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type SidesList struct {
	Sides []Side `json:"list"`
}

type Side struct {
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

func (sides_list *SidesList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/sides/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/sides/list.json")
	}
	err2 := json.Unmarshal(raw, &sides_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the sides_list")
	}
}

func (sides_list *SidesList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(sides_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal sides_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/sides/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/sides/list.json file")
	}
}

func (elem *SidesList) Update(arg *SidesList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *SidesList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Sides {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
