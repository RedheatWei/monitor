package host

type Collect_host struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	InfoStathostname string `xorm:"notnull 'InfoStathostname'"`
	InfoStatuptime uint64 `xorm:"notnull 'InfoStatuptime'"`
	InfoStatbootTime uint64 `xorm:"notnull 'InfoStatbootTime'"`
	InfoStatprocs uint64 `xorm:"notnull 'InfoStatprocs'"`
	InfoStatos string `xorm:"notnull 'InfoStatos'"`
	InfoStatplatform string `xorm:"notnull 'InfoStatplatform'"`
	InfoStatplatformFamily string `xorm:"notnull 'InfoStatplatformFamily'"`
	InfoStatplatformVersion string `xorm:"notnull 'InfoStatplatformVersion'"`
	InfoStatkernelVersion string `xorm:"notnull 'InfoStatkernelVersion'"`
	InfoStatvirtualizationSystem string `xorm:"notnull 'InfoStatvirtualizationSystem'"`
	InfoStatvirtualizationRole string `xorm:"notnull 'InfoStatvirtualizationRole'"`
	InfoStathostid string `xorm:"notnull 'InfoStathostid'"`
	//TemperatureStat
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
	CreateTime int64 `xorm:"notnull created 'CreateTime'"`
}

type Collect_host_userstat struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	UserStatUser string `xorm:"notnull 'UserStatUser'"`
	UserStatTerminal string `xorm:"notnull 'UserStatTerminal'"`
	UserStatHost string `xorm:"notnull 'UserStatHost'"`
	UserStatStarted int `xorm:"notnull 'UserStatStarted'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
}