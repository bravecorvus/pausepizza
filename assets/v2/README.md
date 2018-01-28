# Flow Chart

## 1. App Start
`fetch('ip:port/v2/landing')`

```json
{
  "KitchenOpen ": true,
  "OvenOpen" : false,
  "LateMenu" : true
}
```

## 2. App UI Rendering Choices

### If `KitchenOpen == false`
Pause Kitchen Closed (No further UI options available)

### If `KitchenOpen == true`

#### `OvenOpen == true` and `LateMenu == false`
All UI options available


#### `OvenOpen == true` and `LateMenu == true`
Togo: All options available
Pick-up: Pizzas, milkshake, icecream only 

#### `OvenOpen == false` and `LateMenu == true`
Togo: Not available
Pick-up: Milkshake, icecream only 

#### `OvenOpen == false` and `LateMenu == false`
Togo: Not available
Pick-up: Milkshake, icecream only, nachos


API Endpoints


[`http://ip:port/v2/landing`](server_to_client/landing.json)
[`http://ip:port/v2/togo`](server_to_client/togo.json)
[`http://ip:port/v2/pickup`](server_to_client/pickup.json)

[`http://ip:port/v2/food/pizza/specialty/list`](server_to_client/specialtypizza.json)
[`http://ip:port/v2/food/pizza/specialty/type/chickenbaconranch`](server_to_client/chickenbaconranch.json)

[`http://ip:port/v2/food/pizza/normal/list`](server_to_client/normallist.json)
[`http://ip:port/v2/food/pizza/normal/type/pepperoni`](server_to_client/pepperoni.json)

[`http://ip:port/v2/food/pizza/ingredients/list`](server_to_client/ingredientslist.json)
[`http://ip:port/v2/food/pizza/ingredients/type/mushroom`](server_to_client/mushroom.json)

[`http://ip:port/v2/desserts/milkshake`](server_to_client/milkshake.json)

[`http://ip:port/v2/food/chickenfingers`](server_to_client/chickenfingers.json)
[`http://ip:port/v2/food/cheesybread`](server_to_client/cheesybread.json)

[`http://ip:port/v2/dorms`](server_to_client/dorms.json)
