package server

import (
	"log"
	"net/http"
	"todo_backend/db"
	"todo_backend/routes"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var createTableQuery string = `CREATE TABLE IF NOT EXISTS todos_test (
	id int(11) NOT NULL auto_increment,
	title varchar(250) NOT NULL DEFAULT '0',
	description varchar(500) DEFAULT "No description",
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	completed BOOLEAN DEFAULT 0,
	PRIMARY KEY(id)
);`
type Server struct {
	cfg mysql.Config
}

func (s *Server) InitServer(cfg mysql.Config,query string){
	r := mux.NewRouter() 
	r.HandleFunc("/todo",routes.AddTodo).Methods("POST");
	r.HandleFunc("/todo",routes.ReadAllTodos).Methods("GET");
	r.HandleFunc("/todo/{id}",routes.CompleteTodo).Methods("PATCH");
	r.HandleFunc("/todo/{id}",routes.DeleteTodo).Methods("DELETE");
	addr:=":3001"
	db.Db.InitDb(cfg,query)
	if err:= http.ListenAndServe(addr,r); err!=nil{
		log.Fatal(err)
	}
}