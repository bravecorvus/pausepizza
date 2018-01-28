// Package appetizers contains the structs to store the information about the appetizers endpoints as well as any endpoint under appetizers.
package appetizers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/utils"
)

// The most parent struct of appetizer
// Compare at v5/appetizers/list.json with JSON tags at the end of each field declaration line for a better idea on how everything works
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

// Initialize() will initialize the values for an existing AppetizersList object by getting data from the respective endpoint list.json file and unmarshaling them into the struct.
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

// WriteFile() will write the current values of the AppetizersList instance that this function is operating on into the relevant list.json file
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

// Update() will reassign the pointer to the AppetizerList that this function will operate on to a new pointer passed as an argument.
// Furthermore, it will write these changes back to the JSON files.
func (elem *AppetizersList) Update(arg *AppetizersList) {

	*elem = *arg
	elem.WriteFile()

}

// Function FindFilenames() will look for a given title in its list of objects and return two strings:
//	1) The link to the normal colored image for an item
//	1) The link to the monochromatic image for an item (used to represent items that cannot be clicked on the menu in the Client Ordering App.
func (elem *AppetizersList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Appetizers {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
