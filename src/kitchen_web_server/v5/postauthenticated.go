package v5

import (
	"encoding/json"
	"log"
	"net/http"

	"../appetizers"
	"../auth"
	"../desserts"
	"../drinks"
	"../ingredients"
	"../landing"
	"../photoshopjr"
	"../pizza"
	"../response"
	"../sides"
	"../utils"

	"github.com/gorilla/mux"
)

func (obj *ObjectStore) postAuthenticatedGetAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Landing
	if vars["slug2"] == "landing" {
		if vars["slug3"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/landing/list.json")
		} else if vars["slug3"] == "set" {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/landing/set/list.json")
		}

		// Ingredients
	} else if vars["slug2"] == "ingredients" {
		http.ServeFile(w, r, utils.AssetsDir()+"v5/ingredients/list.json")
	}

	// Pizza
	if vars["slug2"] == "pizza" {
		if vars["slug3"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/pizza/list.json")
		} else {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/pizza/"+vars["slug3"]+"/list.json")
		}

		// Deserts
	} else if vars["slug2"] == "desserts" {
		if vars["slug3"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/desserts/list.json")
		} else {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/desserts/"+vars["slug3"]+"/list.json")
		}

		// Appetizers
	} else if vars["slug2"] == "appetizers" {
		if vars["slug3"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/appetizers/list.json")
		} else {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/appetizers/"+vars["slug3"]+"/list.json")
		}

		// Drinks
	} else if vars["slug2"] == "drinks" {
		if vars["slug3"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/drinks/list.json")
		} else {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/drinks/"+vars["slug3"]+"/list.json")
		}

		// Sides
	} else if vars["slug2"] == "sides" {
		if vars["slug3"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/sides/list.json")
		} else {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/sides/"+vars["slug3"]+"/list.json")
		}

	} else if vars["slug2"] == "auth" {
		if vars["slug3"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/auth/list.json")
		}

	} else if vars["slug2"] == "tokens" {
		if vars["slug3"] == "" {
			http.ServeFile(w, r, utils.AssetsDir()+"v5/tokens/list.json")
		}
	}

}

