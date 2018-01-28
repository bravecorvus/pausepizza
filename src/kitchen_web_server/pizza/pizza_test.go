package pizza

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"../utils"
)

// This will check the Initialize function.
// It will manually use commands unmarshal the json list, then compare it with a PizzaList object intialized with object.Initialize()).
// Since all the other Initialize functions for other struct types. Hence, it will only be tested here.
//
//	List []PizzaItem `json:"list"`
//		Title       string `json:"title"`
//		Available   bool   `json:"available"`
//		Deliverable bool   `json:"deliverable"`
//		Image       Img    `json:"image"`
//			Normal     string `json:"normal"`
//			Monochrome string `json:"monochrome"`
func TestInitialize(t *testing.T) {

	expected := &PizzaList{}
	raw, err1 := ioutil.ReadFile(utils.TestingAssetsDir() + "v5/pizza/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/pizza/list.json")
	}
	err2 := json.Unmarshal(raw, &expected)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the test variables")
	}

	actual := &PizzaList{}
	actual.Initialize()
	for i, pizza := range actual.List {
		assert.Equal(t, expected.List[i].Available, pizza.Available, "Asserting equality of PizzaList field Available")
		assert.Equal(t, expected.List[i].Deliverable, pizza.Deliverable, "Asserting equality of PizzaList field Deliverable")
		assert.Equal(t, expected.List[i].Image, pizza.Image, "Asserting equality of PizzaList field Available Image")
		assert.Equal(t, expected.List[i].Title, pizza.Title, "Asserting equality of PizzaList field Title")
	}

}

// This will check the Update function.
// Will keep the original JSON values of a PizzaList struct to return later, and attempt to manually change values to something to check for, and then try to Update function to see if the those values got written into the JSON files by creating a fresh instance of PizzaList that got rendered from the JSON file that supposedly got these new values.
// Finally, return the values of the JSON files back to the original JSON values saved at the beginning of the function,
// Since all the other Update functions for other struct types. Hence, it will only be tested here.
func TestUpdate(t *testing.T) {
	original := &PizzaList{}
	raw, err1 := ioutil.ReadFile(utils.TestingAssetsDir() + "v5/pizza/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/pizza/list.json")
	}
	err2 := json.Unmarshal(raw, &original)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the PizzaItem")
	}

	test := &PizzaList{}
	testingdata := `{"list":[{"title":"Test1","available":true,"deliverable":true,"image":{"normal":"test1.jpg","monochrome":"test1.mono.jpg"}},{"title":"Test2","available":true,"deliverable":true,"image":{"normal":"test2.jpg","monochrome":"test2.mono.jpg"}},{"title":"Test3","available":true,"deliverable":true,"image":{"normal":"test3.jpg","monochrome":"test3.mono.jpg"}}]}`
	bytes, err := json.Marshal(testingdata)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(bytes, &test)
	test.WriteFile()
	for i, pizza := range test.List {
		assert.Equal(t, pizza.Title, "Test"+strconv.Itoa(i+1))
		assert.Equal(t, pizza.Available, true)
		assert.Equal(t, pizza.Deliverable, true)
		assert.Equal(t, pizza.Image.Normal, "test"+strconv.Itoa(i+1)+".jpg")
		assert.Equal(t, pizza.Image.Normal, "test"+strconv.Itoa(i+1)+".mono.jpg")
	}
	original.WriteFile()
}
