// models for define struct and contract/interface
// or we can define a new type data
package models

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

type Response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Task    interface{} `json:"data,omitempty"` // omitempty -> datanya tidak tampil jika value = nil
	Message string `json:"message"`
}
