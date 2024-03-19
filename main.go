package main

import (
	"log"
	"net/http"
	"os"
	"todo_backend/db"
	"todo_backend/routes"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
) 

func main(){ 
	godotenv.Load()
	cfg:= mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "todos",
	}
	r:= mux.NewRouter() 
	r.HandleFunc("/todo",routes.AddTodo).Methods("POST");
	r.HandleFunc("/todo",routes.ReadAllTodos).Methods("GET");
	r.HandleFunc("/todo/{id}",routes.CompleteTodo).Methods("PATCH");
	addr:=":3001"
	db.Db.InitDb(cfg)
	if err:= http.ListenAndServe(addr,r); err!= nil{
		log.Fatal(err)
	} 
}