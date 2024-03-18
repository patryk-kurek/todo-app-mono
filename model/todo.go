package model

import "fmt"
 
type Todo struct {
	Title string `json: "title"`
	Description string `json: "description"`
	Created_at string `json: "created_at"` // name must match!
	Completed int `json: "completed"`
}

func (todo *Todo) GetAddQueryString() string{
	return fmt.Sprintf(`
	INSERT INTO 
	todos (title,description,created_at,completed)
	VALUES 
	("%v","%v","%v",%v);
	`,todo.Title,todo.Description,todo.Created_at,todo.Completed);
}