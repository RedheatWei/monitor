package load

import (
	"monitor/server/base"
	"fmt"
	"time"
	"reflect"
	"encoding/json"
	"monitor/server/db/opentsdb"
)
//定义收集的项目
type loadDB struct {
	AvgStatload1 *opentsdb.Collect_load
	AvgStatload5 *opentsdb.Collect_load
	AvgStatload15 *opentsdb.Collect_load
	MiscStatprocsRunning *opentsdb.Collect_load
	MiscStatprocsBlocked *opentsdb.Collect_load
	MiscStatctxt *opentsdb.Collect_load
}
//主函数
func InsertLoadDB(js base.SysLoadInfo,server string){
	var loadDB loadDB
	loadDB.AvgStatload1 = load1(js,server)
	loadDB.AvgStatload5 = load5(js,server)
	loadDB.AvgStatload15 = load15(js,server)
	loadDB.MiscStatprocsRunning = procsRunning(js,server)
	loadDB.MiscStatprocsBlocked = procsBlocked(js,server)
	loadDB.MiscStatctxt = ctxt(js,server)
	v := reflect.ValueOf(loadDB)
	for k := 0; k < v.NumField(); k++{
		val := v.Field(k).Interface()
		b,err := json.Marshal(val)
		if err!=nil{
			fmt.Println(err)
		}
		str := string(b)
		fmt.Println(str)
		if str != "null"{
			opentsdb.SendToTsDb(str)
		}
	}
}
//1分钟load
func load1(js base.SysLoadInfo,server string) *opentsdb.Collect_load{
	load := new(opentsdb.Collect_load)
	load.Metric = "sys.load.1m"
	load.Value = js.AvgStat.Load1
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
//5分钟load
func load5(js base.SysLoadInfo,server string) *opentsdb.Collect_load{
	load := new(opentsdb.Collect_load)
	load.Metric = "sys.load.5m"
	load.Value = js.AvgStat.Load5
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
//15分钟load
func load15(js base.SysLoadInfo,server string) *opentsdb.Collect_load{
	load := new(opentsdb.Collect_load)
	load.Metric = "sys.load.15m"
	load.Value = js.AvgStat.Load15
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
//procsRunning
func procsRunning(js base.SysLoadInfo,server string) *opentsdb.Collect_load{
	load := new(opentsdb.Collect_load)
	load.Metric = "sys.load.procsRunning"
	load.Value = float64(js.MiscStat.ProcsRunning)
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
//procsBlocked
func procsBlocked(js base.SysLoadInfo,server string) *opentsdb.Collect_load{
	load := new(opentsdb.Collect_load)
	load.Metric = "sys.load.procsBlocked"
	load.Value = float64(js.MiscStat.ProcsBlocked)
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
//ctxt
func ctxt(js base.SysLoadInfo,server string) *opentsdb.Collect_load{
	load := new(opentsdb.Collect_load)
	load.Metric = "sys.load.ctxt"
	load.Value = float64(js.MiscStat.Ctxt)
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}