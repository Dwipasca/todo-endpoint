package models

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

type Response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Task    string `json:"task"`
	Message string `json:"message"`
}
