package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"todo_backend/model"

	"github.com/go-sql-driver/mysql"
)
 
type Database struct {
	db *sql.DB
}

var Db Database

func (d *Database) InitDb(cfg mysql.Config,queries ...string) (string,error){
	db,err := sql.Open("mysql",cfg.FormatDSN())
	if err != nil{
		return err.Error(),err
	}
	pingErr:= db.Ping()
	if pingErr != nil{
		return err.Error(),err
	}
	fmt.Println("Connected!");
	var res sql.Result
	for _,query := range queries{
		res,err = db.Exec(query)
		if err != nil{
			return err.Error(),err
		}
	}
	d.db = db
	fmt.Println(res.RowsAffected())
	return "succesfully initialized db",nil
}

func (d *Database) DeleteTodo(id string) (string,error){
	d.checkConnection()
	out,err := d.db.Exec("DELETE FROM todos WHERE id=?;",id)
	if err != nil{
		log.Println(err.Error())
		return err.Error(),err
	}
	affectedRows,_ := out.RowsAffected()
	if affectedRows == 0 {
		nilAffectedRowsErr := errors.New("Nil Affected Rows!")
		log.Println(nilAffectedRowsErr.Error())
		return nilAffectedRowsErr.Error(),nilAffectedRowsErr
	}
	return "Succesfully deleted todo with id="+id,nil
}

func (d *Database) MakeTodoCompleted(id string) (string,error){
	d.checkConnection();
	out,err := d.db.Exec("UPDATE todos SET completed=1 WHERE id=?;",id)
	if err != nil{
		log.Println(err.Error())
		return err.Error(),err
	} 
	affectedRows,_ := out.RowsAffected()
	if affectedRows == 0 {
		nilAffectedRowsErr := errors.New("Nil Affected Rows!")
		log.Println(nilAffectedRowsErr.Error())
		return nilAffectedRowsErr.Error(),nilAffectedRowsErr
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
