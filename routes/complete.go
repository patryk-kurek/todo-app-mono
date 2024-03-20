package routes

import (
	"encoding/json"
	"net/http"
	"todo_backend/db"

	"github.com/gorilla/mux"
)

type responseUpdate struct {
	Message string 
}

func CompleteTodo(w http.ResponseWriter,r *http.Request){ 
	id:= mux.Vars(r)["id"]
	w.Header().Set("Content-Type","application/json")
	message,err:=db.Db.MakeTodoCompleted(id)
	if err!=nil{
		response := responseUpdate{err.Error()}
		jsonValue,_ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonValue)
	} else {
		response:=responseUpdate{message}
		jsonValue,_:= json.Marshal(response)
		w.Write(jsonValue)
	}
}
