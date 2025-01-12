package main

import (
	"net/http"
	"todo-endpoint/handlers"
)

func main() {
	handler := http.NewServeMux()
	baseUrl := "localhost"
	port := "8080"

	handler.HandleFunc("/api/v1/todos/detail/", handlers.DetailTodo)

	handler.HandleFunc("/api/v1/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodPost :
				handlers.AddTodo(w,r)
			case http.MethodPut : 
				handlers.UpdateTodo(w,r)
			case http.MethodDelete : 
				handlers.DeleteTodo(w,r)
			case http.MethodGet: 
				handlers.GetAllTodo(w,r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})	

	http.ListenAndServe(baseUrl+":"+port, handler)

}
