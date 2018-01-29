package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/utils"
)

// Struct Super will contain the both the current superadmin username and password as well as the email list for those that need to receive the combination by email.
type Super struct {
	SA  SuperAdmin `json:"superadmin"`
	EML []string   `json:"emailList"`
}

type SuperAdmin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (super *Super) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/superadmin/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/superadmin/list.json")
	}
	err2 := json.Unmarshal(raw, &super)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling to the super struct")
	}
}

func (super *Super) WriteFile() {
	writeFile, err1 := json.MarshalIndent(super, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal super to file")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/superadmin/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/superadmin/list.json file")
	}
}

// Validate() will true if the username and password combination is indeed the super user
func (super *Super) Validate(username, password string) bool {
	if super.SA.Username == username && super.SA.Password == password {
		return true
	} else {
		return false
	}
}

// Update will update the superadmin struct with new values from the arguments, writing those changes to file, and then email everyone on the email list with the new username and password
func (super *SuperAdmin) Update(arg *SuperAdmin) {
	*super = *arg
}

func (super *Super) EveryDay() {
	username := generateRandomHash()
	pass := generateRandomHash()
	newSuperAdmin := SuperAdmin{Username: username, Password: pass}
	super.SA.Update(&newSuperAdmin)
	super.WriteFile()
	for _, to := range super.EML {
		SendEmail(to, super.SA.Username, super.SA.Password)
	}
}
