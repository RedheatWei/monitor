package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net"
	"monitor/base"
)
func connDB() *sql.DB{
	db, err := sql.Open("mysql", "root:vv231@unix(/var/lib/mysql/mysql.sock)/monitor?charset=utf8")
	if err != nil{
		fmt.Println(err)
	}
	return db
}
func InsertDB(js base.JsonInfo,addr *net.UDPAddr){
	db := connDB()
	stmt, err := db.Prepare("INSERT INTO jvm(LoadedClassCount,UnloadedClassCount,TotalLoadedClassCount,HeapMemoryUsageInit,HeapMemoryUsageCommitted,HeapMemoryUsageMax,HeapMemoryUsageUsed,NonHeapMemoryUsageInit,NonHeapMemoryUsageCommitted,NonHeapMemoryUsageMax,NonHeapMemoryUsageUsed,ClassPath,Uptime,TotalStartedThreadCount,PeakThreadCount,CurrentThreadCpuTime,ThreadCount,DaemonThreadCount,AgentId,TimeStamp,ServerIp) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil{
		fmt.Println(err)
	}
	res, err := stmt.Exec(js.LoadedClassCount,js.UnloadedClassCount,js.TotalLoadedClassCount,js.HeapMemoryUsageInit,js.HeapMemoryUsageCommitted,js.HeapMemoryUsageMax,js.HeapMemoryUsageUsed,js.NonHeapMemoryUsageInit,js.NonHeapMemoryUsageCommitted,js.NonHeapMemoryUsageMax,js.NonHeapMemoryUsageUsed,js.ClassPath,js.Uptime,js.TotalStartedThreadCount,js.PeakThreadCount,js.CurrentThreadCpuTime,js.ThreadCount,js.DaemonThreadCount,js.AgentId,js.TimeStamp,addr)
	if err != nil{
		fmt.Println(err)
	}
	id,err := res.LastInsertId()
	if err!=  nil{
		fmt.Println(err)
	}
	fmt.Println(id)
}