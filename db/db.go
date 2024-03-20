package db

import (
	"database/sql"
	"fmt"
	"log"
	"todo_backend/model"

	"github.com/go-sql-driver/mysql"
)
 
type Database struct {
	db *sql.DB
}

var Db Database

var createTableQuery string = `CREATE TABLE IF NOT EXISTS todos (
	id int(11) NOT NULL auto_increment,
	title varchar(250) NOT NULL DEFAULT '0',
	description varchar(500) DEFAULT "No description",
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	completed BOOLEAN DEFAULT 0,
	PRIMARY KEY(id)
);`

func (d *Database) InitDb(cfg mysql.Config){
	db,err := sql.Open("mysql",cfg.FormatDSN())
	if err != nil{
		log.Fatal(err)
	}
	pingErr:= db.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!");
	res,err := db.Exec(createTableQuery) 
	if err!= nil{
		log.Fatal(err) 
	} 
	d.db = db
	fmt.Println(res.RowsAffected())
}

func (d *Database) DeleteTodo(id string) (string,error){
	d.checkConnection()
	_,err := d.db.Exec("DELETE FROM todos WHERE id=?;",id)
	if err != nil{
		log.Println(err.Error())
		return err.Error(),err
	}
	return "Succesfully deleted todo with id="+id,nil
}

func (d *Database) MakeTodoCompleted(id string) (string,error){
	d.checkConnection();
	_,err := d.db.Exec("UPDATE todos SET completed=1 WHERE id=?;",id)
	if err != nil{
		log.Println(err.Error())
		return err.Error(),err
	}
	return "Succesfully updated todo with id="+id,nil
}

func (d *Database) AddToDo(element *model.Todo) (string,error){
	d.checkConnection()
	res,err:= d.db.Exec(element.GetAddQueryString()); 
	if err!=nil{
		log.Println(err.Error())
		return err.Error(),err
	}
	row := d.db.QueryRow("SELECT LAST_INSERT_ID()");
	row.Scan(&element.Id) 

	fmt.Println(res.RowsAffected())
	return "Succesfully added data",nil
}

func (d *Database) ReadTodos() ([]model.Todo,error){
	rows, err:= d.db.Query("SELECT * FROM todos;");
	if err != nil {
		log.Println(err.Error())
		return []model.Todo{},err
	}
	var todos []model.Todo

	for rows.Next(){
		var todo model.Todo
		if err:= rows.Scan(&todo.Id,&todo.Title,&todo.Description,&todo.Created_at,&todo.Completed); err!=nil{
			return todos,err
		}
		todos = append(todos, todo)
	}

	return todos,err
}	

func (d *Database) checkConnection(){
	pingErr:=d.db.Ping()
	if pingErr != nil { 
		log.Fatal(pingErr)
	} else {
		log.Println("Connection Stable!");	 
	}
}
