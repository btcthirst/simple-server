package models

type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type Mess struct {
	Message string `json:"message"`
}
