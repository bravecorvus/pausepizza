package landing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"../auth"
	"../utils"
)

type Img struct {
	Normal     string `json:"normal"`
	Monochrome string `json:"monochrome"`
}

type Landing struct {
	PI    PauseInfo `json:"pauseinfo"`
	Items []Item    `json:"list"`
}

type PauseInfo struct {
	KitchenOpen       bool `json:"kitchenOpen"`
	DeliveryAvailable bool `json:"deliveryAvailable"`
	OvenOn            bool `json:"ovenOn"`
	LateMenu          bool `json:"lateMenu"`
}

type Item struct {
	Title string `json:"title"`
	Image Img    `json:"image"`
	Api   string `json:"api"`
}

func (landing *Landing) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/landing/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/landing/list.json")
	}
	err2 := json.Unmarshal(raw, &landing)
	// fmt.Println("Landing init", landing)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the landing")
	}
}

func (landing *Landing) WriteFile() {

	writeFile, err1 := json.MarshalIndent(landing, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal landing")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/landing/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/landing/list.json file")
	}

}

func (landing *Landing) Update(arg *Landing) {

	*landing = *arg
	landing.WriteFile()

}

func (landing *Landing) FindFilenames(title string) (string, string) {
	for _, item := range landing.Items {
		if item.Title == title {
			return utils.StripPath(item.Image.Normal), utils.StripPath(item.Image.Monochrome)
		}

	}
	return "", ""
}

type OpenClose struct {
	Open  string `json:"open"`
	Close string `json:"close"`
}

type TimesObject struct {
	Sunday    OpenClose `json:"Sunday"`
	Monday    OpenClose `json:"Monday"`
	Tuesday   OpenClose `json:"Tuesday"`
	Wednesday OpenClose `json:"Wednesday"`
	Thursday  OpenClose `json:"Thursday"`
	Friday    OpenClose `json:"Friday"`
	Saturday  OpenClose `json:"Saturday"`
}

type TimeItem struct {
	Parameter string      `json:"parameter"`
	Times     TimesObject `json:"times"`
}

type TimesItems struct {
	List []TimeItem `json:"list"`
}

func (times *TimesItems) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/landing/set/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/landing/set/list.json")
	}
	err2 := json.Unmarshal(raw, &times)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling times")
	}
	// fmt.Println("times init", times)
}

func (times *TimesItems) Update(arg *TimesItems) {

	*times = *arg
	times.writeFile()

}

func (times *TimesItems) writeFile() {
	writeFile, err1 := json.MarshalIndent(times, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal times")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/landing/set/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/landing/set/list.json file")
	}
}

