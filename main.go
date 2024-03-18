package main

import (
	"os"
	"todo_backend/db"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
) 
 
var Db db.Database

func main(){ 
	godotenv.Load()
	cfg:= mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "todos",
	}
	Db.InitDb(cfg)
}