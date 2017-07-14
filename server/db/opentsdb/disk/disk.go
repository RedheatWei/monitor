package disk

import (
	"monitor/server/base"
	"fmt"
	"time"
	"reflect"
	"encoding/json"
	"monitor/server/db/opentsdb"
	"github.com/shirou/gopsutil/disk"
)
//定义收集的项目
type diskDB struct {
	UsageStatTotal *opentsdb.Collect_disk
	UsageStatFree *opentsdb.Collect_disk
	UsageStatUsed *opentsdb.Collect_disk
	UsageStatUsedPercent *opentsdb.Collect_disk
	UsageStatInodesTotal *opentsdb.Collect_disk
	UsageStatInodesUsed *opentsdb.Collect_disk
	UsageStatInodesFree *opentsdb.Collect_disk
	UsageStatInodesUsedPercent *opentsdb.Collect_disk
}
//主函数
func InsertDiskDB(js base.SysDiskInfo,serverIpInfo base.ServerIpInfo){
	var diskDB diskDB
	//var sli_all_disk [][]interface{}
	var sli_str []interface{}
	for _,usage_stat := range js.UsageStat{
		diskDB.UsageStatTotal = usageStatTotal(js,serverIpInfo,usage_stat)
		diskDB.UsageStatFree = usageStatFree(js,serverIpInfo,usage_stat)
		diskDB.UsageStatUsed = usageStatUsed(js,serverIpInfo,usage_stat)
		diskDB.UsageStatUsedPercent = usageStatUsedPercent(js,serverIpInfo,usage_stat)
		diskDB.UsageStatInodesTotal = usageStatInodesTotal(js,serverIpInfo,usage_stat)
		diskDB.UsageStatInodesUsed = usageStatInodesUsed(js,serverIpInfo,usage_stat)
		diskDB.UsageStatInodesFree = usageStatInodesFree(js,serverIpInfo,usage_stat)
		diskDB.UsageStatInodesUsedPercent = usageStatInodesUsedPercent(js,serverIpInfo,usage_stat)
		v := reflect.ValueOf(diskDB)
		for k := 0; k < v.NumField(); k++{
			if ! v.Field(k).IsNil(){
				val := v.Field(k).Interface()
				sli_str = append(sli_str,val)
			}
		}
		//sli_all_disk = append(sli_all_disk,sli_str)
	}
	b,err := json.Marshal(sli_str)
	if err!=nil{
		fmt.Println(err)
	}
	opentsdb.SendToTsDb(string(b))
}
//组合数据,便于修改tag
func createCollect(js base.SysDiskInfo,serverIpInfo base.ServerIpInfo,usage_stat disk.UsageStat) *opentsdb.Collect_disk{
	collect := new(opentsdb.Collect_disk)
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = serverIpInfo.Server
	collect.Tags.Ip = serverIpInfo.Ip
	collect.Tags.CtimeStamp = js.TimeStamp
	collect.Tags.UsageStatPath = usage_stat.Path
	return collect
}

func usageStatTotal(js base.SysDiskInfo,serverIpInfo base.ServerIpInfo,usage_stat disk.UsageStat) *opentsdb.Collect_disk{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.disk.usageStatTotal"
	collect.Value = float64(usage_stat.Total)
	return collect
}
func usageStatFree(js base.SysDiskInfo,serverIpInfo base.ServerIpInfo,usage_stat disk.UsageStat) *opentsdb.Collect_disk{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.disk.usageStatFree"
	collect.Value = float64(usage_stat.Free)
	return collect
}
func usageStatUsed(js base.SysDiskInfo,serverIpInfo base.ServerIpInfo,usage_stat disk.UsageStat) *opentsdb.Collect_disk{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.disk.usageStatUsed"
	collect.Value = float64(usage_stat.Used)
	return collect
}
func usageStatUsedPercent(js base.SysDiskInfo,serverIpInfo base.ServerIpInfo,usage_stat disk.UsageStat) *opentsdb.Collect_disk{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.disk.usageStatUsedPercent"
	collect.Value = float64(usage_stat.UsedPercent)
	return collect
}
func usageStatInodesTotal(js base.SysDiskInfo,serverIpInfo base.ServerIpInfo,usage_stat disk.UsageStat) *opentsdb.Collect_disk{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.disk.usageStatInodesTotal"
	collect.Value = float64(usage_stat.InodesTotal)
	return collect
}
func usageStatInodesUsed(js base.SysDiskInfo,serverIpInfo base.ServerIpInfo,usage_stat disk.UsageStat) *opentsdb.Collect_disk{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.disk.usageStatInodesUsed"
	collect.Value = float64(usage_stat.InodesUsed)
	return collect
}
func usageStatInodesFree(js base.SysDiskInfo,serverIpInfo base.ServerIpInfo,usage_stat disk.UsageStat) *opentsdb.Collect_disk{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.disk.usageStatInodesFree"
	collect.Value = float64(usage_stat.InodesFree)
	return collect
}
func usageStatInodesUsedPercent(js base.SysDiskInfo,serverIpInfo base.ServerIpInfo,usage_stat disk.UsageStat) *opentsdb.Collect_disk{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.disk.usageStatInodesUsedPercent"
	collect.Value = float64(usage_stat.InodesUsedPercent)
	return collect
}