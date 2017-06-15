package jvm

import (
	"monitor/base"
	"fmt"
)

func InsertJvmDB(js base.JvmInfo,addr string){
	//db := ConnDB()
	//defer db.Close()
	//stmt, err := db.Prepare("INSERT INTO collect_jvm(LoadedClassCount,UnloadedClassCount,TotalLoadedClassCount,HeapMemoryUsageInit,HeapMemoryUsageCommitted,HeapMemoryUsageMax,HeapMemoryUsageUsed,NonHeapMemoryUsageInit,NonHeapMemoryUsageCommitted,NonHeapMemoryUsageMax,NonHeapMemoryUsageUsed,ClassPath,Uptime,TotalStartedThreadCount,PeakThreadCount,CurrentThreadCpuTime,ThreadCount,DaemonThreadCount,AgentId,TimeStamp,ServerIp) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//res, err := stmt.Exec(js.LoadedClassCount,js.UnloadedClassCount,js.TotalLoadedClassCount,js.HeapMemoryUsageInit,js.HeapMemoryUsageCommitted,js.HeapMemoryUsageMax,js.HeapMemoryUsageUsed,js.NonHeapMemoryUsageInit,js.NonHeapMemoryUsageCommitted,js.NonHeapMemoryUsageMax,js.NonHeapMemoryUsageUsed,js.ClassPath,js.Uptime,js.TotalStartedThreadCount,js.PeakThreadCount,js.CurrentThreadCpuTime,js.ThreadCount,js.DaemonThreadCount,js.AgentId,js.TimeStamp,addr)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//id,err := res.LastInsertId()
	//if err!=  nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(id)
}
