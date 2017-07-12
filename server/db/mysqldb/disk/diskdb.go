package disk

import (
	"monitor/server/base"
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"monitor/server/db"
)

func InsertDiskDB(js base.SysDiskInfo,serverid int64){
	db := db.ConnDB()
	net := new(Collect_disk)
	net.ServerId = serverid
	net.TimeStamp = js.TimeStamp
	affected, err := db.Insert(net)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
	stat_len := len(js.IOCountersStat)
	if stat_len > 0{
		insertIOStatDB(js.IOCountersStat,serverid,js.TimeStamp,stat_len)
	}
	stat_len2 := len(js.PartitionStat)
	if stat_len2 > 0{
		insertPartitionStatDB(js.PartitionStat,serverid,js.TimeStamp,stat_len2)
	}
	stat_len3 := len(js.UsageStat)
	if stat_len3 > 0 {
		insertUsageStatDB(js.UsageStat,serverid,js.TimeStamp,stat_len3)
	}
}
func insertIOStatDB(io map[string]disk.IOCountersStat,serverid int64,timestamp int64,stat_len int){
	db := db.ConnDB()
	userstats := make([]*Collect_disk_iocountersstat,stat_len)
	var i int = 0
	for key,userstat_s := range io{
		userstats[i] = new(Collect_disk_iocountersstat)
		userstats[i].ServerId = serverid
		userstats[i].IOCountersStatDeviceName = key
		userstats[i].IOCountersStatreadCount = userstat_s.ReadCount
		userstats[i].IOCountersStatmergedReadCount =  userstat_s.MergedReadCount
		userstats[i].IOCountersStatwriteCount =  userstat_s.WriteCount
		userstats[i].IOCountersStatmergedWriteCount =  userstat_s.MergedWriteCount
		userstats[i].IOCountersStatreadBytes =  userstat_s.ReadBytes
		userstats[i].IOCountersStatwriteBytes =  userstat_s.WriteBytes
		userstats[i].IOCountersStatreadTime =  userstat_s.ReadTime
		userstats[i].IOCountersStatwriteTime =  userstat_s.WriteTime
		userstats[i].IOCountersStatiopsInProgress =  userstat_s.IopsInProgress
		userstats[i].IOCountersStatioTime =  userstat_s.IoTime
		userstats[i].IOCountersStatweightedIO =  userstat_s.WeightedIO
		userstats[i].IOCountersStatname =  userstat_s.Name
		userstats[i].IOCountersStatserialNumber =  userstat_s.SerialNumber
		userstats[i].TimeStamp = timestamp
		i++
	}
	affected, err := db.Insert(userstats)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
}
func insertPartitionStatDB(partition []disk.PartitionStat,serverid int64,timestamp int64,stat_len int){
	db := db.ConnDB()
	userstats := make([]*Collect_disk_partitionstat,stat_len)
	for key,userstat_s := range partition{
		userstats[key] = new(Collect_disk_partitionstat)
		userstats[key].ServerId = serverid
		userstats[key].PartitionStatdevice = userstat_s.Device
		userstats[key].PartitionStatmountpoint = userstat_s.Mountpoint
		userstats[key].PartitionStatfstype = userstat_s.Fstype
		userstats[key].PartitionStatopts = userstat_s.Opts
		userstats[key].TimeStamp = timestamp
	}
	affected, err := db.Insert(userstats)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
}
func insertUsageStatDB(us []disk.UsageStat,serverid int64,timestamp int64,stat_len int){
	db := db.ConnDB()
	usagestats := make([]*Collect_disk_usagestat,stat_len)
	for key,userstat_s := range us{
		usagestats[key] = new(Collect_disk_usagestat)
		usagestats[key].ServerId = serverid
		usagestats[key].UsageStatPath = userstat_s.Path
		usagestats[key].UsageStatFstype = userstat_s.Fstype
		usagestats[key].UsageStatTotal = userstat_s.Total
		usagestats[key].UsageStatFree = userstat_s.Free
		usagestats[key].UsageStatUsed = userstat_s.Used
		usagestats[key].UsageStatUsedPercent = userstat_s.UsedPercent
		usagestats[key].UsageStatInodesTotal = userstat_s.InodesTotal
		usagestats[key].UsageStatInodesUsed = userstat_s.InodesUsed
		usagestats[key].UsageStatInodesFree = userstat_s.InodesFree
		usagestats[key].UsageStatInodesUsedPercent = userstat_s.InodesUsedPercent
		usagestats[key].TimeStamp = timestamp
	}
	affected, err := db.Insert(usagestats)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)

}