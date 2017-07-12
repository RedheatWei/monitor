package mem

type Collect_mem struct {
	Id int64 `xorm:"autoincr notnull unique 'id'"`
	ServerId      int64 `xorm:"notnull 'Serverid'"`
	VirtualMemoryStattotal uint64  `xorm:"notnull 'VirtualMemoryStattotal'"`
	VirtualMemoryStatavailable uint64 `xorm:"notnull 'VirtualMemoryStatavailable'"`
	VirtualMemoryStatused uint64 `xorm:"notnull 'VirtualMemoryStatused'"`
	VirtualMemoryStatusedPercent float64 `xorm:"notnull 'VirtualMemoryStatusedPercent'"`
	SysMemoryStatPercent float64 `xorm:"notnull 'SysMemoryStatPercent'"`
	VirtualMemoryStatfree uint64 `xorm:"notnull 'VirtualMemoryStatfree'"`
	VirtualMemoryStatactive uint64 `xorm:"notnull 'VirtualMemoryStatactive'"`
	VirtualMemoryStatinactive uint64 `xorm:"notnull 'VirtualMemoryStatinactive'"`
	VirtualMemoryStatwired uint64 `xorm:"notnull 'VirtualMemoryStatwired'"`
	VirtualMemoryStatbuffers uint64 `xorm:"notnull 'VirtualMemoryStatbuffers'"`
	VirtualMemoryStatcached uint64 `xorm:"notnull 'VirtualMemoryStatcached'"`
	VirtualMemoryStatwriteback uint64 `xorm:"notnull 'VirtualMemoryStatwriteback'"`
	VirtualMemoryStatdirty uint64 `xorm:"notnull 'VirtualMemoryStatdirty'"`
	VirtualMemoryStatwritebacktmp uint64 `xorm:"notnull 'VirtualMemoryStatwritebacktmp'"`
	VirtualMemoryStatshared uint64 `xorm:"notnull 'VirtualMemoryStatshared'"`
	VirtualMemoryStatslab uint64 `xorm:"notnull 'VirtualMemoryStatslab'"`
	VirtualMemoryStatpagetables uint64 `xorm:"notnull 'VirtualMemoryStatpagetables'"`
	VirtualMemoryStatswapcached uint64 `xorm:"notnull 'VirtualMemoryStatswapcached'"`
	SwapMemoryStattotal uint64 `xorm:"notnull 'SwapMemoryStattotal'"`
	SwapMemoryStatused uint64  `xorm:"notnull 'SwapMemoryStatused'"`
	SwapMemoryStatfree uint64 `xorm:"notnull 'SwapMemoryStatfree'"`
	SwapMemoryStatusedPercent float64 `xorm:"notnull 'SwapMemoryStatusedPercent'"`
	SwapMemoryStatsin uint64 `xorm:"notnull 'SwapMemoryStatsin'"`
	SwapMemoryStatsout uint64 `xorm:"notnull 'SwapMemoryStatsout'"`
	TimeStamp      int64 `xorm:"notnull 'TimeStamp'"`
	CreateTime int64 `xorm:"notnull created 'CreateTime'"`
}