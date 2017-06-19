package cpu


type Collect_cpu struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
	CreateTime int64 `xorm:"notnull created 'CreateTime'"`
}

type Collect_cpu_userstat struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	InfoStatcpu int32 `xorm:"notnull 'InfoStatcpu'"`
	InfoStatvendorId string `xorm:"notnull 'InfoStatvendorId'"`
	InfoStatfamily string `xorm:"notnull 'InfoStatfamily'"`
	InfoStatmodel string `xorm:"notnull 'InfoStatmodel'"`
	InfoStatstepping int32 `xorm:"notnull 'InfoStatstepping'"`
	InfoStatphysicalId string `xorm:"notnull 'InfoStatphysicalId'"`
	InfoStatcoreId string `xorm:"notnull 'InfoStatcoreId'"`
	InfoStatcores int32 `xorm:"notnull 'InfoStatcores'"`
	InfoStatmodelName string `xorm:"notnull 'InfoStatmodelName'"`
	InfoStatmhz float64 `xorm:"notnull 'InfoStatmhz'"`
	InfoStatcacheSize int32 `xorm:"notnull 'InfoStatcacheSize'"`
	InfoStatmicrocode string `xorm:"notnull 'InfoStatmicrocode'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
}