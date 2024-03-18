package model

import "fmt"
 
type Todo struct {
	Title string `json: "title"`
	Description string `json: "description"`
	CreatedAt string `json: "created_at"`
}

func (todo *Todo) GetAddQueryString() string{
	return fmt.Sprintf(`
	INSERT INTO 
	todos (title,description,created_at)
	VALUES 
	(%v,%v,%v);
	`,todo.Title,todo.Description,todo.CreatedAt);
}