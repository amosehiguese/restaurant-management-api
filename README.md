# Restaurant Management API

The restaurant management API is a comprehensive solution designed to streamline and automate various aspects of restaurant operations. It provides a set of endpoints to manage menus, orders, reservations, tables, users, and more...

### Run the Server

To run the server, make sure you have postgresql and redis running.
Execute:
```bash
make start
```

### Libraries and features 
- Command Line flags with flag library
- Structured logging with zap
- Graceful Exits
- Chi as HTTP framework
- Custom Middlewares
- Pagination
- Documentation as Code with http-swagger

### Deployment and CI/CD Stack
- Docker
- Kubernetes (AWS EKS)
- Helm
- AWS IAM
- Cloud Formation
- Code Build
- Code Pipeline

### API ENDPOINTS

`GET '/api/v1/menu'`

- returns all menu
- Request Arguments: None
- Returns: json

```json
{
    "data": [
        {
            "id": "f312c0a8-7080-4b12-9abe-a37482b57efc",
            "name": "Breakfast Menu",
            "description": "Start your day with our delicious breakfast options!"
        }
    ],
    "success": true
}
```

`GET '/api/v1/menu/{id}'`

- return a single menu with the given id
- Request Arguments: `id` - integer
- Returns: json

```json
{
    "menu": {
        "id": "661b9383-ec2c-450c-9770-1aaf6b1972c3",
        "name": "Breakfast Menu",
        "description": "Start your day with our delicious breakfast options"
    },
    "success": true
}
```

`POST '/api/v1/admin/menu'`

- Sends a post request to create a menu
- Requires Admin access token
- Request Body:

```json
{
    "name": "Breakfast Menu",
    "description": "Start your day with our delicious breakfast options!"
}
```
- Returns: an object containing the id of the just created menu

```json
{
    "menu_id": "f312c0a8-7080-4b12-9abe-a37482b57efc",
    "success": true
}
```

`PATCH '/api/v1/admin/menu/{id}'`

- Sends a patch request to update a menu with the specified id
- Request Arguments: `id` - integer
- Requires Admin access token
- Request Body:

```json
{
    "name": "Breakfast Menu",
    "description": "Start your day with our delicious breakfast options!"
}
```
- Returns: an object containing a success message

```json
{
    "msg": "Successfully update menu with id 661b9383-ec2c-450c-9770-1aaf6b1972c3",
    "success": true
}
```

`DELETE '/api/v1/admin/menu/{id}'`

- Deletes a specified menu corresponding to the id of the menu
- Request Arguments: `id` - integer
- Requires Admin access token
- Returns: an object containing a success message

```json
{   
    "msg":"Menu with id 661b9383-ec2c-450c-9770-1aaf6b1972c3 is successfully deleted",
    "success":true
}
```


<!-- dishes -->
`GET '/api/v1/menu{id}/dishes'`

- returns all dishes belonging to a specified menu
- Request Arguments: `id` - integer
- Returns: json

```json
{
    "data": [
        {
            "id": "bc09787b-9a00-48ea-91b4-6b9f76166130",
            "name": "Pancakes",
            "description": "Delicious fluffy pancakes",
            "price": "7.99",
            "menu_id": "2a780496-e12f-4679-a5d5-e2bcd22d1660"
        }
    ],
    "success": true
}
```

`GET '/api/v1/menu/{id}/dishes/{dishID}'`

- return a single dish with the given id
- Request Arguments: `id` - integer
- Request Arguments: `dishID` - integer
- Returns: json

```json
{
    "dish": {
        "id": "bc09787b-9a00-48ea-91b4-6b9f76166130",
        "name": "Pancakes",
        "description": "Delicious fluffy pancakes",
        "price": "7.99",
        "menu_id": "2a780496-e12f-4679-a5d5-e2bcd22d1660"
    },
    "success": true
}
```

`POST '/api/v1/admin/menu/{id}/dishes'`

- Sends a post request to create a dish
- Request Arguments: `id` - integer
- Requires Admin access token
- Request Body:

```json
{
"name": "Pancakes",
"description": "Delicious fluffy pancakes",
"price": "7.99"
}
```
- Returns: an object containing the id of the just created dish

