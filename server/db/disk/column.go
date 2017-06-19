package disk

type Collect_disk struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
	CreateTime int64 `xorm:"notnull created 'CreateTime'"`
}

type Collect_disk_iocountersstat struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	IOCountersStatDeviceName string `xorm:"notnull 'IOCountersStatDeviceName'"`
	IOCountersStatreadCount uint64 `xorm:"notnull 'IOCountersStatreadCount'"`
	IOCountersStatmergedReadCount uint64 `xorm:"notnull 'IOCountersStatmergedReadCount'"`
	IOCountersStatwriteCount uint64 `xorm:"notnull 'IOCountersStatwriteCount'"`
	IOCountersStatmergedWriteCount uint64 `xorm:"notnull 'IOCountersStatmergedWriteCount'"`
	IOCountersStatreadBytes uint64 `xorm:"notnull 'IOCountersStatreadBytes'"`
	IOCountersStatwriteBytes uint64 `xorm:"notnull 'IOCountersStatwriteBytes'"`
	IOCountersStatreadTime uint64 `xorm:"notnull 'IOCountersStatreadTime'"`
	IOCountersStatwriteTime uint64 `xorm:"notnull 'IOCountersStatwriteTime'"`
	IOCountersStatiopsInProgress uint64 `xorm:"notnull 'IOCountersStatiopsInProgress'"`
	IOCountersStatioTime uint64 `xorm:"notnull 'IOCountersStatioTime'"`
	IOCountersStatweightedIO uint64 `xorm:"notnull 'IOCountersStatweightedIO'"`
	IOCountersStatname string `xorm:"notnull 'IOCountersStatname'"`
	IOCountersStatserialNumber string `xorm:"notnull 'IOCountersStatserialNumber'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
}

type Collect_disk_partitionstat struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	PartitionStatdevice  string  `xorm:"notnull 'PartitionStatdevice'"`
	PartitionStatmountpoint  string `xorm:"notnull 'PartitionStatmountpoint'"`
	PartitionStatfstype  string `xorm:"notnull 'PartitionStatfstype'"`
	PartitionStatopts  string `xorm:"notnull 'PartitionStatopts'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
}