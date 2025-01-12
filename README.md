# Simple todo-endpoint | Golang

This is simple todo list endpoint that I created with go language

# Endpoint

## Create Todo

- URL : `http://localhost:8080/api/v1/todos`
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
    "message": "Successfully created new todo"
  }
  ```

## Update Todo

- URL : `http://localhost:8080/api/v1/todos?id={id}`
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
    "message": "Successfully update todo"
  }
  ```

## Delete Todo

- URL : `http://localhost:8080/api/v1/todos?id={id}`
- Method: `DELETE`
- response
  ```json
  {
    "status": "success",
    "code": 200,
    "message": "Successfully deleted todo"
  }
  ```

## List Todo

- URL : `http://localhost:8080/api/v1/todos`
- Method: `GET`
- response
  ```json
  {
    "status": "success",
    "code": 200,
    "data": [
      {
        "id": 2,
        "task": "Buy new macbook"
      },
      {
        "id": 3,
        "task": "Buy new macbook"
      }
    ],
    "message": "List todo retrieved successfully"
  }
  `
  ```

## Detail Todo

- URL : `http://localhost:8080/api/v1/todos/detail/{id}`
- Method: `GET`
- response
  ```json
  {
    "status": "success",
    "code": 200,
    "data": {
      "id": 3,
      "task": "Buy new macbook"
    },
    "message": "Todo retrieved successfully"
  }
  ```
