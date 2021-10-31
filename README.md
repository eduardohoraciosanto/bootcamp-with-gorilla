# Golang Bootcamp API with Gorilla

This is a solution of the third exercise of the Boot-Camp using Gorilla Mux.

---

## Exercise API

This exercise is about a REST API that lets you:

- Create a Shopping Cart
- Get all available Items
- Add Items to a particular Shopping Cart
- Modify the amount of a particular item in a particular Shopping Cart
- Delete a particular item in a particular Shopping Cart
- Delete all items of a particular Shopping Cart
- Delete a particular Shopping Cart

## Database

I'll use Redis as a Database.
Since the idea is a simple Cart API, a NO-SQL Database seems perfect.

## Documentation
Documentation of the endpoints will be done using OpenAPI spec in Swagger format.  

---

**Note:** Available articles to be added to the shopping cart are provided by the following third party web service.

| Method   |      URL    | Desc |
|----------|-------------|---   |
| GET | https://bootcamp-products.getsandbox.com/products | To get all available products |
| GET | https://bootcamp-products.getsandbox.com/products/{id} | To get an specific product by id. It returns `404` if the _id_ is not found |

---

## Endpoints

### Create Cart

*Method*: POST  
*URL*: /cart  
*Expected Request Body*: Empty  
*Expected Response Bodies*  

```
Success: Status 200 OK
Data:
{
    meta:{
        version: {string},
    },
    data:{
        id: {UUID},
        items: []
    }
}

Error: Status 500 Internal Server Error
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_internal
        description: Internal Server Error
    }
}
```

### Get Available Items

*Method*: GET  
*URL*: /items/available
*Expected Request Body*:Empty  
*Expected Response Bodies*  

```
Success: Status 200 OK
Data: Available Items
{
    meta:{
        version: {string},
    },
    data:{
        items: [
            {
                id: {string}
                name: {string},
                price: {float}
            }
        ]
    }
}

Error: Status 500 Internal Server Error
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_internal
        description: Internal Server Error
    }
}
```

### Add Item to Cart

*Method*: POST  
*URL*: /cart/{cart_id}/item  
- cart_id is the UUID of the particular cart  
*Expected Request Body*:
```
Data:
{
    id: {string},
    quantity: {int}
    
}

The id MUST be one of the available items provided via the -Get all available Items- endpoint
```  

*Expected Response Bodies*  

```
Success: Status 200 OK
Data: Full cart with updated Items
{
    meta:{
        version: {string},
    },
    data:{
        id: {UUID},
        items: [
            {
                name: {string}
                price: {float}
                quantity: {int}
            }
        ]
    }
}

Error: Status 422 Unprocessable Entity
When the Item is not one of the available
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_item_unavailable
        description: Item not available
    }
}

Error: Status 404 Not Found
When the cart ID is not in the database
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_cart_not_found
        description: Cart ID not found
    }
}
Error: Status 400 Bad Request
When the request body is malformed or incorrect
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_bad_request
        description: Request body malformed or incorrect
    }
}

Error: Status 500 Internal Server Error
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_internal
        description: Internal Server Error
    }
}
```

### Modify Item Quantity

*Method*: PUT  
*URL*: /cart/{cart_id}/item/{item_id}  
- cart_id is the UUID of the particular cart  
- item_id is the ID of the particular item to modify  
*Expected Request Body*:
```
Data:
{
    quantity: {int}
}

The item_id MUST be one of the available items provided via the -Get all available Items- endpoint
```  

*Expected Response Bodies*  

```
Success: Status 200 OK
Data: Full cart with updated Items
{
    meta:{
        version: {string},
    },
    data:{
        id: {UUID},
        items: [
            {
                name: {string}
                price: {float}
                quantity: {int}
            }
        ]
    }
}

Error: Status 422 Unprocessable Entity
When the Item is not one of the available
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_item_unavailable
        description: Item not available
    }
}

Error: Status 404 Not Found
When the cart ID is not in the database
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_cart_not_found
        description: Cart ID not found
    }
}

Error: Status 404 Not Found
When the Item ID is not in the cart
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_item_not_found
        description: Item ID not found in Cart
    }
}

Error: Status 400 Bad Request
When the request body is malformed or incorrect
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_bad_request
        description: Request body malformed or incorrect
    }
}

Error: Status 500 Internal Server Error
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_internal
        description: Internal Server Error
    }
}
```

### Delete Item from Cart

*Method*: DELETE  
*URL*: /cart/{cart_id}/item/{item_id}  
- cart_id is the UUID of the particular cart  
- item_id is the ID of the particular item to modify  
*Expected Request Body*: Empty  
*Expected Response Bodies*  

```
Success: Status 200 OK
Data: Full cart with updated Items
{
    meta:{
        version: {string},
    },
    data:{
        id: {UUID},
        items: [
            {
                name: {string}
                price: {float}
                quantity: {int}
            }
        ]
    }
}

Error: Status 404 Not Found
When the cart ID is not in the database
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_cart_not_found
        description: Cart ID not found
    }
}

Error: Status 404 Not Found
When the Item ID is not in the cart
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_item_not_found
        description: Item ID not found in Cart
    }
}

Error: Status 500 Internal Server Error
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_internal
        description: Internal Server Error
    }
}
```

### Delete All Items from a Cart

*Method*: DELETE  
*URL*: /cart/{cart_id}/item/all
- cart_id is the UUID of the particular cart  
*Expected Request Body*: Empty  
*Expected Response Bodies*  

```
Success: Status 200 OK
Data: Full cart with no Items
{
    meta:{
        version: {string},
    },
    data:{
        id: {UUID},
        items: []
    }
}

Error: Status 404 Not Found
When the cart ID is not in the database
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_cart_not_found
        description: Cart ID not found
    }
}

Error: Status 500 Internal Server Error
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_internal
        description: Internal Server Error
    }
}
```

### Delete Cart

*Method*: DELETE  
*URL*: /cart/{cart_id}
- cart_id is the UUID of the particular cart  
*Expected Request Body*: Empty  
*Expected Response Bodies*  

```
Success: Status 204 No Content
Data: Empty

Error: Status 404 Not Found
When the cart ID is not in the database
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_cart_not_found
        description: Cart ID not found
    }
}

Error: Status 500 Internal Server Error
Data:
{
    meta:{
        version: {string},
    },
    error:{
        code: err_internal
        description: Internal Server Error
    }
}
```



