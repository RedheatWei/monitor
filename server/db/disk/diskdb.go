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
}
func insertIOStatDB(io map[string]disk.IOCountersStat,serverid int64,timestamp int64,stat_len int){
	db := db.ConnDB()
	userstats := make(map[string]*Collect_disk_iocountersstat,stat_len)
	for key,userstat_s := range io{
		userstats[key] = new(Collect_disk_iocountersstat)
		userstats[key].IOCountersStatDeviceName = key
		userstats[key].ServerId = serverid
		userstats[key].IOCountersStatname = userstat_s.Name
		userstats[key].TimeStamp = timestamp
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