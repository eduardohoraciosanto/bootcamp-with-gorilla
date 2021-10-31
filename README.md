# Golang Bootcamp API with Gorilla

This is a solution of the third exercise of the Boot-Camp using Gorilla Mux.

The application is Dockerized, has a docker-compose file to boot the service quickly and a Makefile to start working ASAP.

	make init
	make build
	make up

And you are ready to Go!

`init` will populate the `.env` file needed for injecting environment variables.  
`build` will create the development image to code inside of it.  
`up` will run the API, exposing ports specified in the docker-compose file.  

This lets the developer focus on the code, running it inside the container resembling production.

---

## Exercise API

This exercise is about a REST API that lets you:

- Create a Shopping Cart
- Get all available Items from the external provider
- Get a particular Item from the external provider
- Add Items to a particular Shopping Cart
- Modify the amount of a particular item in a particular Shopping Cart
- Delete a particular item in a particular Shopping Cart
- Delete all items of a particular Shopping Cart
- Delete a particular Shopping Cart

### Database

We'll use Redis as a Database.
Since the idea is a simple Cart API, a NO-SQL Database seems perfect.

### Documentation
Documentation of the endpoints will be done using OpenAPI spec in Swagger format.  

---

### External API
**Note:** Available articles to be added to the shopping cart are provided by the following third party web service.

| Method   |      URL    | Desc |
|----------|-------------|---   |
| GET | https://bootcamp-products.getsandbox.com/products | To get all available products |
| GET | https://bootcamp-products.getsandbox.com/products/{id} | To get an specific product by id. It returns `404` if the _id_ is not found |

---

## Endpoints

More details are available in the swagger, here's just the list of available endpoints.
<img width="646" alt="image" src="https://user-images.githubusercontent.com/42719608/139605073-87f8b9df-5499-4633-bdc8-2c3bd83d5e01.png">

## Unit Testing

The unit testing was done using the dependency injection technique. This enabled coverage level to reach 100%

Same as the development, the unit testing is performed inside the docker container, to do so run the following:

	make devshell
	make t

`devshell` will run the develpoment container and start a terminal inside of it.  
`t` will run the unit testing and provide the coverage level for each package.  

Here's a screen capture of the test coverage:
<img width="899" alt="image" src="https://user-images.githubusercontent.com/42719608/139605136-8f5c66ed-a305-4a26-a2cb-75288acd0b05.png">


