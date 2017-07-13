package mem

import (
	"monitor/server/db/opentsdb"
	"time"
	"reflect"
	"encoding/json"
	"fmt"
	"monitor/server/base"
)
//定义收集的项目
type memDB struct {
	virtualMemoryStattotal *opentsdb.Collect
	virtualMemoryStatavailable *opentsdb.Collect
	virtualMemoryStatused *opentsdb.Collect
	virtualMemoryStatusedPercent *opentsdb.Collect
	sysMemoryStatPercent *opentsdb.Collect
	virtualMemoryStatfree *opentsdb.Collect
	virtualMemoryStatactive *opentsdb.Collect
	virtualMemoryStatinactive *opentsdb.Collect
	virtualMemoryStatwired *opentsdb.Collect
	virtualMemoryStatbuffers *opentsdb.Collect
	virtualMemoryStatcached *opentsdb.Collect
	virtualMemoryStatwriteback *opentsdb.Collect
	virtualMemoryStatdirty *opentsdb.Collect
	virtualMemoryStatwritebacktmp *opentsdb.Collect
	virtualMemoryStatshared *opentsdb.Collect
	virtualMemoryStatslab *opentsdb.Collect
	virtualMemoryStatpagetables *opentsdb.Collect
	virtualMemoryStatswapcached *opentsdb.Collect
	swapMemoryStattotal *opentsdb.Collect
	swapMemoryStatused *opentsdb.Collect
	swapMemoryStatfree *opentsdb.Collect
	swapMemoryStatusedPercent *opentsdb.Collect
	swapMemoryStatsin *opentsdb.Collect
	swapMemoryStatsout *opentsdb.Collect
}
//主函数
func InsertMemDB(js base.SysMemInfo,server string){
	var memDB memDB
	memDB.virtualMemoryStattotal = virtualMemoryStattotal(js,server)
	memDB.virtualMemoryStatavailable = virtualMemoryStatavailable(js,server)
	memDB.virtualMemoryStatused = virtualMemoryStatused(js,server)
	memDB.virtualMemoryStatusedPercent = virtualMemoryStatusedPercent(js,server)
	memDB.sysMemoryStatPercent = sysMemoryStatPercent(js,server)
	memDB.virtualMemoryStatfree = virtualMemoryStatfree(js,server)
	memDB.virtualMemoryStatactive = virtualMemoryStatactive(js,server)
	memDB.virtualMemoryStatinactive = virtualMemoryStatinactive(js,server)
	memDB.virtualMemoryStatwired = virtualMemoryStatwired(js,server)
	memDB.virtualMemoryStatbuffers = virtualMemoryStatbuffers(js,server)
	memDB.virtualMemoryStatcached = virtualMemoryStatcached(js,server)
	memDB.virtualMemoryStatwriteback = virtualMemoryStatwriteback(js,server)
	memDB.virtualMemoryStatdirty = virtualMemoryStatdirty(js,server)
	memDB.virtualMemoryStatwritebacktmp = virtualMemoryStatwritebacktmp(js,server)
	memDB.virtualMemoryStatshared = virtualMemoryStatshared(js,server)
	memDB.virtualMemoryStatslab = virtualMemoryStatslab(js,server)
	memDB.virtualMemoryStatpagetables = virtualMemoryStatpagetables(js,server)
	memDB.virtualMemoryStatswapcached = virtualMemoryStatswapcached(js,server)
	memDB.swapMemoryStattotal = swapMemoryStattotal(js,server)
	memDB.swapMemoryStatused = swapMemoryStatused(js,server)
	memDB.swapMemoryStatfree = swapMemoryStatfree(js,server)
	memDB.swapMemoryStatusedPercent = swapMemoryStatusedPercent(js,server)
	memDB.swapMemoryStatsin = swapMemoryStatsin(js,server)
	memDB.swapMemoryStatsout = swapMemoryStatsout(js,server)
	v := reflect.ValueOf(memDB)
	var sli_str []interface{}
	for k := 0; k < v.NumField(); k++{
		if ! v.Field(k).IsNil(){
			val := v.Field(k).Interface()
			sli_str = append(sli_str,val)
		}
	}
	b,err := json.Marshal(sli_str)
	if err!=nil{
		fmt.Println(err)
	}
	opentsdb.SendToTsDb(string(b))
}
//组合数据,便于修改tag
func createCollect(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := new(opentsdb.Collect)
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = server
	collect.Tags.CtimeStamp = js.TimeStamp
	return collect
}

