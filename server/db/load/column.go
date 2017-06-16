package load

type Collect_load struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	AvgStatload1      float64 `xorm:"notnull 'AvgStatload1'"`
	AvgStatload5      float64 `xorm:"notnull 'AvgStatload5'"`
	AvgStatload15      float64 `xorm:"notnull 'AvgStatload15'"`
	MiscStatprocsRunning      int `xorm:"notnull 'MiscStatprocsRunning'"`
	MiscStatprocsBlocked      int `xorm:"notnull 'MiscStatprocsBlocked'"`
	MiscStatctxt      int `xorm:"notnull 'MiscStatctxt'"`
	CreateTime int64 `xorm:"notnull created 'CreateTime'"`
}

