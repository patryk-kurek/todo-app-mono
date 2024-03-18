package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
) 

var db *sql.DB

func main(){ 
	godotenv.Load()
	cfg:= mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "todos",
	}
	var err error 
	db, err = sql.Open("mysql",cfg.FormatDSN())
	if err != nil{
		log.Fatal(err)
	}

	pingErr:= db.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}