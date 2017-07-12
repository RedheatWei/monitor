package mem

import (
	"monitor/server/base"
	"fmt"
	"monitor/server/db"
)

func InsertMemDB(js base.SysMemInfo,serverid int64){
	db := db.ConnDB()
	mem := new(Collect_mem)
	mem.ServerId = serverid
	mem.VirtualMemoryStattotal = js.VirtualMemoryStat.Total
	mem.VirtualMemoryStatavailable = js.VirtualMemoryStat.Available
	mem.VirtualMemoryStatused = js.VirtualMemoryStat.Used
	mem.VirtualMemoryStatusedPercent = js.VirtualMemoryStat.UsedPercent
	mem.SysMemoryStatPercent = js.SysMemoryStatPercent
	mem.VirtualMemoryStatfree = js.VirtualMemoryStat.Free
	mem.VirtualMemoryStatactive = js.VirtualMemoryStat.Active
	mem.VirtualMemoryStatinactive = js.VirtualMemoryStat.Inactive
	mem.VirtualMemoryStatwired = js.VirtualMemoryStat.Wired
	mem.VirtualMemoryStatbuffers = js.VirtualMemoryStat.Buffers
	mem.VirtualMemoryStatcached = js.VirtualMemoryStat.Cached
	mem.VirtualMemoryStatwriteback = js.VirtualMemoryStat.Writeback
	mem.VirtualMemoryStatdirty = js.VirtualMemoryStat.Dirty
	mem.VirtualMemoryStatwritebacktmp = js.VirtualMemoryStat.WritebackTmp
	mem.VirtualMemoryStatshared = js.VirtualMemoryStat.Shared
	mem.VirtualMemoryStatslab = js.VirtualMemoryStat.Slab
	mem.VirtualMemoryStatpagetables = js.VirtualMemoryStat.PageTables
	mem.VirtualMemoryStatswapcached = js.VirtualMemoryStat.SwapCached
	mem.SwapMemoryStattotal = js.SwapMemoryStat.Total
	mem.SwapMemoryStatused = js.SwapMemoryStat.Used
	mem.SwapMemoryStatfree = js.SwapMemoryStat.Free
	mem.SwapMemoryStatusedPercent = js.SwapMemoryStat.UsedPercent
	mem.SwapMemoryStatsin = js.SwapMemoryStat.Sin
	mem.SwapMemoryStatsout = js.SwapMemoryStat.Sout
	mem.TimeStamp = js.TimeStamp
	affected, err := db.Insert(mem)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
}