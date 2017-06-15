package jvm

type Collect_jvm struct {
	Id int64 `xorm:"autoincr notnull unique"`
	LoadedClassCount      int32 `xorm:"notnull"`
	UnloadedClassCount    int32 `xorm:"notnull"`
	TotalLoadedClassCount int32 `xorm:"notnull"`
	HeapMemoryUsageInit      int64 `xorm:"notnull"`
	HeapMemoryUsageCommitted int64 `xorm:"notnull"`
	HeapMemoryUsageMax       int64 `xorm:"notnull"`
	HeapMemoryUsageUsed      int64 `xorm:"notnull"`
	NonHeapMemoryUsageInit      int64 `xorm:"notnull"`
	NonHeapMemoryUsageCommitted int64 `xorm:"notnull"`
	NonHeapMemoryUsageMax       int64 `xorm:"notnull"`
	NonHeapMemoryUsageUsed      int64 `xorm:"notnull"`
	ClassPath string `xorm:"notnull"`
	Uptime    int64 `xorm:"notnull"`
	TotalStartedThreadCount int32 `xorm:"notnull"`
	PeakThreadCount         int32 `xorm:"notnull"`
	CurrentThreadCpuTime    int64 `xorm:"notnull"`
	ThreadCount             int32 `xorm:"notnull"`
	DaemonThreadCount       int32 `xorm:"notnull"`
	AgentId   string `xorm:"notnull"`
	TimeStamp int64 `xorm:"notnull"`
	CreateTime int64 `xorm:"notnull created"`
	ServerIp string `xorm:"notnull"`
}