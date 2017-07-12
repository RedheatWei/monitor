package info

type Server_info_ip struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId int64 `xorm:"notnull 'ServerId'"`
	Ip string `xorm:"notnull 'Ip'"`
	Type string `xorm:"notnull 'Type'"`
	CreateTime int64 `xorm:"notnull created 'CreateTime'"`
}
type Server_info_base struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	Server string `xorm:"notnull 'Server'"`
	CreateTime int64 `xorm:"notnull created 'CreateTime'"`
}