```json
{
    "dish_id": "bc09787b-9a00-48ea-91b4-6b9f76166130",
    "success": true
}
```

`PATCH '/api/v1/admin/menu/{id}/dishes/{dishID}'`

- Sends a patch request to update a dish with the specified dishID
- Request Arguments: `id` - integer
- Request Arguments: `dishID` - integer
- Requires Admin access token
- Request Body:

```json
{
    "name": "Pancakes",
    "description": "Delicious fluffy pancakes",
    "price": "9.99"
}
```
- Returns: an object containing a success message

```json
{
    "msg": "Successfully update dish with id 304fe67f-adda-47ba-8c75-418e3512856d",
    "success": true
}
```

`DELETE '/api/v1/admin/menu/{id}/dishes/{dishID}'`

- Deletes a specified dish corresponding to the id of the dish
- Request Arguments: `id` - integer
- Request Arguments: `dishID` - integer
- Requires Admin access token
- Returns: an object containing a success message

```json
{   
    "msg":"Dish with id 661b9383-ec2c-450c-9770-1aaf6b1972c3 is successfully deleted",
    "success":true
}
```

<!-- Tables -->
`GET '/api/v1/tables'`

- returns all tables
- Request Arguments: None
- Returns: json

```json
{
    "data": [
        {
            "id": "08a81a18-ce5a-4d69-b701-6b479af8b8cd",
            "number": 5,
            "capacity": 4,
            "status": "available"
        },
        {
            "id": "09e644ea-9d1f-441e-a732-9e8d8e10b0ba",
            "number": 3,
            "capacity": 6,
            "status": "available"
        }
    ],
    "success": true
}
```

`GET '/api/v1/tables/{id}'`

- return a single table with the given id
- Request Arguments: `id` - integer
- Returns: json

```json
{
    "table": {
        "id": "08a81a18-ce5a-4d69-b701-6b479af8b8cd",
        "number": 5,
        "capacity": 4,
        "status": "available"
    },
    "success": true
}
```

`POST '/api/v1/admin/tables'`

- Sends a post request to create a table
- Requires Admin access token
- Request Body:

```json
{
  "number": 5,
  "capacity": 4,
  "status": "available"
}
```
- Returns: an object containing the id of the just created table

```json
{
    "table_id": "08a81a18-ce5a-4d69-b701-6b479af8b8cd",
    "success": true
}
```

`PATCH '/api/v1/admin/tables/{id}'`

- Sends a patch request to update a table with the specified id
- Request Arguments: `id` - integer
- Requires Admin access token
- Request Body:

```json
{
  "number": 5,
  "capacity": 17,
  "status": "available"
}
```
- Returns: an object containing a success message

```json
{
    "msg": "Successfully update table with id 08a81a18-ce5a-4d69-b701-6b479af8b8cd",
    "success": true
}
```

`DELETE '/api/v1/admin/tables/{id}'`

- Deletes a specified table corresponding to the id of the table
- Request Arguments: `id` - integer
- Requires Admin access token
- Returns: an object containing a success message

```json
{   
    "msg":"Table with id 08a81a18-ce5a-4d69-b701-6b479af8b8cd is successfully deleted",
    "success":true
}
```

<!-- Reservations -->
`GET '/api/v1/p/reservations'`

- returns all reservations
- Request Arguments: None
- Returns: json

```json
{
    "data": [
        {
            "id": "de719e76-5e90-4806-b3d3-10a6774745f5",
            "table_id": "08a81a18-ce5a-4d69-b701-6b479af8b8cd",
            "reservation_date": "2023-12-25T00:00:00Z",
            "reservation_time": "0000-01-01T19:00:00Z",
            "status": "available",
            "created_at": "2024-01-08T11:32:46.08547Z",
            "updated_at": "0001-01-01T00:00:00Z"
        },
        {
            "id": "e17f5d23-f773-42a6-804f-25baa20a0bef",
            "table_id": "08a81a18-ce5a-4d69-b701-6b479af8b8cd",
            "reservation_date": "2023-12-25T00:00:00Z",
            "reservation_time": "0000-01-01T19:00:00Z",
            "status": "available",
            "created_at": "2024-01-08T11:35:17.536142Z",
            "updated_at": "0001-01-01T00:00:00Z"
        }
    ],
    "success": true
}
```

