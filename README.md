# Simple todo-endpoint | Golang

This is simple todo list endpoint that I created with go language

# Endpoint

## Create Todo

- URL : `localhost:8000/api/v1/todo`
- Method: `POST`
- request body
  ```json
  {
    "task": "Buy new macbook"
  }
  ```
- response
  ```json
  {
    "status": "success",
    "code": 201,
    "task": "Buy new macbook",
    "message": "Successfully created new todo"
  }
  ```

## Update Todo

- URL : `localhost:8000/api/v1/todo/update/`
- Method: `PUT`
- request body
  ```json
  {
    "task": "Buy new monitor"
  }
  ```
- response
  ```json
  {
    "status": "success",
    "code": 200,
    "task": "buy new chair",
    "message": "Successfully updated todo"
  }
  ```

## Delete Todo

- URL : `localhost:8000/api/v1/todo/delete/`
- Method: `DELETE`
- response
  ```json
  {
    "status": "success",
    "code": 200,
    "task": "",
    "message": "Successfully deleted todo"
  }
  ```

## List Todo

- URL : `localhost:8000/api/v1/todos`
- Method: `GET`
- response
  ```json
  [
    {
      "id": 1,
      "task": "Buy new macbook"
    },
    {
      "id": 2,
      "task": "Buy new chair"
    }
  ]
  ```

## Detail Todo

- URL : `localhost:8000/api/v1/todo/detail/`
- Method: `GET`
- response
  ```json
  {
    "status": "success",
    "code": 200,
    "task": "Buy new macbook",
    "message": "Todo retrieved successfully"
  }
  ```