// EveryMinute will get once a minute (by the c.Add() function attachment in the kitchen_web_server main program).
// First, runs the removeOld() function of the Tokens list to ensure that tokens that are older than their 24 hour time limit are deleted immediately
// This function will automatically fix the values of landing/list.json based on a minute by minute cron job that checks to see the following: Is the Pause Open? Is the Oven On? Is Delivery Available? and is the Pause serving the Late Menu
// It will check for the values stored in the landing/set/list.json file to see any of the time periods specified is met (e.g. 11:30PM on Saturday has been passed)
// If a condition is met, or becomes unmet, then landing object is modified accordingly
func (times *TimesItems) EveryMinute(landing *Landing) {

	auth.DeleteOldTokens()

	t := time.Now()
	weekday := t.Weekday().String()

	if weekday == "Sunday" {

		for _, tm := range times.List {
			parametername := tm.Parameter
			opentime := tm.Times.Sunday.Open
			closetime := tm.Times.Sunday.Close
			// fmt.Println(tm.Parameter, tm.Times.Sunday.Open, tm.Times.Sunday.Close)
			currentlyShouldBeOpen := checkIfCurrentlyEnabled(t, weekday, opentime, closetime)
			var currentlyOpen bool
			// fmt.Println("landing.PI.KitchenOpen)", landing.PI.KitchenOpen)
			// fmt.Println("parametername", parametername, "\n")
			if parametername == "kitchenOpen" {
				currentlyOpen = landing.PI.KitchenOpen
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.KitchenOpen = currentlyShouldBeOpen
					// fmt.Println("landing.PI.KitchenOpen", landing.PI.KitchenOpen)
					landing.WriteFile()
				}
			} else if parametername == "deliveryAvailable" {
				currentlyOpen = landing.PI.DeliveryAvailable
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.DeliveryAvailable = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "ovenOn" {
				currentlyOpen = landing.PI.OvenOn
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.OvenOn = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "lateMenu" {
				currentlyOpen = landing.PI.LateMenu
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.LateMenu = currentlyShouldBeOpen
					landing.WriteFile()
				}
			}

			// fmt.Println(parametername, opentime, closetime)
		}

	} else if weekday == "Monday" {

		for _, tm := range times.List {
			parametername := tm.Parameter
			opentime := tm.Times.Monday.Open
			closetime := tm.Times.Monday.Close
			// fmt.Println(tm.Parameter, tm.Times.Monday.Open, tm.Times.Monday.Close)
			currentlyShouldBeOpen := checkIfCurrentlyEnabled(t, weekday, opentime, closetime)
			var currentlyOpen bool
			// fmt.Println("landing.PI.KitchenOpen)", landing.PI.KitchenOpen)
			// fmt.Println("parametername", parametername, "\n")
			if parametername == "kitchenOpen" {
				currentlyOpen = landing.PI.KitchenOpen
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.KitchenOpen = currentlyShouldBeOpen
					// fmt.Println("landing.PI.KitchenOpen", landing.PI.KitchenOpen)
					landing.WriteFile()
				}
			} else if parametername == "deliveryAvailable" {
				currentlyOpen = landing.PI.DeliveryAvailable
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.DeliveryAvailable = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "ovenOn" {
				currentlyOpen = landing.PI.OvenOn
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.OvenOn = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "lateMenu" {
				currentlyOpen = landing.PI.LateMenu
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.LateMenu = currentlyShouldBeOpen
					landing.WriteFile()
				}
			}

			// fmt.Println(parametername, opentime, closetime)
		}

	} else if weekday == "Tuesday" {

		for _, tm := range times.List {
			parametername := tm.Parameter
			opentime := tm.Times.Tuesday.Open
			closetime := tm.Times.Tuesday.Close
			// fmt.Println(tm.Parameter, tm.Times.Tuesday.Open, tm.Times.Tuesday.Close)
			currentlyShouldBeOpen := checkIfCurrentlyEnabled(t, weekday, opentime, closetime)
			var currentlyOpen bool
			// fmt.Println("landing.PI.KitchenOpen)", landing.PI.KitchenOpen)
			// fmt.Println("parametername", parametername, "\n")
			if parametername == "kitchenOpen" {
				currentlyOpen = landing.PI.KitchenOpen
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.KitchenOpen = currentlyShouldBeOpen
					// fmt.Println("landing.PI.KitchenOpen", landing.PI.KitchenOpen)
					landing.WriteFile()
				}
			} else if parametername == "deliveryAvailable" {
				currentlyOpen = landing.PI.DeliveryAvailable
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.DeliveryAvailable = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "ovenOn" {
				currentlyOpen = landing.PI.OvenOn
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.OvenOn = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "lateMenu" {
				currentlyOpen = landing.PI.LateMenu
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.LateMenu = currentlyShouldBeOpen
					landing.WriteFile()
				}
			}

			// fmt.Println(parametername, opentime, closetime)
		}
	} else if weekday == "Wednesday" {

		for _, tm := range times.List {
			parametername := tm.Parameter
			opentime := tm.Times.Wednesday.Open
			closetime := tm.Times.Wednesday.Close
			// fmt.Println(tm.Parameter, tm.Times.Wednesday.Open, tm.Times.Wednesday.Close)
			currentlyShouldBeOpen := checkIfCurrentlyEnabled(t, weekday, opentime, closetime)
			var currentlyOpen bool
			// fmt.Println("landing.PI.KitchenOpen)", landing.PI.KitchenOpen)
			// fmt.Println("parametername", parametername, "\n")
			if parametername == "kitchenOpen" {
				currentlyOpen = landing.PI.KitchenOpen
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.KitchenOpen = currentlyShouldBeOpen
					// fmt.Println("landing.PI.KitchenOpen", landing.PI.KitchenOpen)
					landing.WriteFile()
				}
			} else if parametername == "deliveryAvailable" {
				currentlyOpen = landing.PI.DeliveryAvailable
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.DeliveryAvailable = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "ovenOn" {
				currentlyOpen = landing.PI.OvenOn
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.OvenOn = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "lateMenu" {
				currentlyOpen = landing.PI.LateMenu
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.LateMenu = currentlyShouldBeOpen
					landing.WriteFile()
				}
			}

			// fmt.Println(parametername, opentime, closetime)
		}

	} else if weekday == "Thursday" {

		for _, tm := range times.List {
			parametername := tm.Parameter
			opentime := tm.Times.Thursday.Open
			closetime := tm.Times.Thursday.Close
			// fmt.Println(tm.Parameter, tm.Times.Thursday.Open, tm.Times.Thursday.Close)
			currentlyShouldBeOpen := checkIfCurrentlyEnabled(t, weekday, opentime, closetime)
			var currentlyOpen bool
			// fmt.Println("landing.PI.KitchenOpen)", landing.PI.KitchenOpen)
			// fmt.Println("parametername", parametername, "\n")
			if parametername == "kitchenOpen" {
				currentlyOpen = landing.PI.KitchenOpen
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.KitchenOpen = currentlyShouldBeOpen
					// fmt.Println("landing.PI.KitchenOpen", landing.PI.KitchenOpen)
					landing.WriteFile()
				}
			} else if parametername == "deliveryAvailable" {
				currentlyOpen = landing.PI.DeliveryAvailable
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.DeliveryAvailable = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "ovenOn" {
				currentlyOpen = landing.PI.OvenOn
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.OvenOn = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "lateMenu" {
				currentlyOpen = landing.PI.LateMenu
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.LateMenu = currentlyShouldBeOpen
					landing.WriteFile()
				}
			}

			// fmt.Println(parametername, opentime, closetime)
		}

	} else if weekday == "Friday" {

		for _, tm := range times.List {
			parametername := tm.Parameter
			opentime := tm.Times.Friday.Open
			closetime := tm.Times.Friday.Close
			// fmt.Println(tm.Parameter, tm.Times.Friday.Open, tm.Times.Friday.Close)
			currentlyShouldBeOpen := checkIfCurrentlyEnabled(t, weekday, opentime, closetime)
			var currentlyOpen bool
			// fmt.Println("landing.PI.KitchenOpen)", landing.PI.KitchenOpen)
			// fmt.Println("parametername", parametername, "\n")
			if parametername == "kitchenOpen" {
				currentlyOpen = landing.PI.KitchenOpen
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.KitchenOpen = currentlyShouldBeOpen
					// fmt.Println("landing.PI.KitchenOpen", landing.PI.KitchenOpen)
					landing.WriteFile()
				}
			} else if parametername == "deliveryAvailable" {
				currentlyOpen = landing.PI.DeliveryAvailable
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.DeliveryAvailable = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "ovenOn" {
				currentlyOpen = landing.PI.OvenOn
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.OvenOn = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "lateMenu" {
				currentlyOpen = landing.PI.LateMenu
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.LateMenu = currentlyShouldBeOpen
					landing.WriteFile()
				}
			}

			// fmt.Println(parametername, opentime, closetime)
		}

	} else if weekday == "Saturday" {

		for _, tm := range times.List {
			parametername := tm.Parameter
			opentime := tm.Times.Saturday.Open
			closetime := tm.Times.Saturday.Close
			// fmt.Println(tm.Parameter, tm.Times.Saturday.Open, tm.Times.Saturday.Close)
			currentlyShouldBeOpen := checkIfCurrentlyEnabled(t, weekday, opentime, closetime)
			var currentlyOpen bool
			// fmt.Println("landing.PI.KitchenOpen)", landing.PI.KitchenOpen)
			// fmt.Println("parametername", parametername, "\n")
			if parametername == "kitchenOpen" {
				currentlyOpen = landing.PI.KitchenOpen
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.KitchenOpen = currentlyShouldBeOpen
					// fmt.Println("landing.PI.KitchenOpen", landing.PI.KitchenOpen)
					landing.WriteFile()
				}
			} else if parametername == "deliveryAvailable" {
				currentlyOpen = landing.PI.DeliveryAvailable
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.DeliveryAvailable = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "ovenOn" {
				currentlyOpen = landing.PI.OvenOn
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.OvenOn = currentlyShouldBeOpen
					landing.WriteFile()
				}
			} else if parametername == "lateMenu" {
				currentlyOpen = landing.PI.LateMenu
				// fmt.Println("currentlyShouldBeOpen", currentlyShouldBeOpen)
				// fmt.Println("currentlyOpen", currentlyOpen)
				if currentlyShouldBeOpen != currentlyOpen {
					landing.PI.LateMenu = currentlyShouldBeOpen
					landing.WriteFile()
				}
			}

			// fmt.Println(parametername, opentime, closetime)
		}

	}

}