func virtualMemoryStattotal(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStattotal"
	collect.Value = float64(js.VirtualMemoryStat.Total)
	return collect
}
func virtualMemoryStatavailable(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatavailable"
	collect.Value = float64(js.VirtualMemoryStat.Available)
	return collect
}
func virtualMemoryStatused(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatused"
	collect.Value = float64(js.VirtualMemoryStat.Used)
	return collect
}
func virtualMemoryStatusedPercent(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatusedPercent"
	collect.Value = js.VirtualMemoryStat.UsedPercent
	return collect
}
func sysMemoryStatPercent(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.sysMemoryStatPercent"
	collect.Value = js.SysMemoryStatPercent
	return collect
}
func virtualMemoryStatfree(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatfree"
	collect.Value = float64(js.VirtualMemoryStat.Free)
	return collect
}
func virtualMemoryStatactive(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatactive"
	collect.Value = float64(js.VirtualMemoryStat.Active)
	return collect
}
func virtualMemoryStatinactive(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.1m"
	collect.Value = float64(js.VirtualMemoryStat.Active)
	return collect
}
func virtualMemoryStatwired(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatwired"
	collect.Value = float64(js.VirtualMemoryStat.Wired)
	return collect
}
func virtualMemoryStatbuffers(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatbuffers"
	collect.Value = float64(js.VirtualMemoryStat.Buffers)
	return collect
}
func virtualMemoryStatcached(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatcached"
	collect.Value = float64(js.VirtualMemoryStat.Cached)
	return collect
}
func virtualMemoryStatwriteback(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatwriteback"
	collect.Value = float64(js.VirtualMemoryStat.Writeback)
	return collect
}
func virtualMemoryStatdirty(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatdirty"
	collect.Value = float64(js.VirtualMemoryStat.Dirty)
	return collect
}
func virtualMemoryStatwritebacktmp(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatwritebacktmp"
	collect.Value = float64(js.VirtualMemoryStat.WritebackTmp)
	return collect
}
func virtualMemoryStatshared(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatshared"
	collect.Value = float64(js.VirtualMemoryStat.Shared)
	return collect
}
func virtualMemoryStatslab(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatslab"
	collect.Value = float64(js.VirtualMemoryStat.Slab)
	return collect
}
func virtualMemoryStatpagetables(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatpagetables"
	collect.Value = float64(js.VirtualMemoryStat.PageTables)
	return collect
}func virtualMemoryStatswapcached(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.mem.virtualMemoryStatswapcached"
	collect.Value = float64(js.VirtualMemoryStat.SwapCached)
	return collect
}
func swapMemoryStattotal(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.swap.swapMemoryStattotal"
	collect.Value = float64(js.SwapMemoryStat.Total)
	return collect
}
func swapMemoryStatused(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.swap.swapMemoryStatused"
	collect.Value = float64(js.SwapMemoryStat.Used)
	return collect
}
func swapMemoryStatfree(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.swap.swapMemoryStatfree"
	collect.Value = float64(js.SwapMemoryStat.Free)
	return collect
}
func swapMemoryStatusedPercent(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.swap.swapMemoryStatusedPercent"
	collect.Value = float64(js.SwapMemoryStat.UsedPercent)
	return collect
}
func swapMemoryStatsin(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.swap.swapMemoryStatsin"
	collect.Value = float64(js.SwapMemoryStat.Sin)
	return collect
}
func swapMemoryStatsout(js base.SysMemInfo,server string) *opentsdb.Collect{
	collect := createCollect(js,server)
	collect.Metric = "sys.swap.swapMemoryStatsout"
	collect.Value = float64(js.SwapMemoryStat.Sout)
	return collect
}