func (obj *ObjectStore) postAuthenticatedPostAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Println(vars)

	// Landing
	if vars["slug2"] == "landing" {
		// fmt.Println("Landing router")
		if vars["slug3"] == "" {
			decoder := json.NewDecoder(r.Body)
			var l landing.Landing
			err := decoder.Decode(&l)
			if err != nil {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "landing page JSON API invalid"})
				log.Fatal(err)

			}
			defer r.Body.Close()
			obj.Landing_List.Update(&l)
		} else if vars["slug3"] == "set" {
			decoder := json.NewDecoder(r.Body)
			var s landing.TimesItems
			err := decoder.Decode(&s)
			if err != nil {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "landing/set page JSON API invalid"})
				log.Fatal(err)
			}
			defer r.Body.Close()
			obj.Times.Update(&s)

		} else {
			title, _ := utils.FixTitle(vars["slug3"])

			if title == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
			}

			color_filename, mono_filename := obj.Landing_List.FindFilenames(title)
			if color_filename == "" && mono_filename == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
			} else {
				photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
			}

		}

		// Ingredients
	} else if vars["slug2"] == "ingredients" {
		if vars["slug3"] == "" {
			decoder := json.NewDecoder(r.Body)
			var i ingredients.IngredientsList
			err := decoder.Decode(&i)
			if err != nil {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "ingredients page JSON API invalid"})
				log.Fatal(err)
			}
			defer r.Body.Close()
			obj.Ingredients_List.Update(&i)
		} else {
			title, _ := utils.FixTitle(vars["slug3"])
			if title == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
			}
			color_filename, mono_filename := obj.Ingredients_List.FindFilenames(title)
			if color_filename == "" && mono_filename == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
			} else {
				photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
			}
		}

		// Pizza
	} else if vars["slug2"] == "pizza" {
		if vars["slug3"] == "" {
			decoder := json.NewDecoder(r.Body)
			var p pizza.PizzaList
			err := decoder.Decode(&p)
			if err != nil {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "pizza page JSON API invalid"})
				log.Fatal(err)
			}
			defer r.Body.Close()
			obj.Pizza_List.Update(&p)

		} else if vars["slug3"] == "specialty" {
			if vars["slug4"] == "" {
				decoder := json.NewDecoder(r.Body)
				var p pizza.SpecialtyList
				err := decoder.Decode(&p)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "pizza/specialty page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Pizza_Specialty_List.Update(&p)

			} else {
				title, param := utils.FixTitle(vars["slug4"])
				if title == "" || param == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug4"]})
				}
				// fmt.Println("title", title)
				// fmt.Println("param", param)
				color_filename, mono_filename := obj.Pizza_Specialty_List.FindFilenames(title, param)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug4"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else if vars["slug3"] == "build" {
			if vars["slug4"] == "" {
				decoder := json.NewDecoder(r.Body)
				var p pizza.BuildList
				err := decoder.Decode(&p)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "pizza/build page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Pizza_Build_List.Update(&p)

			} else {
				title, param := utils.FixTitle(vars["slug4"])
				if title == "" || param == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug4"]})
				}
				color_filename, mono_filename := obj.Pizza_Build_List.FindFilenames(title, param)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug4"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else {
			title, _ := utils.FixTitle(vars["slug3"])
			if title == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
			}
			color_filename, mono_filename := obj.Pizza_List.FindFilenames(title)
			if color_filename == "" && mono_filename == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
			} else {
				photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
			}
		}

		// Deserts
	} else if vars["slug2"] == "desserts" {

		if vars["slug3"] == "" {
			decoder := json.NewDecoder(r.Body)
			var d desserts.DessertsList
			err := decoder.Decode(&d)
			if err != nil {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "desserts page JSON API invalid"})
				log.Fatal(err)
			}
			defer r.Body.Close()
			obj.Desserts_List.Update(&d)

		} else if vars["slug3"] == "cookie" {
			if vars["slug4"] == "" {

				decoder := json.NewDecoder(r.Body)
				var d desserts.CookieList
				err := decoder.Decode(&d)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "desserts/cookie page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Cookie_List.Update(&d)

			} else {
				title, param := utils.FixTitle(vars["slug4"])
				if title == "" || param == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug4"]})
				}
				color_filename, mono_filename := obj.Cookie_List.FindFilenames(title, param)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug4"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else if vars["slug3"] == "icecream" {
			if vars["slug4"] == "" {
				decoder := json.NewDecoder(r.Body)
				var d desserts.IceCreamList
				err := decoder.Decode(&d)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "desserts/icecream page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Icecream_List.Update(&d)

			} else {
				title, param := utils.FixTitle(vars["slug4"])
				if title == "" || param == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug4"]})
				}
				color_filename, mono_filename := obj.Icecream_List.FindFilenames(title, param)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug4"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else if vars["slug3"] == "milkshake" {
			if vars["slug4"] == "" {
				decoder := json.NewDecoder(r.Body)
				var d desserts.MilkshakeList
				err := decoder.Decode(&d)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "desserts/milkshake page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Milkshake_List.Update(&d)

			} else {
				title, param := utils.FixTitle(vars["slug4"])
				if title == "" || param == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug4"]})
				}
				color_filename, mono_filename := obj.Milkshake_List.FindFilenames(title, param)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug4"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else {
			title, _ := utils.FixTitle(vars["slug3"])
			if title == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
			}
			color_filename, mono_filename := obj.Desserts_List.FindFilenames(title)
			if color_filename == "" && mono_filename == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
			} else {
				photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
			}

		}

		// Appetizers
	} else if vars["slug2"] == "appetizers" {

		if vars["slug3"] == "" {
			decoder := json.NewDecoder(r.Body)
			var a appetizers.AppetizersList
			err := decoder.Decode(&a)
			if err != nil {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "appetizers page JSON API invalid"})
				log.Fatal(err)
			}
			defer r.Body.Close()
			obj.Appetizers_List.Update(&a)

		} else if vars["slug3"] == "cheesybread" {
			if vars["slug4"] == "" {
				decoder := json.NewDecoder(r.Body)
				var a appetizers.CheesybreadList
				err := decoder.Decode(&a)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "appetizers/cheesybread page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Cheesybread_List.Update(&a)
			} else {
				title, _ := utils.FixTitle(vars["slug4"])
				if title == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug4"]})
				}
				color_filename, mono_filename := obj.Cheesybread_List.FindFilenames(title)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug4"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else if vars["slug3"] == "chickenfingers" {
			if vars["slug4"] == "" {
				decoder := json.NewDecoder(r.Body)
				var a appetizers.ChickenfingersList
				err := decoder.Decode(&a)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "appetizers/chickenfingers page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Chickenfingers_List.Update(&a)
			} else {
				title, _ := utils.FixTitle(vars["slug4"])
				if title == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug4"]})
				}
				color_filename, mono_filename := obj.Chickenfingers_List.FindFilenames(title)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else if vars["slug3"] == "quesadilla" {
			if vars["slug4"] == "" {
				decoder := json.NewDecoder(r.Body)
				var a appetizers.QuesadillaList
				err := decoder.Decode(&a)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "appetizers/quesadilla page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Quesadilla_List.Update(&a)
			} else {
				title, param := utils.FixTitle(vars["slug4"])
				if title == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug4"]})
				}
				color_filename, mono_filename := obj.Quesadilla_List.FindFilenames(title, param)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else {
			title, _ := utils.FixTitle(vars["slug3"])
			if title == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
			}
			color_filename, mono_filename := obj.Appetizers_List.FindFilenames(title)
			if color_filename == "" && mono_filename == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
			} else {
				photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
			}
		}

		// Drinks
	} else if vars["slug2"] == "drinks" {

		if vars["slug3"] == "" {

			decoder := json.NewDecoder(r.Body)
			var d drinks.DrinksList
			err := decoder.Decode(&d)
			if err != nil {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "drinks page JSON API invalid"})
				log.Fatal(err)
			}
			defer r.Body.Close()
			obj.Drinks_List.Update(&d)

		} else if vars["slug3"] == "bottled" {
			if vars["slug4"] == "" {

				decoder := json.NewDecoder(r.Body)
				var d drinks.BottledDrinksList
				err := decoder.Decode(&d)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "drinks/bottled page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Bottled_Drinks_List.Update(&d)

			} else {

				title, _ := utils.FixTitle(vars["slug4"])
				if title == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
				}
				color_filename, mono_filename := obj.Bottled_Drinks_List.FindFilenames(title)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else if vars["slug3"] == "fountain" {
			if vars["slug4"] == "" {

				decoder := json.NewDecoder(r.Body)
				var d drinks.FountainDrinksList
				err := decoder.Decode(&d)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "drinks/fountain page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Fountain_Drinks_List.Update(&d)

			} else {

				title, _ := utils.FixTitle(vars["slug4"])
				if title == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
				}
				color_filename, mono_filename := obj.Fountain_Drinks_List.FindFilenames(title)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else {

			title, _ := utils.FixTitle(vars["slug3"])
			if title == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
			}
			color_filename, mono_filename := obj.Drinks_List.FindFilenames(title)
			if color_filename == "" && mono_filename == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
			} else {
				photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
			}

		}

		// Sides
	} else if vars["slug2"] == "sides" {
		if vars["slug3"] == "" {

			decoder := json.NewDecoder(r.Body)
			var s sides.SidesList
			err := decoder.Decode(&s)
			if err != nil {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "sides page JSON API invalid"})
				log.Fatal(err)
			}
			defer r.Body.Close()
			obj.Sides_List.Update(&s)

		} else if vars["slug3"] == "chips" {

			if vars["slug4"] == "" {

				decoder := json.NewDecoder(r.Body)
				var s sides.ChipsList
				err := decoder.Decode(&s)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "sides/chips page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()

				obj.Chips_List.Update(&s)
			} else {

				title, _ := utils.FixTitle(vars["slug4"])
				if title == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
				}
				color_filename, mono_filename := obj.Chips_List.FindFilenames(title)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}

			}

		} else if vars["slug3"] == "sauces" {
			if vars["slug4"] == "" {

				decoder := json.NewDecoder(r.Body)
				var s sides.SaucesList
				err := decoder.Decode(&s)
				if err != nil {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "sides/sauces page JSON API invalid"})
					log.Fatal(err)
				}
				defer r.Body.Close()
				obj.Sauces_List.Update(&s)

			} else {

				title, _ := utils.FixTitle(vars["slug4"])
				if title == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
				}
				color_filename, mono_filename := obj.Sauces_List.FindFilenames(title)
				if color_filename == "" && mono_filename == "" {
					json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
				} else {
					photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
				}
			}

		} else {
			title, _ := utils.FixTitle(vars["slug3"])
			if title == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "No such item for " + vars["slug3"]})
			}
			color_filename, mono_filename := obj.Sides_List.FindFilenames(title)
			if color_filename == "" && mono_filename == "" {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "Image link could not be found for " + vars["slug3"]})
			} else {
				photoshopjr.ProcessImage(w, r, color_filename, mono_filename)
			}
		}

	} else if vars["slug2"] == "auth" {
		if vars["slug3"] == "" {

			decoder := json.NewDecoder(r.Body)
			var a auth.ValidUsersList
			err := decoder.Decode(&a)
			if err != nil {
				json.NewEncoder(w).Encode(response.Response{Status: false, Message: "auth page JSON API invalid"})
				log.Fatal(err)
			}
			defer r.Body.Close()
			obj.Admins.Update(&a)

		}
	}

}