func ConvertTimeStrings(arg string) string {
	split := strings.Split(arg, ":")
	ampm := split[1][2:]
	hour, min := split[0], utils.TruncateString(split[1], 2)

	hourint, err1 := strconv.Atoi(hour)
	if err1 != nil {
		fmt.Println("Couldn't convert " + hour + " into an integer")
	}
	if ampm == "PM" {
		if hourint == 12 {
			hourint = 0
		} else {
			hourint += 12
		}
	}
	return strconv.Itoa(hourint) + ":" + min + ":00"
}

// checkIfCurrentlyEnabled is a function that computes whether or not a certain parameter of of the landing.pauseinfo API endpoint meets the time criteria of current time being within operation hours, and returns that result as a bool
func checkIfCurrentlyEnabled(now time.Time, dayname, opentime, closetime string) bool {
	daynameTruncated := dayname[:3]
	month := utils.TruncateString(now.Month().String(), 3)
	var date string
	if now.Day() > 9 {
		date = strconv.Itoa(now.Day())
	} else {
		date = "0" + strconv.Itoa(now.Day())
	}
	convertedOpenTime := ConvertTimeStrings(opentime)
	convertedCloseTime := ConvertTimeStrings(closetime)
	timeZoneName, _ := now.Zone()
	year := strconv.Itoa(now.Year())

	// fmt.Println(daynameTruncated, month, convertedOpenTime, convertedCloseTime, timeZoneName, year)

	todaysOpenTime, errOpenTime := time.Parse(time.RFC1123, daynameTruncated+", "+date+" "+month+" "+year+" "+convertedOpenTime+" "+timeZoneName)
	if errOpenTime != nil {
		fmt.Println("Unable to parse the opening time of ", opentime)
	}
	todaysCloseTime, errCloseTime := time.Parse(time.RFC1123, daynameTruncated+", "+date+" "+month+" "+year+" "+convertedCloseTime+" "+timeZoneName)
	if errCloseTime != nil {
		fmt.Println("Unable to parse the opening time of ", closetime)
	}
	// fmt.Println(todaysOpenTime)
	// fmt.Println(todaysCloseTime)
	if todaysOpenTime.Sub(now).Seconds() < 0 && todaysCloseTime.Sub(now).Seconds() < 0 {
		// fmt.Println("Current time is past operating hours")
		return false
	} else if todaysOpenTime.Sub(now).Hours() > 0 && todaysCloseTime.Sub(now).Hours() > 0 {
		// fmt.Println("Current time is before operating hours")
		return false
	} else {
		// fmt.Println("Current time is during operating hours")
		return true
	}

}
