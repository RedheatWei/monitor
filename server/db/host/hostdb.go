package host

import (
	"github.com/shirou/gopsutil/host"
	"monitor/server/base"
	"fmt"
	"monitor/server/db"
	"github.com/xormplus/xorm"
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
	//insertUserStatDB(db,js.UserStat,serverid,js.TimeStamp)
	//affected, err := db.Insert(js.UserStat)

	fmt.Println(affected)
}
func insertUserStatDB(db *xorm.Engine,us []host.UserStat,serverid int64,timestamp int64){
	userstats := make([]*Collect_host_userstat,1)
	//userstat := new(Collect_host_userstat)
	for key,userstat_s := range us{
		userstats[key].ServerId = serverid
		userstats[key].UserStatUser = userstat_s. User
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