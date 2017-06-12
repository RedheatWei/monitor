package models

import (
	"github.com/astaxie/beego/orm"
)

type Jvm struct {
	Id int32
	LoadedClassCount int32
	UnloadedClassCount int32
	TotalLoadedClassCount int32
	HeapMemoryUsageInit int64
	HeapMemoryUsageCommitted int64
	HeapMemoryUsageMax int64
	HeapMemoryUsageUsed int64
	NonHeapMemoryUsageInit int64
	NonHeapMemoryUsageCommitted int64
	NonHeapMemoryUsageMax int64
	NonHeapMemoryUsageUsed int64
	ClassPath string
	Uptime int64
	TotalStartedThreadCount int32
	PeakThreadCount int32
	CurrentThreadCpuTime int64
	ThreadCount int32
	DaemonThreadCount int32
	AgentId string
	TimeStamp int64
	ServerIp string
}

func init(){
	orm.RegisterDataBase("monitor", "mysql", "root:vv231@unix(/var/lib/mysql/mysql.sock)/monitor?charset=utf8", 30)
	orm.RegisterModel(new(Jvm))
}
func Sel() RawSeter{
	o := orm.NewOrm()
	r := o.Raw("select * from jvm")
	return  r
}
