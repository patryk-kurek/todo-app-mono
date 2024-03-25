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

var createTableQuery string = `CREATE TABLE IF NOT EXISTS todos (
	id int(11) NOT NULL auto_increment,
	title varchar(250) NOT NULL DEFAULT '0',
	description varchar(500) DEFAULT "No description",
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	completed BOOLEAN DEFAULT 0,
	PRIMARY KEY(id)
);`

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
	r.HandleFunc("/todo/{id}",routes.DeleteTodo).Methods("DELETE");
	addr:=":3001"
	db.Db.InitDb(cfg,createTableQuery)
	if err:= http.ListenAndServe(addr,r); err!= nil{
		log.Fatal(err)
	} 
}
