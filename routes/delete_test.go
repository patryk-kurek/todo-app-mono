package routes_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo_backend/db"
	"todo_backend/routes"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type deleteTestSuite struct {
	suite.Suite
}

func (s *deleteTestSuite) SetupSuite(){
	cfg := mysql.Config{
		User: "root",
		Passwd: "1234",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "todos_test",		
	}
	createTableQuery := `CREATE TABLE IF NOT EXISTS todos (
		id int(11) NOT NULL auto_increment,
		title varchar(250) NOT NULL DEFAULT '0',
		description varchar(500) DEFAULT "No description",
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		completed BOOLEAN DEFAULT 0,
		PRIMARY KEY(id)
	);`
	dropTableQuery := "DROP TABLE todos" 
	addTestDataQuery:= fmt.Sprintf(`
	INSERT INTO 
	todos (title,description,created_at,completed)
	VALUES 
	("%v","%v","%v",%v);
	`,"test","This is testing todo","2023-10-10",0)
	db.Db.InitDb(cfg,dropTableQuery,createTableQuery,addTestDataQuery)
}

func (s *deleteTestSuite) TestDeleteTodoCorrectId(){
	s.Run("return 200",func(){
		id := "1"
		status := s.sendRequest(id)
		s.Require().Equal(200,status)
	})
}

func (s *deleteTestSuite) TestDeleteTodoIncorrectId(){ 
	s.Run("return 500",func(){
		id := "100"
		status := s.sendRequest(id)
		s.Require().Equal(500,status)
	})
}

func (s *deleteTestSuite) TestDeleteTodoNotId(){
	s.Run("return 500",func(){
		id:= "this is not even id" 
		status := s.sendRequest(id)
		s.Require().Equal(500,status)
	})
}

func (s *deleteTestSuite) sendRequest(id string)int{
	w := httptest.NewRecorder()
	req,_ := http.NewRequest("DELETE","/todo/"+id,nil)
	vars := map[string]string {
		"id": id,
	}	
	req = mux.SetURLVars(req,vars)
	req.Header.Set("Content-Type","application/json")
	routes.DeleteTodo(w,req)
	res := w.Result()
	defer res.Body.Close()
	return res.StatusCode
}

func TestDeleteTodosDbSuite(t *testing.T){
	suite.Run(t,new(deleteTestSuite))
}