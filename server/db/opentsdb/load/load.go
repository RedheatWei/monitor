package load

import (
	"monitor/server/base"
	"monitor/server/db/opentsdb"
	"fmt"
	"time"
)

func InsertLoadDB(js base.SysLoadInfo,server string){
	_,err := opentsdb.ConnDb()
	if err != nil{
		fmt.Println(err)
	}
	var load Collect_load
	load.Metric = "collect.sys.server"
	load.Value = server
	load.TimeStamp = time.Now().Unix()
	load.Tags.AvgStatload1 = js.AvgStat.Load1
	load.Tags.AvgStatload5 = js.AvgStat.Load5
	load.Tags.AvgStatload15 = js.AvgStat.Load15
	load.Tags.MiscStatprocsRunning = js.MiscStat.ProcsRunning
	load.Tags.MiscStatprocsBlocked = js.MiscStat.ProcsBlocked
	load.Tags.MiscStatctxt = js.MiscStat.Ctxt
	load.Tags.TimeStamp = js.TimeStamp
	fmt.Println(string(load))
}