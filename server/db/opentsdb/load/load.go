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
	AvgStatload1 *Collect_load
	AvgStatload5 *Collect_load
	AvgStatload15 *Collect_load
	MiscStatprocsRunning *Collect_load
	MiscStatprocsBlocked *Collect_load
	MiscStatctxt *Collect_load
}
//主函数
func InsertLoadDB(js base.SysLoadInfo,server string){
	load := new(Collect_load)
	var loadDB loadDB
	loadDB.AvgStatload1 = load1(load,js,server)
	loadDB.AvgStatload1 = load5(load,js,server)
	loadDB.AvgStatload1 = load15(load,js,server)
	loadDB.AvgStatload1 = procsRunning(load,js,server)
	loadDB.AvgStatload1 = procsBlocked(load,js,server)
	loadDB.AvgStatload1 = ctxt(load,js,server)
	t := reflect.TypeOf(loadDB)
	v := reflect.ValueOf(loadDB)
	for k := 0; k < t.NumField(); k++{
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
func load1(load *Collect_load,js base.SysLoadInfo,server string) *Collect_load{
	load.Metric = "sys.load.1m"
	load.Value = js.AvgStat.Load1
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
//5分钟load
func load5(load *Collect_load,js base.SysLoadInfo,server string) *Collect_load{
	load.Metric = "sys.load.5m"
	load.Value = js.AvgStat.Load5
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
//15分钟load
func load15(load *Collect_load,js base.SysLoadInfo,server string) *Collect_load{
	load.Metric = "sys.load.15m"
	load.Value = js.AvgStat.Load15
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
//procsRunning
func procsRunning(load *Collect_load,js base.SysLoadInfo,server string) *Collect_load{
	load.Metric = "sys.load.procsRunning"
	load.Value = float64(js.MiscStat.ProcsRunning)
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
//procsBlocked
func procsBlocked(load *Collect_load,js base.SysLoadInfo,server string) *Collect_load{
	load.Metric = "sys.load.procsBlocked"
	load.Value = float64(js.MiscStat.ProcsBlocked)
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
//ctxt
func ctxt(load *Collect_load,js base.SysLoadInfo,server string) *Collect_load{
	load.Metric = "sys.load.ctxt"
	load.Value = float64(js.MiscStat.Ctxt)
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}