`GET '/api/v1/p/reservations/{id}'`

- return a single reservation with the given id
- Requires Authenticated user access token
- Request Arguments: `id` - integer
- Returns: json

```json
{
    "reservation": {
        "id": "de719e76-5e90-4806-b3d3-10a6774745f5",
        "table_id": "08a81a18-ce5a-4d69-b701-6b479af8b8cd",
        "reservation_date": "2023-12-25T00:00:00Z",
        "reservation_time": "0000-01-01T19:00:00Z",
        "status": "available",
        "created_at": "2024-01-08T11:32:46.08547Z",
        "updated_at": "2024-01-08T11:32:46.08547Z"
    },
    "success": true
}
```

`POST '/api/v1/p/reservations'`

- Sends a post request to create a reservation
- Requires Authenticated user access token
- Request Body:

```json
{
    "table_id": "08a81a18-ce5a-4d69-b701-6b479af8b8cd",
    "reservation_date": "2023-12-25",
    "reservation_time": "19:00:00",
    "status": "available"
}
```
- Returns: an object containing the id of the just created reservation

```json
{
    "reservation_id": "e17f5d23-f773-42a6-804f-25baa20a0bef",
    "success": true
}
```

`PATCH '/api/v1/p/reservation/{id}'`

- Sends a patch request to update a reservation with the specified id
- Request Arguments: `id` - integer
- Requires Authenticated user access token
- Request Body:

```json
{
    "table_id": "08a81a18-ce5a-4d69-b701-6b479af8b8cd",
    "reservation_date": "2023-12-25",
    "reservation_time": "19:45:00",
    "status": "canceled"
}
```
- Returns: an object containing a success message

```json
{
    "msg": "Successfully update reservation with id e17f5d23-f773-42a6-804f-25baa20a0bef",
    "success": true
}
```

`DELETE '/api/v1/p/reservations/{id}'`

- Deletes a specified reservation corresponding to the id of the reservation
- Request Arguments: `id` - integer
- Requires Authenticated user access token
- Returns: an object containing a success message

```json
{   
    "msg":"Reservation with id e17f5d23-f773-42a6-804f-25baa20a0bef is successfully deleted",
    "success":true
}
```

<!-- Invoices -->
`GET '/api/v1/admin/invoices'`

- returns all invoices
- Requires Admin access token
- Request Arguments: None
- Returns: json

```json
{
    "data": [
        {
            "id": "9c8d1aea-0174-4a90-ab5e-22ab9e76cb7f",
            "order_id": "ee3f1a10-533f-48ce-9e2e-f0ad3daa4259",
            "invoice_date": "2024-01-08T00:00:00Z",
            "total_amount": "69.97",
            "tax": "5.60",
            "discount": "0.08",
            "grand_total": "75.57"
        },
        {
            "id": "6b50dabf-d066-4bc0-8c93-89401ced3521",
            "order_id": "538b12c5-4e2e-450e-ae3e-d1c311e346fc",
            "invoice_date": "2024-01-08T00:00:00Z",
            "total_amount": "70.97",
            "tax": "7.60",
            "discount": "0.00",
            "grand_total": "77.97"
        }
    ],
    "success": true
}
```

`GET '/api/v1/p/invoice/{id}'`

- return a single invoice with the given id
- Requires Authenticated user access token
- Request Arguments: `id` - integer
- Returns: json

```json
{
    "invoice": {
        "id": "9c8d1aea-0174-4a90-ab5e-22ab9e76cb7f",
        "order_id": "ee3f1a10-533f-48ce-9e2e-f0ad3daa4259",
        "invoice_date": "2024-01-08T00:00:00Z",
        "total_amount": "69.97",
        "tax": "5.60",
        "discount": "0.08",
        "grand_total": "75.57"
    },
    "success": true
}
```

`POST '/api/v1/p/invoices'`

- Sends a post request to create a reservation
- Requires Authenticated user access token
- Request Body:

