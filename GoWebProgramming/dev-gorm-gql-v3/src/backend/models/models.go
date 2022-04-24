package models

type Todo struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
