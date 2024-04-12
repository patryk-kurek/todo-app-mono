package db_test

import (
	"testing"
	"todo_backend/db"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

type dbTestSuite struct { 
	suite.Suite
	cfg mysql.Config
	database db.Database
	createTableQuery string
}

func (s *dbTestSuite) SetupSuite() {
	s.cfg = mysql.Config{ // I shuld create separate db to tests 
		User: "root",
		Passwd: "1234",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "todos_test",
	}
	s.database = db.Database{}
	s.createTableQuery=`CREATE TABLE IF NOT EXISTS todos (
		id int(11) NOT NULL auto_increment,
		title varchar(250) NOT NULL DEFAULT '0',
		description varchar(500) DEFAULT "No description",
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		completed BOOLEAN DEFAULT 0,
		PRIMARY KEY(id)
	);`
}

func (s *dbTestSuite) TestInitDb(){
	s.Run("returns no errors when connecting",func(){
		s.checkIfInitDb()
	})
}

func (s *dbTestSuite) checkIfInitDb(){
	_,err:=s.database.InitDb(s.cfg,s.createTableQuery)
	s.Require().Equal(nil,err,"expected to initialize db");
}

func TestDatabaseTestSuite(t *testing.T){
	suite.Run(t,new(dbTestSuite))
}