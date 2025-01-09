package main

import (
	"net/http"
	"todo-endpoint/handlers"
)

func main() {
	handler := http.NewServeMux()
	baseUrl := "localhost"
	port := "8080"

	handler.HandleFunc("/api/v1/todo/add", handlers.AddTodo)
	handler.HandleFunc("/api/v1/todo/update/", handlers.UpdateTodo)
	handler.HandleFunc("/api/v1/todo/delete/", handlers.DeleteTodo)
	handler.HandleFunc("/api/v1/todo/detail/", handlers.DetailTodo)
	handler.HandleFunc("/api/v1/todos", handlers.GetAllTodo)

	http.ListenAndServe(baseUrl+":"+port, handler)
}