package jvm

type Collect_jvm struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId int32 `xorm:"notnull 'ServerId'"`
	LoadedClassCount      int32 `xorm:"notnull 'LoadedClassCount'"`
	UnloadedClassCount    int32 `xorm:"notnull 'UnloadedClassCount'"`
	TotalLoadedClassCount int32 `xorm:"notnull 'TotalLoadedClassCount'"`
	HeapMemoryUsageInit      int64 `xorm:"notnull 'HeapMemoryUsageInit'"`
	HeapMemoryUsageCommitted int64 `xorm:"notnull 'HeapMemoryUsageCommitted'"`
	HeapMemoryUsageMax       int64 `xorm:"notnull 'HeapMemoryUsageMax'"`
	HeapMemoryUsageUsed      int64 `xorm:"notnull 'HeapMemoryUsageUsed'"`
	NonHeapMemoryUsageInit      int64 `xorm:"notnull 'NonHeapMemoryUsageInit'"`
	NonHeapMemoryUsageCommitted int64 `xorm:"notnull 'NonHeapMemoryUsageCommitted'"`
	NonHeapMemoryUsageMax       int64 `xorm:"notnull 'NonHeapMemoryUsageMax'"`
	NonHeapMemoryUsageUsed      int64 `xorm:"notnull 'NonHeapMemoryUsageUsed'"`
	ClassPath string `xorm:"notnull 'ClassPath'"`
	Uptime    int64 `xorm:"notnull 'Uptime'"`
	TotalStartedThreadCount int32 `xorm:"notnull 'TotalStartedThreadCount'"`
	PeakThreadCount         int32 `xorm:"notnull 'PeakThreadCount'"`
	CurrentThreadCpuTime    int64 `xorm:"notnull 'CurrentThreadCpuTime'"`
	ThreadCount             int32 `xorm:"notnull 'ThreadCount'"`
	DaemonThreadCount       int32 `xorm:"notnull 'DaemonThreadCount'"`
	AgentId   string `xorm:"notnull 'AgentId'"`
	TimeStamp int64 `xorm:"notnull 'TimeStamp'"`
	CreateTime int64 `xorm:"notnull created 'CreateTime'"`
}