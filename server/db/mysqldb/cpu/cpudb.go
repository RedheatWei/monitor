package cpu

import (
	"monitor/server/base"
	"fmt"
	db "monitor/server/db/mysqldb"
	"github.com/shirou/gopsutil/cpu"
)

func InsertCpuDB(js base.SysCpuInfo,serverid int64){
	db := db.ConnDB()
	cpu := new(Collect_cpu)
	cpu.ServerId = serverid
	cpu.TimeStamp = js.TimeStamp
	affected, err := db.Insert(cpu)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
	stat_len := len(js.InfoStat)
	if stat_len > 0{
		insertInfoStatDB(js.InfoStat,serverid,js.TimeStamp,stat_len)
	}
}
func insertInfoStatDB(is []cpu.InfoStat,serverid int64,timestamp int64,stat_len int){
	db := db.ConnDB()
	userstats := make([]*Collect_cpu_infostat,stat_len)
	for key,userstat_s := range is{
		userstats[key] = new(Collect_cpu_infostat)
		userstats[key].ServerId = serverid
		userstats[key].InfoStatcpu = userstat_s.CPU
		userstats[key].InfoStatvendorId = userstat_s.VendorID
		userstats[key].InfoStatfamily = userstat_s.Family
		userstats[key].InfoStatmodel = userstat_s.Model
		userstats[key].InfoStatstepping = userstat_s.Stepping
		userstats[key].InfoStatphysicalId = userstat_s.PhysicalID
		userstats[key].InfoStatcoreId = userstat_s.CoreID
		userstats[key].InfoStatcores = userstat_s.Cores
		userstats[key].InfoStatmodelName = userstat_s.ModelName
		userstats[key].InfoStatmhz = userstat_s.Mhz
		userstats[key].InfoStatcacheSize = userstat_s.CacheSize
		userstats[key].InfoStatmicrocode = userstat_s.Microcode
		userstats[key].TimeStamp = timestamp
	}
	affected, err := db.Insert(userstats)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
}