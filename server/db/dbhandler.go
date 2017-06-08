package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func ConnDB(){
	db, err := sql.Open("mysql", "root:vv231@unix(/var/lib/mysql/mysql.sock)/monitor?charset=utf8")
	defer db.Close()
	if err != nil{
		fmt.Println(err)
	}
	result, err := db.Exec("SELECT * FROM jvm")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
}