package db

import (
	"database/sql"
	//"github.com/go-sql-driver/mysql"
	"fmt"
)

func ConnDB(){
	db, err := sql.Open("mysql", "root:vv231@/monitor")
	if err != nil{
		fmt.Println(err)
	}
	result, err := db.Exec("SELECT * FROM jvm")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
}