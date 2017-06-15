package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

//func ConnDB() *sql.DB{
//	db, err := sql.Open("mysql", "monitor:monitor@tcp(127.0.0.1:3306)/monitor?charset=utf8")
//	if err != nil{
//		fmt.Println(err)
//	}
//	return db
//}
