package routes

import (
	"encoding/json"
	"net/http"
	"todo_backend/db"
	"todo_backend/model"
)

type responseRead struct {
	Message string `json: message`
	Todos []model.Todo `json: todos`
 }

func ReadAllTodos(w http.ResponseWriter,r *http.Request){	
	todos,err:=db.Db.ReadTodos()
	w.Header().Set("Content-Type","application/json")
	if err != nil{
		response:= responseRead{err.Error(),[]model.Todo{}}
		jsonValue,_:=json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonValue)
	} else {
		response:= responseRead{"Succesfully get all todos",todos}
		jsonValue,_ := json.Marshal(response)
		w.Write(jsonValue)
	}	
} 