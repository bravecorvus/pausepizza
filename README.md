# Pause Pizza Backend Implementations
### Andrew Lee

---
## Introduction

**Please view [Wiki](https://github.com/gilgameshskytrooper/pausepizza/wiki) in order to see implement system**
The Kitchen Web Server (which does most of the heavy lifting) has been heavily documented. If you have any questions about code specifically, head over to the project [godocs page](https://godoc.org/github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server).

### Quick charts of the packages that make up the server

#### [Customer Ordering App](https://github.com/gilgameshskytrooper/pausepizza/tree/master/src/customer_web_server)

| Package | Summary | Link |
| -- | -- | -- |
| main | main function which drives the application | [link](https://github.com/gilgameshskytrooper/pausepizza/blob/master/src/customer_web_server/customer_web_server.go) |
| utils | useful functions used in many places of the application | [link](https://github.com/gilgameshskytrooper/pausepizza/tree/master/src/customer_web_server/utils) |
| orders | package which stores the specific structs used to store the JSON objects within the running app | [link](https://github.com/gilgameshskytrooper/pausepizza/blob/master/src/customer_web_server/orders/orders.go) |
| v5 | router version 5 defines the routes of the entire application | [link](https://github.com/gilgameshskytrooper/pausepizza/tree/master/src/customer_web_server/v5) |

#### [Kitchen Management App](https://github.com/gilgameshskytrooper/pausepizza/tree/master/src/kitchen_web_server)

| Package | Summary | Link |
| -- | -- | -- |
| main | main function which drives the application | [link](https://github.com/gilgameshskytrooper/pausepizza/blob/master/src/kitchen_web_server/kitchen_web_server.go) |
| auth | package which contains most of the authentication related structs ValidUsersList and TokenList as well as the related functions. | [link](https://github.com/gilgameshskytrooper/pausepizza/tree/master/src/kitchen_web_server/auth) |
| appetizers/desserts/drinks/ingredients/landing/orders/pizza/sides | these packages all define all the structs used to hold an in program copy of what is contained in the JSON files as well as the related functions such as initializing them, writing them | [link](https://github.com/gilgameshskytrooper/pausepizza/tree/master/src/kitchen_web_server) |
| utils | contains many functions which are generic but are useful in many other packages | [link](https://github.com/gilgameshskytrooper/pausepizza/blob/master/src/kitchen_web_server/utils/utils.go) |
| photoshopjr | package used to provide server-side image processing | [link](https://github.com/gilgameshskytrooper/pausepizza/blob/master/src/kitchen_web_server/photoshopjr/photoshopjr.go) |
| v5 | router version 5 defines the routes of the entire application as well as the Cache struct which is a parent struct to all other JSON structs (appetizers/desserts/drinks/ingredients/landing/orders/pizza/sides); it also contains, more complex endpoints due to the secure nature of the kitchen management app | [link](https://github.com/gilgameshskytrooper/pausepizza/tree/master/src/kitchen_web_server/v5) |

### Summary

To summarize what these two servers do, the client ordering primary job to provide RESTful service by tying specific API endpoints to specific JSON files. These files are used by the front-end (written in React Native) to render the application. Furthermore, it also hosts other non-REST related static assets such as images. Finally, the client ordering app has exactly one `POST` endpoint to submit new orders.

The kitchen management app fulfills a simlar role in passing JSON objects to the kitchen management front-end, but contains a `POST` endpoint for every single `GET` endpoint (to be used by workers of the Pause Pizza Kitchen in order to add, modify, or remove any items on the menu). To keep things simple, the `POST` endpoints accepts a JSON objects that are identical in structure to the one that is returned on a `GET` request.

Although most production systems avoid storing persistent data into plaintext, I believe there was a key architectural pattern difference makes a database structure unsuited for our application.

With our application, the amount of times that JSON objects are retrived via `GET` requests would far outnumber the number of any `POST` requests by kitchen managers. Hence, the fact that databases mitigate the pains of having a large number of writes would not benefit these services in any significant way, and instead, yield poorer results for all `GET` requests (as instead of serving up static JSON files that are already stored on disk, both servers would have to consult with a database, serialize the data in an intermediate structure [i.e. JSON], and then send it). Hence, the stateful storage of both of these applications is actually directly stored in static JSON files rather than in a database.

---

## Client Ordering Server
The client ordering API is available at the url `https://customer.pause.pizza/v5/...`.

The API server returns STATEful resource (JSON files) which will be parsed by the React Native front-end, and converted into objects to populate the app.

> Examples:

When the user first opens, the app, the front-end will request the endpoint `https://customer.pause.pizza/v5/landing`, which returns the following JSON object

```json
{
	"pauseinfo": {
		"kitchenOpen": true,
		"deliveryAvailable": true,
		"ovenOn": true,
		"lateMenu": false
	},
	"list": [
		{
			"title": "Pizza",
			"image": {
				"normal": "",
				"monochrome": ""
			},
			"api": "v5/pizza"
		},
		{
			"title": "Dessert",
			"image": {
				"normal": "",
				"monochrome": ""
			},
			"api": "v5/desserts"
		},
		{
			"title": "Appetizers",
			"image": {
				"normal": "",
				"monochrome": ""
			},
			"api": "v5/appetizers"
		},
		{
			"title": "Drinks",
			"image": {
				"normal": "",
				"monochrome": ""
			},
			"api": "v5/drinks"
		},
		{
			"title": "Sides",
			"image": {
				"normal": "",
				"monochrome": ""
			},
			"api": "v5/sides"
		}
	]
}
```

the pauseinfo object towards the top defines the parameter of what gets rendered at the very beginning (i.e. If `kitchenOpen==false`, then nothing form the app will run, or if `deliveryAvailable==false`, then delivery items will be turned off)

As you can see, each subsequent endpoint is defined as part of the parent API resource

For example, in order to get the resources of the Pizza endpoint, we will go to `https://customer.pause.pizza/v5/pizza`


> Some more links
`https://customer.pause.pizza/v5/pizza/specialty`
`https://customer.pause.pizza/v5/desserts`
`https://customer.pause.pizza/v5/appetizers/cheesybread`


---

## Management App

The management app manipulates the data of the Client Ordering application to reflect changes in the Pause such as the addition of new menu items, the deletion of certain items, running out of food etc. Despite the fact that the the kitchen web server and the client ordering web server share the same JSON files, the kitchen server contains `GET` endpoints that are not present on the client side.

Unlike the client ordering app, the management app needs to be secure. However, sending the username and password every time to request a resource in a header like this `curl -H "Content-Type: application/json" -X POST -d '{"username":"leeas","password":"abcd"}' https://kitchen.pause.pizza/v5/pizza` is not ideal as it is transfering sensitive data.

Therefore, I designed a token based system where the user *initially* authenticates using a username and password, gets the server to randomly generate a 40 character alphanumeric string as a secure token (which is given a expiration time), and can only access points of the app using this secure token.

> Example
The kitchen management app API is located at `https://kitchen.pause.pizza/v5/...`

However, trying to access the endpoints in the same way we did in the client app like this `https://kitchen.pause.pizza/v5/landing` will return the following message:

```json
{"status":false,"message":"User did not access the link with a valid token. Please log in"}
```

Hence, the user must first log in using a username and password saved on the server side.

Now, to test the authentication:

We will use curl to attempt to receive a token using a non-valid username and password.

```
curl -H "Content-Type: application/json" -X POST -d '{"username":"notarealaccount","password":"notarealpassword"}' https://kitchen.pause.pizza/v5/login
```

The example above, is showing that only authorized username+password combinations will work. It will yield the response

```json
{"status":false,"message":"Username and Password combination does not exist. If you forgot your password, please talk to another manager for the app "}
```

However, a real user with a real password will be able to receive a secure token
```
curl -H "Content-Type: application/json" -X POST -d '{"username":"leeas","password":"abcd"}' https://kitchen.pause.pizza/v5/login
```

Which will return a JSON object similar to the following:
```json
{"associatedUser":"leeas","value":"N95RxRTZHWUsaD6HEdz0ThbXfQ6pYSQ3n267l1VQ","timestamp":"2018-01-24T13:11:23.309813635-06:00"}
```

Then, we use the value of `value` in the JSON object that was returned to us by the server as the second slug of all request (both `GET` and `POST`)

```
https://kitchen.pause.pizza/v5/N95RxRTZHWUsaD6HEdz0ThbXfQ6pYSQ3n267l1VQ/landing
```

***Note, the above example will not work verbatum since the token expires every 24 hours (Which is what we want)*** .


Server also runs a cron function to change the values of the `landing` endpoint of the client app according to the values stored at the `v5/landing/set/list.json` object.

```
curl https://kitchen.pause.pizza/v5/[token]/landing/set
```


## Ordering Process
The following is a step-by-step example of the lifecycle of a client's interaction with the application.

1. Client opens app which renders all possible order items from the `GET` endpoints from the client ordering web server
2. Client selects items which get stored in the front-ends application state as a JSON object.
3. When client submits order, it `POSTS` a request from client ordering app including Items list, cost, name of person, phone number, delivery location etc. to the `v5/checkout` endpoint to the client web server.
4. client ordering backend receives order. The backend will create a instance of a struct with the above information, and add a new order ID based on a randomly generated 40 character string.
5. client ordering backend will find the Super Admin username from the `assets/v5/superadmin/list.json`, and then try to find a token in `assets/v5/tokens/list.json` that matches the user name. If the token exists, the client ordering Backend will attempt to `POST` the order (including the new order ID generated during step 2) at the `v5/neworder` endpoint on the Kitchen management using the generated token. If the token does not exist, the client ordering backend will use the Super Admin username and password to generate a new token and go through the same exact steps to post to the kitchen management app.
6. Kitchen management app registers new order, and adds it to the queue. It will send out a server to client notification regarding the addition of the new order (in order to dynamically modify the UI if the orders page is being accessed at the moment) through a websocket (this time, all users of the kitchen management app will get a notification).
7. When the order gets processed on the kitchen management backend, the kitchen app will respond with a JSON struct (with a status of `success`). If something fails, it will respond with a status of `fail`.
8. If the above process registers a success, the client server will return an object also with a status = `success` portion as well as the generated orderID
9. The client ordering app front-end will then attempt to subsribe to a websocket using the endpoint `v5/websocket/[generated order id]` (Not currently implemented).
10. When the Kitchen sends an order complete `POST` to the kitchen server backend, the kitchen backend will `POST` to the `v5/ordercomplete` of the client ordering backend. Then the client ordering backend will send a notification to the specific user who sent the order whenever it is done.

> Demo
This is the full JSON object we will be sending to the client backend
```json
{
	"dorm": "Mohn",
	"itemsOrdered": [
		{
			"category": "Pizza",
			"extraIncrement": [
				"Chicken",
				"Bacon"
			],
			"increment": "Large",
			"item": "Ole Pizza"
		}
	],
	"name": "Andrew",
	"phone": "6198896620",
	"price": 11.5
}
```


Before `POSTing` this, we will check the server to see all the current orders and make sure the order got added after we do the post.


```
https://kitchen.pause.pizza/v5/[token]/orders
```

Then we do the `POST` request to the client backend.
The client app will then parse this into a struct, send over the new order to the kitchen management app (which appends this to the list of items)


```
curl -H "Content-Type: application/json" -X POST -d '{"dorm":"Mohn","itemsOrdered":[{"category":"Pizza","extraIncrement":["Chicken","Bacon"],"increment":"Large","item":"Ole Pizza"}],"name":"Andrew","phone":"6198896620","price":11.5}' https://kitchen.pause.pizza/v5/[token]/checkout
```

```
https://kitchen.pause.pizza/v5/[token]/orders
```

Then, we simulate a pizza complete post as follows:

```
curl -X POST https://kitchen.pause.pizza/v5/[token]/ordercomplete/[order#]
```

The server side checks to make sure that order is a valid order, removes it from the list of orders, and sends a similar `POST` request to the client ordering app.

---

## Future plans
1. Finish implementing the websockets servers (both kitchen management and client ordering side) to allow for notifications (which requires bidirectional streaming not present in the default HTTP/1.1 specs)
2.  Move the authentication portion to be on a database (I'm thinking of using Redis since I have had a very positive experience with it and most of the objects can be stored very closely to their structure in JSON form, which would minimize the amount of refractoring code in order to accomidate this)
