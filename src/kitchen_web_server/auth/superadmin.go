package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../utils"
)

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

func (super *Super) Validate(username, password string) bool {
	if super.SA.Username == username && super.SA.Password == password {
		return true
	} else {
		return false
	}
}

// func SendEmail(to, subject, message string) {
func (super *Super) Update(arg *Super) {
	*super = *arg
	super.WriteFile()
	for _, to := range super.EML {
		SendEmail(to, super.SA.Username, super.SA.Password)
	}
}
