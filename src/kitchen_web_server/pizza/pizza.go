package pizza

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type PriceStruct struct {
	Increment string  `json:"increment"`
	Price     float32 `json:"price"`
}

type ImgStruct struct {
	Increment string `json:"increment"`
	Image     Img    `json:"image"`
}

type Img struct {
	Normal     string `json:"normal"`
	Monochrome string `json:"monochrome"`
}

type PizzaList struct {
	List []PizzaItem `json:"list"`
}

type PizzaItem struct {
	Title       string `json:"title"`
	Available   bool   `json:"available"`
	Deliverable bool   `json:"deliverable"`
	Image       Img    `json:"image"`
}

func (pizza_list *PizzaList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/pizza/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/pizza/list.json")
	}
	err2 := json.Unmarshal(raw, &pizza_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the pizza_list")
	}
}

func (pizza_list *PizzaList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(pizza_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal pizza_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/pizza/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/pizza/list.json file")
	}
}

func (p *PizzaList) Update(arg *PizzaList) {

	*p = *arg
	p.WriteFile()

}

func (p *PizzaList) FindFilenames(title string) (string, string) {
	for _, item := range p.List {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