```json
{
    "order_id": "538b12c5-4e2e-450e-ae3e-d1c311e346fc",
    "total_amount": "69.97",
    "discount": "0.08",
    "tax": "5.60",
    "grand_total": "75.57"
}
```
- Returns: an object containing the id of the just created invoice

```json
{
    "invoice_id": "6b50dabf-d066-4bc0-8c93-89401ced3521",
    "success": true
}
```

`PATCH '/api/v1/admin/invoices/{id}'`

- Sends a patch request to update a reservation with the specified id
- Request Arguments: `id` - integer
- Requires Admin access token
- Request Body:

```json
{
    "order_id": "6b50dabf-d066-4bc0-8c93-89401ced3521",
    "total_amount": "99.97",
    "discount": "0.10",
    "tax": "5.60",
    "grand_total": "75.57"
}
```
- Returns: an object containing a success message

```json
{
    "msg": "Successfully update invoice with id 6b50dabf-d066-4bc0-8c93-89401ced3521",
    "success": true
}
```

<!-- Roles -->
`GET '/api/v1/admin/roles'`

- returns all roles
- Requires Admin access token
- Request Arguments: None
- Returns: json

```json
{
    "data": [
        {
            "id": 1,
            "name": "admin",
            "description": "Administrator role"
        },
        {
            "id": 2,
            "name": "user",
            "description": "Authenticated user role"
        },
        {
            "id": 3,
            "name": "anonymous",
            "description": "Unauthenticated user role"
        }
    ],
    "success": true
}
```

`POST '/api/v1/admin/roles'`

- Sends a post request to create a roles
- Requires Admin access token
- Request Body:

```json
{
    "name": "staff",
    "description":"staff role"
}
```
- Returns: an object containing the details of the created role

```json
{
    "role": {
        "id": 4,
        "name": "staff",
        "description": "staff role"
    },
    "success": true
}
```

`PATCH '/api/v1/admin/roles/{id}'`

- Sends a patch request to update a role with the specified id
- Request Arguments: `id` - integer
- Requires Admin access token
- Request Body:

```json
{
    "name": "staff",
    "description":"staff rolee"
}
```
- Returns: an object containing a success message

```json
{
    "msg": "Successfully update role with id 13",
    "success": true
}
```

`DELETE '/api/v1/admin/roles/{id}'`

- Deletes a specified roles corresponding to the id of the roles
- Request Arguments: `id` - integer
- Requires Admin access token
- Returns: an object containing a success message

```json
{   
    "msg":"Role with id 13 is successfully deleted",
    "success":true
}
```


<!-- User -->
`GET '/api/v1/admin/users'`

- returns all users
- Requires Admin access token
- Request Arguments: None
- Returns: json

```json
{
    "data": [
        {
            "id": "7ab000a7-2942-42ee-af1b-a36e052d19d3",
            "first_name": "john",
            "last_name": "okon",
            "username": "john",
            "email": "john@gmail.com",
            "password_hash": "$2a$04$1eLqir9kaAcxGj1Ma/pUye4i5wmNulcYpGLq4Zp9Uc8DCBB0KwpLm",
            "user_role": 2,
            "created_at": "2024-01-08T13:38:55.6088Z",
            "updated_at": {
                "Time": "0001-01-01T00:00:00Z",
                "Valid": false
            }
        },
        {
            "id": "a532c2cf-af57-4847-a596-07af49e807a0",
            "first_name": "sam",
            "last_name": "crooke",
            "username": "sammy",
            "email": "crookesam@gmail.com",
            "password_hash": "$2a$04$DIFzUsicuLBCS3ceTVTSt.5RrZNJCsd6ja.Lot3.YOoCeNmQX.KUG",
            "user_role": 2,
            "created_at": "2024-01-08T13:39:25.617949Z",
            "updated_at": {
                "Time": "0001-01-01T00:00:00Z",
                "Valid": false
            }
        }
    ],
    "success": true
}
```

`GET '/api/v1/p/users/{id}'`

- return a single user with the given id
- Requires Admin access token
- Request Arguments: `id` - integer
- Returns: json

