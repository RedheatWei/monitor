package db
type Server_info_base struct {
	Id int64 `xorm:"autoincr notnull unique"`
	Server      string `xorm:"notnull"`
	CreateTime int64 `xorm:"notnull created"`
}
type Server_info_ip struct {
	Id int64 `xorm:"autoincr notnull unique"`
	Serverid      int64 `xorm:"notnull index"`
	Ip      string `xorm:"notnull index"`
	Type      string `xorm:"notnull"`
	CreateTime int64 `xorm:"notnull created"`
}
