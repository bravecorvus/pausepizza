package appetizers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/utils"
)

// The most parent struct of chicken finger
// Compare at v5/appetizers/chickenfingers/list.json with JSON tags at the end of each field declaration line for a better idea on how everything works.
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

// Initialize() will initialize the values for an existing ChickenfingersList object by getting data from the respective endpoint list.json file and unmarshaling them into the struct.
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

// WriteFile() will write the current values of the ChickenfingersList instance that this function is operating on into the relevant list.json file
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

// Update() will reassign the pointer to the AppetizerList that this function will operate on to a new pointer passed as an argument.
// Furthermore, it will write these changes back to the JSON files.
func (elem *ChickenfingersList) Update(arg *ChickenfingersList) {

	*elem = *arg
	elem.WriteFile()

}

// Function FindFilenames() will look for a given title in its list of objects and return two strings:
//	1) The link to the normal colored image for an item
//	1) The link to the monochromatic image for an item (used to represent items that cannot be clicked on the menu in the Client Ordering App.
func (elem *ChickenfingersList) FindFilenames(title string) (string, string) {
	for _, item := range elem.Chickenfingers {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}
