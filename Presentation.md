# Presentation 1/29/18
### Andrew Lee

## Customer Backend POST
In addition to all the GET endpoints, there is now a POST endpoint which parses takes the following JSON object, and creates a new order on the Kitchen Managing side application.


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

### TODO
Implement websockets on Client Side
1. Client sends order to backend
2. backend processes order, returns a orderID
3. Client then connects to a web socket using the orderID as a slug
4. When the order is complete, that specific user will be notified.


Implement websockets on Kitchen side
1. Websocket connection established the first time user logs in
2. Every user on the Kitchen App will be notified anytime a new order is POSTED from the Client ordering app.

## Kitchen Management Stuff

### Image Processing



The server will take a valid JPEG or PNG and then convert it to a 500x500px JPEG (as well as create a monochromatic version of it)

The following will not work.
```
/Users/andrewlee/go/src/github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/development/sendimage
curl -F 'file=@pizzagif.gif'  localhost:7000/v5/8Dk6b5yndkjItF4RoeDhcz6ESFtHVTMA2jSNyL6Z/pizza/specialty/Chicken_Bacon_Ranch_Pizza.Large
```

Link
```
http://localhost:8000/files/chickenbaconranchlarge.jpg
```

This will work.
```
curl -F 'file=@pizza.JPEG'  localhost:7000/v5/8Dk6b5yndkjItF4RoeDhcz6ESFtHVTMA2jSNyL6Z/pizza/specialty/Chicken_Bacon_Ranch_Pizza.Large
```

### Documentation
[godoc](https://godoc.org/github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server)

### Superadmin
The superadmin account is an admin account on the kitchen management server that cannot be deleted (which will be a fail safe for if a user+pass system becomes compromised).

It is reset every day at midnight, and the combination is sent to people on a mailing list (Airmail).

### Tests
go test ./pizza

### All Post endpoints complete on server side

## TODO
- Statically compiled Docker scratch images
- Implement SSL to make the API more secure
