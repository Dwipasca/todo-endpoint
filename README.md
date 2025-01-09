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
  {
    "code": 200,
    "status": "Success",
    "message": "Success get list notes",
    "data": [
      {
        "id": 1,
        "title": "golang",
        "body": "bahasa golang dari google"
      },
      {
        "id": 2,
        "title": "js",
        "body": "javascript beda dari java"
      }
    ]
  }
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