```json
{
    "success": true,
    "user_id": {
        "id": "a532c2cf-af57-4847-a596-07af49e807a0",
        "first_name": "sam",
        "last_name": "crooke",
        "username": "sammy",
        "email": "crookesam@gmail.com",
        "password_hash": "$2a$04$DIFzUsicuLBCS3ceTVTSt.5RrZNJCsd6ja.Lot3.YOoCeNmQX.KUG",
        "user_role": 2,
        "created_at": "2024-01-08T13:39:25.617949Z",
        "updated_at": {
            "Time": "0001-01-01T00:00:00Z",
            "Valid": false
        }
    }
}
```

`PATCH '/api/v1/p/users/{id}'`

- Sends a patch request to update a user with the specified id
- Request Arguments: `id` - integer
- Requires Admin access token
- Request Body:

```json
{
    "first_name":"isaac",
    "last_name": "okon",
    "username": "isaacokon"
}
```
- Returns: an object containing a success message

```json
{
    "msg": "Successfully update user with id 7ab000a7-2942-42ee-af1b-a36e052d19d3",
    "success": true
}
```

<!-- Orders -->
`GET '/api/v1/p/orders'`

- returns all orders
- Request Arguments: None
- Returns: json

```json
{
    "data": [
        {
            "id": "ee3f1a10-533f-48ce-9e2e-f0ad3daa4259",
            "status": "pending",
            "created_at": "2024-01-08T13:11:20.239354Z",
            "updated_at": {
                "Time": "0001-01-01T00:00:00Z",
                "Valid": false
            }
        },
        {
            "id": "538b12c5-4e2e-450e-ae3e-d1c311e346fc",
            "status": "pending",
            "created_at": "2024-01-08T13:12:13.804157Z",
            "updated_at": {
                "Time": "0001-01-01T00:00:00Z",
                "Valid": false
            }
        },
        {
            "id": "72cc040e-dd06-4bde-b480-87363aa538c3",
            "status": "completed",
            "created_at": "2024-01-08T13:52:51.23471Z",
            "updated_at": {
                "Time": "0001-01-01T00:00:00Z",
                "Valid": false
            }
        }
    ],
    "success": true
}
```

`GET '/api/v1/p/orders/{id}'`

- return a single order with the given id
- Requires Authenticated user access token
- Request Arguments: `id` - integer
- Returns: json

```json
{
    "order": {
        "created_at": "2024-01-08T13:11:20.239354Z",
        "order_id": "ee3f1a10-533f-48ce-9e2e-f0ad3daa4259",
        "status": "pending",
        "updated_at": "0001-01-01T00:00:00Z"
    },
    "success": true
}
```

`POST '/api/v1/p/orders'`

- Sends a post request to create an order
- Requires Authenticated user access token
- Request Body:

```json
{
    "status": "pending"
}
```
- Returns: an object containing the id of the just created order

```json
{
    "order_id": "ee3f1a10-533f-48ce-9e2e-f0ad3daa4259",
    "success": true
}
```

`PATCH '/api/v1/p/orders/{id}'`

- Sends a patch request to update a order with the specified id
- Request Arguments: `id` - integer
- Requires Authenticated user access token
- Request Body:

```json
{
    "status": "pending"
}
```
- Returns: an object containing a success message

```json
{
    "msg": "Successfully update order with id ee3f1a10-533f-48ce-9e2e-f0ad3daa4259",
    "success": true
}
```

`DELETE '/api/v1/p/order/{id}'`

- Deletes a specified order corresponding to the id of the order
- Request Arguments: `id` - integer
- Requires Authenticated user access token
- Returns: an object containing a success message

```json
{   
    "msg":"Order with id e17f5d23-f773-42a6-804f-25baa20a0bef is successfully deleted",
    "success":true
}
```


<!-- OrderItems -->
`GET '/api/v1/p/orders/{id}/items'`

- returns all order items belonging to the order's id
- Request Arguments: None
- Returns: json

```json
{
    "data": [
        {
            "id": "c162726f-40bc-4047-8e13-86ab332cc867",
            "order_id": "ee3f1a10-533f-48ce-9e2e-f0ad3daa4259",
            "dish_id": "cc64d751-7d75-46a4-9c42-b0c35ce4a450",
            "quantity": 3
        }
    ],
    "success": true
}
```

