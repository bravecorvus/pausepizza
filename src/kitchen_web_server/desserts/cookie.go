package desserts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

type CookieList struct {
	Cookies []Cookie `json:"list"`
}

type Cookie struct {
	Title                  string        `json:"title"`
	FreeIngredients        []string      `json:"freeIngredients"`
	AddableIngredients     []string      `json:"addableIngredients"`
	AddableIngredientTypes []string      `json:"addableIngredientTypes"`
	Price                  []PriceStruct `json:"price"`
	PricePerExtra          float32       `json:"pricePerExtra"`
	Image                  []ImgStruct   `json:"image"`
	Api                    string        `json:"api"`
}

func (cookie_list *CookieList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/desserts/cookie/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/desserts/cookie/list.json")
	}
	err2 := json.Unmarshal(raw, &cookie_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the cookie_list")
	}
}

func (cookie_list *CookieList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(cookie_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal cookie_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/desserts/cookie/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/desserts/cookie/list.json file")
	}
}

func (elem *CookieList) Update(arg *CookieList) {

	*elem = *arg
	elem.WriteFile()

}

func (elem *CookieList) FindFilenames(title, parameter string) (string, string) {
	for _, item := range elem.Cookies {
		// fmt.Println(item.Title)
		for _, imgstruct := range item.Image {

			if item.Title == title && imgstruct.Increment == parameter {
				return utils.StripPath(imgstruct.Image.Normal), utils.StripPath(imgstruct.Image.Monochrome)
			}
		}

	}
	return "", ""
}
