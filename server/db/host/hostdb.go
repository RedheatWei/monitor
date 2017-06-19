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
	insertUserStatDB(db,js.UserStat,serverid,js.TimeStamp)
	fmt.Println(affected)
}
func insertUserStatDB(db *xorm.Engine,us []host.UserStat,serverid int64,timestamp int64){
	userstat := new(Collect_host_userstat)
	for _,userstat_s := range us{
		userstat.ServerId = serverid
		userstat.UserStatUser = userstat_s. User
		userstat.UserStatTerminal = userstat_s.Terminal
		userstat.UserStatHost = userstat_s.Host
		userstat.UserStatStarted = userstat_s.Started
		userstat.TimeStamp = timestamp
		affected, err := db.Insert(userstat)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(affected)
	}
}