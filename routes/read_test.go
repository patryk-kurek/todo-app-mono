package routes_test

import (
	"fmt"
	"testing"
	"todo_backend/db"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

type readTestSuite struct {
	suite.Suite
}

func (*readTestSuite) SetupSuite(){
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
	
	dropTableQuery:=`DROP TABLE todos;`
	addTestDataQuery:= fmt.Sprintf(`
	INSERT INTO 
	todos (title,description,created_at,completed)
	VALUES 
	("%v","%v","%v",%v);
	`,"test","This is testing todo","2023-10-10",0)

	db.Db.InitDb(cfg,dropTableQuery,createTableQuery,addTestDataQuery)
}

func TestReadTodosSuite(t *testing.T){
	suite.Run(t,new(readTestSuite))
}
 