package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo_backend/db"
	"todo_backend/model"
)

type responseAdd struct {
	Message string 
	Todo model.Todo 
}

func AddTodo(w http.ResponseWriter,r *http.Request){
	var todo model.Todo 
	
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	w.Header().Set("Content-type","application/json")

	_,err = db.Db.AddToDo(&todo)
	if err != nil {
		response := responseAdd{err.Error(),model.Todo{}}
		jsonValue,_ := json.Marshal(response)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(jsonValue)
	} else {
		response := responseAdd{"Succesfully created todo!",todo}
		jsonValue,_ := json.Marshal(response)
		w.Write(jsonValue)
	}

}