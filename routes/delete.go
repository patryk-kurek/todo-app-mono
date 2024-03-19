package routes

import (
	"encoding/json"
	"net/http"
	"todo_backend/db"

	"github.com/gorilla/mux"
)

type responseDelete struct {
	Message string 
}

func DeleteTodo(w http.ResponseWriter,r *http.Request){
	id:=mux.Vars(r)["id"]
	w.Header().Set("Content-Type","application/json")
	message,err := db.Db.DeleteTodo(id)
	if err!=nil{
		response := responseDelete{err.Error()}
		jsonValue,_ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonValue)
	} else {
		response := responseDelete{message}
		jsonValue,_ := json.Marshal(response)
		w.Write(jsonValue)
	}
}