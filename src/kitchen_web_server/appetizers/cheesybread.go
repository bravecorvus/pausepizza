package appetizers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/utils"
)

// The most parent struct of cheesybread
// Compare at v5/appetizers/cheesybread/list.json with JSON tags at the end of each field declaration line for a better idea on how everything works.
type CheesybreadList struct {
	Cheesybreads []Cheesybread `json:"list"`
}

type Cheesybread struct {
	Title                  string   `json:"title"`
	FreeIngredients        []string `json:"freeIngredients"`
	AddableIngredients     []string `json:"addableIngredients"`
	AddableIngredientTypes []string `json:"addableIngredientTypes"`
	Price                  float32  `json:"price"`
	PricePerExtra          float32  `json:"pricePerExtra"`
	Image                  Img      `json:"image"`
	Api                    string   `json:"api"`
}

// Initialize() will initialize the values for an existing CheesybreadList object by getting data from the respective endpoint list.json file and unmarshaling them into the struct.
func (cheesy_list *CheesybreadList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/appetizers/cheesybread/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/appetizers/cheesybread/list.json")
	}
	err2 := json.Unmarshal(raw, &cheesy_list)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the cheesy_list")
	}
}

// WriteFile() will write the current values of the CheesybreadList instance that this function is operating on into the relevant list.json file
func (cheesy_list *CheesybreadList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(cheesy_list, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal cheesy_list")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/appetizers/cheesybread/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/appetizers/cheesybread/list.json file")
	}

}

// Update() will reassign the pointer to the CheesybreadList that this function will operate on to a new pointer passed as an argument.
// Furthermore, it will write these changes back to the JSON files.
func (elem *CheesybreadList) Update(arg *CheesybreadList) {

	*elem = *arg
	elem.WriteFile()

}

// Function FindFilenames() will look for a given title in its list of objects and return two strings:
//	1) The link to the normal colored image for an item
//	1) The link to the monochromatic image for an item (used to represent items that cannot be clicked on the menu in the Client Ordering App.
func (elem *CheesybreadList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Cheesybreads {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