`PUT '/api/v1/p/orders/{id}/items'`

- Sends a post request to create an order item for a specific order
- Requires Authenticated user access token
- Request Body:

```json
{
    "dish_id": "cc64d751-7d75-46a4-9c42-b0c35ce4a450",
    "quantity": 3
}
```
- Returns: an object containing the id of the just created order item

```json
{
    "order_item_id": "c162726f-40bc-4047-8e13-86ab332cc867",
    "success": true
}
```

`PATCH '/api/v1/p/orders/{id}/items/{itemID}'`

- Sends a patch request to update an order item under the specified order id
- Request Arguments: `id` - integer
- Request Arguments: `itemID` - integer
- Requires Authenticated user access token
- Request Body:

```json
{
    "dish_id": "cc64d751-7d75-46a4-9c42-b0c35ce4a450",
    "quantity": 7
}
```
- Returns: an object containing a success message

```json
{
    "msg": "Successfully update order item with id c162726f-40bc-4047-8e13-86ab332cc867",
    "success": true
}
```

`DELETE '/api/v1/p/order/{id}/items/{itemID}'`

- Deletes a specified order item corresponding to the id of the order
- Request Arguments: `id` - integer
- Request Arguments: `itemID` - integer
- Requires Authenticated user access token
- Returns: an object containing a success message

```json
{   
    "msg":"Order item with id e17f5d23-f773-42a6-804f-25baa20a0bef is successfully deleted",
    "success":true
}
```

<!-- Auth -->
`POST '/api/v1/auth/signup'`

- Sends a post request to register to user
- Request Body:

```json
{
    "first_name":"sam",
    "last_name": "crooke",
    "username": "sammy",
    "email": "crookesam@gmail.com",
    "password":"sam123"
}
```
- Returns: an object containing details of the just created user

```json
{
    "success": true,
    "user": {
        "id": "a532c2cf-af57-4847-a596-07af49e807a0",
        "first_name": "sam",
        "last_name": "crooke",
        "username": "sammy",
        "email": "crookesam@gmail.com",
        "password_hash": "",
        "user_role": 2,
        "created_at": "2024-01-08T13:39:25.617949Z",
        "updated_at": {
            "Time": "0001-01-01T00:00:00Z",
            "Valid": false
        }
    }
}
```

`POST '/api/v1/auth/login'`

- Sends a post request to log-in
- Request Body:

```json
{
    "email": "john@gmail.com",
    "password": "john123"
}
```
- Returns: an object containing the user's access and refresh tokens

```json
{
    "success": true,
    "tokens": {
        "access": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE3MDQ3MjAxMjEsImlhdCI6MTcwNDcxOTIyMSwiaWQiOiI3YWIwMDBhNy0yOTQyLTQyZWUtYWYxYi1hMzZlMDUyZDE5ZDMiLCJyb2xlIjoyfQ.VkCEsQ3eakhdFQ4gHvbXn6mHdl4vC-SbdooNvOCK5FQ",
        "refresh": "10ea92439b256422b8078f8dc57653ad1029c93fce6655db454d754529f83911.1707311221"
    }
}
```

`POST '/api/v1/p/token/renew'`

- Sends a post request to renew tokens
- Request Body:

```json
{
    "refresh_token": "10ea92439b256422b8078f8dc57653ad1029c93fce6655db454d754529f83911.1707311221"
}
```
- Returns: an object containing tokens

```json
{
    "success": true,
    "tokens": {
        "access": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE3MDQ3MjA1MjksImlhdCI6MTcwNDcxOTYyOSwiaWQiOiI3YWIwMDBhNy0yOTQyLTQyZWUtYWYxYi1hMzZlMDUyZDE5ZDMiLCJyb2xlIjoyfQ.mYu_OS_ML3Wkp3GyilTY36Ws13rySquSV22lNLUZns4",
        "refresh": "d7ee2454ffb663097d81ac62fa940b0eccd24bbe2c64b9747157d811b2d5c793.1707311629"
    }
}
```























