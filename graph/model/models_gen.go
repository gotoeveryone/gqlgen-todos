// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DeleteTodo struct {
	ID string `json:"id"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type UpdateTodo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
