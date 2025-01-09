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

func write(w http.ResponseWriter, status string, code int, task string, message string) {
	var req = models.Response{
		Status:  status,
		Code:    code,
		Task:    task,
		Message: message,
	}

	res, _ := json.Marshal(req)
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
	if r.Method == http.MethodPost { // if the method is POST
		var newTodo models.Todo
		if err := json.NewDecoder(r.Body).Decode(&newTodo);
		err != nil {
			write(w, "error", http.StatusBadRequest,"", "Invalid input")
			return
		}

		// sync.Mutex (lock and unlock) used to avoid race condition
		todoMux.Lock()
		newTodo.ID = nextID
		nextID++;
		todos = append(todos, newTodo)
		todoMux.Unlock()

		write(w, "success", http.StatusCreated, newTodo.Task, "Successfully created new todo")
	} else {
		write(w, "error", http.StatusMethodNotAllowed,"", "Method is not allowed")
	}
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		// menggunakan substring untuk menghapus semua strings selain dari teks yg ter-tera
		idStr := r.URL.Path[len("/api/v1/todo/update/"):]

		id, err := strconv.Atoi(idStr) // convert from string to int
		if err != nil {
			write(w, "error", http.StatusBadRequest, "", "Id is not found")
			return
		}

		var updateTodo models.Todo
		// newDecoder => membuat decoder baru untuk membaca data yang ada di body request
		// decoder => decode data json yg ada di r.Body kedalam var updateTodo yang pointernya ke object/struct Todo
		if err := json.NewDecoder(r.Body).Decode(&updateTodo)
		err != nil {
			write(w, "error", http.StatusBadRequest,"", "invalid input")
			return
		}

		todoMux.Lock()
		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Task = updateTodo.Task
				todoMux.Unlock()
				write(w, "success", http.StatusOK, updateTodo.Task, "Successfully update todo")
				return
			}
		}
		todoMux.Unlock()

		write(w, "error", http.StatusNotFound, "", "todo not found")
	} else {
		write(w, "error", http.StatusMethodNotAllowed, "", "Method not allowed")
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		idStr := r.URL.Path[len("/api/v1/todo/delete/"):]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			write(w, "error", http.StatusBadRequest, "", "Invalid ID")
			return
		}

		todoMux.Lock()
		for i, todo := range todos {
			if todo.ID == id {
				todos = remove(todos, i)
				todoMux.Unlock()
				write(w, "success", http.StatusOK, "", "Successfully deleted todo")
				return
			}
		}
		todoMux.Unlock()

		write(w, "error", http.StatusNotFound, "", "Todo not found")
	} else {
		write(w, "error", http.StatusMethodNotAllowed, "", "Method not allowed")
	}
}

func DetailTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/v1/todo/detail/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		write(w, "error", http.StatusBadRequest, "", "Invalid ID")
		return
	}

	todoMux.Lock()
	for _, todo := range todos {
		if todo.ID == id {
			todoMux.Unlock()
			write(w, "success", http.StatusOK, todo.Task, "Todo retrieved successfully")
			return
		}
	}
	todoMux.Unlock()

	write(w, "error", http.StatusNotFound, "", "Todo not found")
}

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		todoMux.Lock()
		res, _ := json.Marshal(todos)
		todoMux.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		write(w, "error", http.StatusMethodNotAllowed, "", "Method not allowed")
	}
}