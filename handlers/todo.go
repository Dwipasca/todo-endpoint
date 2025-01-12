package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"todo-endpoint/models"
)

var (
	todos   = []models.Todo{}
	todoMux sync.Mutex
	nextID  = 1
)

func write(w http.ResponseWriter, status string, code int, data interface{}, message string) {
	var req = models.Response{
		Status:  status,
		Code:    code,
		Task:    data,
		Message: message,
	}

	res, _ := json.Marshal(req) // convert byte to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

func remove(slice []models.Todo, s int) []models.Todo {
	// slice[:s] -> get all data from index 0 to index s - 1
	// slice[s+1:] -> get all data after index s + 1
	// append(sllice/arr, index, ...) combine all the data
	return append(slice[:s], slice[s+1:]...)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		write(w, "error", http.StatusMethodNotAllowed, nil, "Method is not allowed")
		return
	}

	var newTodo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&newTodo);
	err != nil {
		write(w, "error", http.StatusBadRequest, nil, "Invalid input")
		return
	}

	// sync.Mutex (lock and unlock) used to avoid race condition
	todoMux.Lock()
	defer todoMux.Unlock()

	newTodo.ID = nextID
	nextID++;

	todos = append(todos, newTodo)

	write(w, "success", http.StatusCreated, nil, "Successfully created new todo")
	
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		write(w, "error", http.StatusMethodNotAllowed, nil, "Method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr) // convert from string to int
	if err != nil {
		write(w, "error", http.StatusBadRequest, nil, "Id is not found")
		return
	}

	var updateTodo models.Todo
	// newDecoder => membuat decoder baru untuk membaca data yang ada di body request
	// decoder => decode data json yg ada di r.Body kedalam var updateTodo yang pointernya ke object/struct Todo
	if err := json.NewDecoder(r.Body).Decode(&updateTodo)
	err != nil {
		write(w, "error", http.StatusBadRequest, nil, "invalid input")
		return
	}

	todoMux.Lock()
	defer todoMux.Unlock()

	for i, todo := range todos {
		if todo.ID == id {

			if todos[i].Task == updateTodo.Task{
				write(w, "info", http.StatusOK, nil, "There is no change")
				return
			}

			todos[i].Task = updateTodo.Task
			write(w, "success", http.StatusOK, nil, "Successfully update todo")
			return
		}
	}

	write(w, "error", http.StatusNotFound, nil, "todo not found")
	
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		write(w, "error", http.StatusMethodNotAllowed, nil, "Method not allowed")
		return
	}
	
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		write(w, "error", http.StatusBadRequest, "", "Invalid ID")
		return
	}

	todoMux.Lock()
	defer todoMux.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			todos = remove(todos, i)
			write(w, "success", http.StatusOK, nil, "Successfully deleted todo")
			return
		}
	}
	
	write(w, "error", http.StatusNotFound, nil, "Todo not found")
	
}

func DetailTodo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		write(w, "error", http.StatusMethodNotAllowed, nil, "Method not allowed")
		return
	}

	// menggunakan substring untuk menghapus semua strings yg ter-tera
	idStr := r.URL.Path[len("/api/v1/todos/detail/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		write(w, "error", http.StatusBadRequest, nil, "Invalid ID")
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			write(w, "success", http.StatusOK, todo, "Todo retrieved successfully")
			return
		}
	}
	
	write(w, "error", http.StatusNotFound, nil, "Todo not found")
}

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		write(w, "error", http.StatusMethodNotAllowed, nil, "Method not allowed")
		return
	}

	write(w, "success", http.StatusOK, todos, "List todo retrieved successfully", )
}
