package database

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	fmt.Println("masuk init lagi")
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_new")
	if err!= nil{
		log.Fatal(err)
	}
}
