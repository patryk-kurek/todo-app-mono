package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo_backend/db"
	"todo_backend/routes"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

type RequestBody struct {
	Title string `json: "title"`
	Description string `json: "description"`
	Created_at string `json: "created_at"`
}

type addTestSuite struct {
	suite.Suite
}

func (s *addTestSuite) SetupSuite() {
	cfg := mysql.Config{
		User: "root",
		Passwd: "1234",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "todos_test",
	}

	createTableQuery:=`CREATE TABLE IF NOT EXISTS todos (
		id int(11) NOT NULL auto_increment,
		title varchar(250) NOT NULL DEFAULT '0',
		description varchar(500) DEFAULT "No description",
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		completed BOOLEAN DEFAULT 0,
		PRIMARY KEY(id)
	);`

	db.Db.InitDb(cfg,createTableQuery) // init global db I don't think this is the best practice
} 

func (s *addTestSuite) TestAddCorrectData(){ 
	s.Run("return 200",func(){
		body:= RequestBody{"test2","description","2003-10-10"}
		jsonData,_ :=  json.Marshal(body)
		status := s.sendRequest(jsonData)
		s.Require().Equal(200,status)
	})
}

func (s *addTestSuite) TestAddIncorrectData(){
	s.Run("return 422",func() {
		body := RequestBody{"test","description hello","20-0232-0323"}
		jsonData,_ :=  json.Marshal(body)
		status:= s.sendRequest(jsonData)
		s.Require().Equal(422,status)
	})
}

func (s* addTestSuite) TestAddIncorrectJSONFormat(){
	s.Run("return 400",func(){
		body := `{"teststtststs"` 
		status := s.sendRequest([]byte(body))
		s.Require().Equal(400,status)
	})
}

func (s* addTestSuite) sendRequest(body []byte) int{
	w := httptest.NewRecorder()
	req,_ := http.NewRequest("POST","/todo",bytes.NewBuffer(body))
	req.Header.Set("Content-Type","application/json")
	routes.AddTodo(w,req)
	res := w.Result()
	defer res.Body.Close()
	return res.StatusCode
}

func TestAddingDataToDbSuite(t *testing.T){
	suite.Run(t,new(addTestSuite))
}
