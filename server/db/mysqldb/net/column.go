package net

type Collect_net struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
	CreateTime int64 `xorm:"notnull created 'CreateTime'"`
}

type Collect_net_iocountersstat struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	IOCountersStatname string `xorm:"notnull 'IOCountersStatname'"`
	IOCountersStatbytesSent uint64 `xorm:"notnull 'IOCountersStatbytesSent'"`
	IOCountersStatbytesRecv uint64 `xorm:"notnull 'IOCountersStatbytesRecv'"`
	IOCountersStatpacketsSent uint64 `xorm:"notnull 'IOCountersStatpacketsSent'"`
	IOCountersStatpacketsRecv uint64 `xorm:"notnull 'IOCountersStatpacketsRecv'"`
	IOCountersStaterrin uint64 `xorm:"notnull 'IOCountersStaterrin'"`
	IOCountersStaterrout uint64 `xorm:"notnull 'IOCountersStaterrout'"`
	IOCountersStatdropin uint64 `xorm:"notnull 'IOCountersStatdropin'"`
	IOCountersStatdropout uint64 `xorm:"notnull 'IOCountersStatdropout'"`
	IOCountersStatfifoin uint64 `xorm:"notnull 'IOCountersStatfifoin'"`
	IOCountersStatfifoout uint64 `xorm:"notnull 'IOCountersStatfifoout'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
}

type Collect_net_interfacestat struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	InterfaceStatmtu int `xorm:"notnull 'InterfaceStatmtu'"`
	InterfaceStatname string `xorm:"notnull 'InterfaceStatname'"`
	InterfaceStathardwareaddr string `xorm:"notnull 'InterfaceStathardwareaddr'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
}
type Collect_net_interfacestat_addrs struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	Addrsaddr string `xorm:"notnull 'Addrsaddr'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
}
