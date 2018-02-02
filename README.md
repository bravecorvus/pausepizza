# Backend Implementations
---
## Golang
### Andrew Lee<br>
**Please view Wiki in order to see implement system**
The Kitchen Web Server (which does most of the heavy lifting) has been heavily documented. If you have any questions about code specifically, head over to the project [godocs page](godoc.org/github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server).

### Client Ordering Server
#### Andrew Lee

The client ordering API is internally available at the url `162.210.90.60:8000/v5/...` (can be access anywhere user is on eduroam, or CS 10 Gig)

The API server returns STATEful resource (JSON files) which will be parsed by the React Native front-end, and converted into objects to populate the App.

> Examples:

When the user first opens, the App, the front-end will request the endpoint `http://162.210.90.60:8000/v5/landing`, which returns the following JSON object

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

the pauseinfo object towards the top defines the parameter of what gets rendered at the very beginning (i.e. If `kitchenOpen==false`, then nothing form the App will run, or if `deliveryAvailable==false`, then delivery items will be turned off)

As you can see, each subsequent endpoint is defined as part of the parent API resource

For example, in order to get the resources of the Pizza endpoint, we will go to `http://162.210.90.60:8000/v5/pizza`


> Some more links
`http://162.210.90.60:8000/v5/pizza/specialty`
`http://162.210.90.60:8000/v5/desserts`
`http://162.210.90.60:8000/v5/appetizers/cheesybread`



### Management App
#### Andrew Lee
The Management App manipulates the data of the Client Ordering application to reflect changes in the Pause such as the addition of new menu items, the deletion of certain items, running out of food etc.

Unlike the client ordering app, the management app needs to be secure. However, sending the username and password every time to request a resource in a header like this `curl -H "Content-Type: application/json" -X POST -d '{"username":"leeas","password":"abcd"}' 162.210.90.60:7000/v5/pizza` is not ideal as it is transfering sensitive data.

Therefore, I designed a token based system where the user *initially* authenticates using a username and password, gets the server to randomly generate a 40 character alphanumeric string as a secure token (which is given a expiration time), and can only access points of the app using this secure token.

> Example
The kitchen management app API is located at `162.210.90.60:7000/v5/...`

However, trying to access the normal endpoints as we did in the client app like this `162.210.90.60:7000/v5/landing` will return the following message:

```json
{"status":false,"message":"User did not access the link with a valid token. Please log in"}
```

Hence, the user must first log in using a username and password saved on the server side.

```
curl -H "Content-Type: application/json" -X POST -d '{"username":"notarealaccount","password":"notarealpassword"}' 162.210.90.60:7000/v5/login
```
The example above, is showing that only authorized username+password combinations will work. It will yield the response

```json
{"status":false,"message":"Username and Password combination does not exist. If you forgot your password, please talk to another manager for the app "}
```

However, a real user with a real password will be able to receive a secure token
```
curl -H "Content-Type: application/json" -X POST -d '{"username":"leeas","password":"abcd"}' 162.210.90.60:7000/v5/login
```

Which will return a JSON object similar to the following:
```json
{"associatedUser":"leeas","value":"N95RxRTZHWUsaD6HEdz0ThbXfQ6pYSQ3n267l1VQ","timestamp":"2018-01-24T13:11:23.309813635-06:00"}
```

Then, we use the `value` as part of any request (both `GET` and `POST`)

```
http://162.210.90.60:7000/v5/N95RxRTZHWUsaD6HEdz0ThbXfQ6pYSQ3n267l1VQ/landing
```

As you can see, the token expired in 30 seconds because that is what I have it set to for this demo, but the expiration of this token will be somewhere between 1 hour and 24 hours during the production phase.


Server also runs a cron function to change the values of the `landing` endpoint of the client app
```
http://162.210.90.60:7000/v5//landing/set
```


### Steps to Ordering
1. POST request from Client Ordering App including Items list, cost, name of person, phone number, delivery location etc. to the `v5/checkout` endpoint
2. Client Ordering backend receives order. The backend will create a instance of a struct with the above information, and add a new order ID based on a randomly generated 40 character string.
3. Client Ordering backend will find the Super Admin username from the `assets/v5superadmin/list.json`, and then try to find a token in `assets/v5/tokens/list.json` that matches the user name. If the token exists, the Client Ordering Backend will attempt to POST the order (including the new order ID generated during step 2) at the `v5/neworder` endpoint on the Kitchen Management using the generated token. If the token does not exist, the Client Ordering backend will use the Super Admin username and password to generate a new token and go through the same exact steps to post to the kitchen management app.
4. Kitchen Management App registers new order, and adds it to the queue. It will send out a server to client notification regarding the addition of the new order (in order to dynamically modify the UI if the Orders page is being accessed at the moment) through a websocket (this time, all users of the kitchen management app will get a notification).
5. When the order gets processed on the kitchen management backend, the kitchen app will respond with a JSON struct (with a status of `success`). If something fails, it will respond with a status of `fail`.
6. If the above process registers a success, the client server will return an object also with a status = `success` portion as well as the generated orderID
7. The client ordering app front-end will then attempt to subsribe to a websocket using the endpoint `v5/websocket/[generated order id]`.
8. When the Kitchen sends an order complete POST to the kitchen server backend, the kitchen backend will POST to the `v5/ordercomplete` of the client ordering backend. Then the client ordering backend will send a notification to the specific user who sent the order whenever it is done.

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


Before POSTing this, we will check the server to see all the current orders and make sure the order got added after we do the post.


```
localhost:7000/v5/8Dk6b5yndkjItF4RoeDhcz6ESFtHVTMA2jSNyL6Z/orders
```

Then we do the POST request to the client backend.
The client app will then parse this into a struct, send over the new order to the kitchen management app (which appends this to the list of items)


```
curl -H "Content-Type: application/json" -X POST -d '{"dorm":"Mohn","itemsOrdered":[{"category":"Pizza","extraIncrement":["Chicken","Bacon"],"increment":"Large","item":"Ole Pizza"}],"name":"Andrew","phone":"6198896620","price":11.5}' localhost:8000/v5/checkout
```

```
localhost:7000/v5/8Dk6b5yndkjItF4RoeDhcz6ESFtHVTMA2jSNyL6Z/orders
```

Then, we simulate a pizza complete post as follows:

```
curl -X POST localhost:7000/v5/8Dk6b5yndkjItF4RoeDhcz6ESFtHVTMA2jSNyL6Z/ordercomplete/order_#
```

The server side checks to make sure that order is a valid order, removes it from the list of orders, and sends a similar POST request to the client ordering app.

---
## Java Customer Ordering Implemented<br>
### Roo Kosherbay

Java implementation of the static backendend API server for the Customer Ordering App.
