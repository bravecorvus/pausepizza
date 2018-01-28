# Backend Implementations

## Client Ordering Server

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



## Management App
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
