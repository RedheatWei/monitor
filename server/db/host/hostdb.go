package host

import (
	"github.com/shirou/gopsutil/host"
	"monitor/server/base"
	"fmt"
	"monitor/server/db"
)
func InsertHostDB(js base.SysHostInfo,serverid int64){
	db := db.ConnDB()
	host := new(Collect_host)
	host.ServerId = serverid
	host.InfoStathostname = js.InfoStat.Hostname
	host.InfoStatuptime = js.InfoStat.Uptime
	host.InfoStatbootTime = js.InfoStat.BootTime
	host.InfoStatprocs = js.InfoStat.Procs
	host.InfoStatos = js.InfoStat.OS
	host.InfoStatplatform = js.InfoStat.Platform
	host.InfoStatplatformFamily = js.InfoStat.PlatformFamily
	host.InfoStatplatformVersion = js.InfoStat.PlatformVersion
	host.InfoStatkernelVersion = js.InfoStat.KernelVersion
	host.InfoStatvirtualizationSystem = js.InfoStat.VirtualizationSystem
	host.InfoStatvirtualizationRole = js.InfoStat.VirtualizationRole
	host.InfoStathostid = js.InfoStat.HostID
	host.TimeStamp = js.TimeStamp
	affected, err := db.Insert(host)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
	user_stat_len := len(js.UserStat)
	fmt.Println(user_stat_len)
	if user_stat_len > 0{
		insertUserStatDB(js.UserStat,serverid,js.TimeStamp,user_stat_len)
	}
}
func insertUserStatDB(us []host.UserStat,serverid int64,timestamp int64,user_stat_len int){
	db := db.ConnDB()
	userstats := make([]*Collect_host_userstat,user_stat_len)
	//userstat := new(Collect_host_userstat)
	for key,userstat_s := range us{
		//fmt.Println("key:",key)
		userstats[key] = new(Collect_host_userstat)
		userstats[key].ServerId = serverid
		userstats[key].UserStatUser = userstat_s.User
		userstats[key].UserStatTerminal = userstat_s.Terminal
		userstats[key].UserStatHost = userstat_s.Host
		userstats[key].UserStatStarted = userstat_s.Started
		userstats[key].TimeStamp = timestamp
	}
	affected, err := db.Insert(userstats)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
}