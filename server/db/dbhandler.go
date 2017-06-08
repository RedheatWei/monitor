package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
type JsonInfo struct {
	//class
	LoadedClassCount int32 `json:"LoadedClassCount"`
	UnloadedClassCount int32 `json:"UnloadedClassCount"`
	TotalLoadedClassCount int32 `json:"TotalLoadedClassCount"`
	//memory
	HeapMemoryUsageInit int64 `json:"HeapMemoryUsageInit"`
	HeapMemoryUsageCommitted int64 `json:"HeapMemoryUsageCommitted"`
	HeapMemoryUsageMax int64 `json:"HeapMemoryUsageMax"`
	HeapMemoryUsageUsed int64 `json:"HeapMemoryUsageUsed"`

	NonHeapMemoryUsageInit int64 `json:"NonHeapMemoryUsageInit"`
	NonHeapMemoryUsageCommitted int64 `json:"NonHeapMemoryUsageCommitted"`
	NonHeapMemoryUsageMax int64 `json:"NonHeapMemoryUsageMax"`
	NonHeapMemoryUsageUsed int64 `json:"NonHeapMemoryUsageUsed"`
	//rumtime
	ClassPath string `json:"ClassPath"`
	Uptime int64 `json:"Uptime"`
	//thread
	TotalStartedThreadCount int32 `json:"TotalStartedThreadCount"`
	PeakThreadCount int32 `json:"PeakThreadCount"`
	CurrentThreadCpuTime int64 `json:"CurrentThreadCpuTime"`
	ThreadCount int32 `json:"ThreadCount"`
	DaemonThreadCount int32 `json:"DaemonThreadCount"`
	//base
	AgentId string `json:"agentId"`
	TimeStamp int64 `json:"TimeStamp"`
}

func connDB() *sql.DB{
	db, err := sql.Open("mysql", "root:vv231@unix(/var/lib/mysql/mysql.sock)/monitor?charset=utf8")
	if err != nil{
		fmt.Println(err)
	}
	return db
}
func InsertDB(js JsonInfo,addr string){
	db := connDB()
	stmt, err := db.Prepare("INSERT INTO jvm(LoadedClassCount,UnloadedClassCount,TotalLoadedClassCount,HeapMemoryUsageInit,HeapMemoryUsageCommitted,HeapMemoryUsageMax,HeapMemoryUsageUsed,NonHeapMemoryUsageInit,NonHeapMemoryUsageCommitted,NonHeapMemoryUsageMax,NonHeapMemoryUsageUsed,ClassPath,Uptime,TotalStartedThreadCount,PeakThreadCount,CurrentThreadCpuTime,ThreadCount,DaemonThreadCount,AgentId,TimeStamp,ServerIp) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
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