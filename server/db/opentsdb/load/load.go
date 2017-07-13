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
	AvgStatload1 *opentsdb.Collect
	AvgStatload5 *opentsdb.Collect
	AvgStatload15 *opentsdb.Collect
	MiscStatprocsRunning *opentsdb.Collect
	MiscStatprocsBlocked *opentsdb.Collect
	MiscStatctxt *opentsdb.Collect
}
//主函数
func InsertLoadDB(js base.SysLoadInfo,server string){
	var loadDB loadDB
	loadDB.AvgStatload1 = load1(js,server)
	loadDB.AvgStatload5 = load5(js,server)
	//loadDB.AvgStatload15 = load15(js,server)
	//loadDB.MiscStatprocsRunning = procsRunning(js,server)
	//loadDB.MiscStatprocsBlocked = procsBlocked(js,server)
	loadDB.MiscStatctxt = ctxt(js,server)
	v := reflect.ValueOf(loadDB)
	var sli_str []interface{}
	var str string
	for k := 0; k < v.NumField(); k++{
		val := v.Field(k).Interface()
		fmt.Println(v.Field(k))
		sli_str = append(sli_str,val)
		b,err := json.Marshal(sli_str)
		if err!=nil{
			fmt.Println(err)
		}
		str = string(b)
		//if str != "null"{
		//}
	}
	//fmt.Println(str)
	opentsdb.SendToTsDb(string(str))
}
//1分钟load
func load1(js base.SysLoadInfo,server string) *opentsdb.Collect{
	collect := new(opentsdb.Collect)
	collect.Metric = "sys.load.1m"
	collect.Value = js.AvgStat.Load1
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = server
	collect.Tags.CtimeStamp = js.TimeStamp
	return collect
}
//5分钟load
func load5(js base.SysLoadInfo,server string) *opentsdb.Collect{
	collect := new(opentsdb.Collect)
	collect.Metric = "sys.load.5m"
	collect.Value = js.AvgStat.Load5
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = server
	collect.Tags.CtimeStamp = js.TimeStamp
	return collect
}
//15分钟load
func load15(js base.SysLoadInfo,server string) *opentsdb.Collect{
	collect := new(opentsdb.Collect)
	collect.Metric = "sys.load.15m"
	collect.Value = js.AvgStat.Load15
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = server
	collect.Tags.CtimeStamp = js.TimeStamp
	return collect
}
//procsRunning
func procsRunning(js base.SysLoadInfo,server string) *opentsdb.Collect{
	collect := new(opentsdb.Collect)
	collect.Metric = "sys.load.procsRunning"
	collect.Value = float64(js.MiscStat.ProcsRunning)
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = server
	collect.Tags.CtimeStamp = js.TimeStamp
	return collect
}
//procsBlocked
func procsBlocked(js base.SysLoadInfo,server string) *opentsdb.Collect{
	collect := new(opentsdb.Collect)
	collect.Metric = "sys.load.procsBlocked"
	collect.Value = float64(js.MiscStat.ProcsBlocked)
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = server
	collect.Tags.CtimeStamp = js.TimeStamp
	return collect
}
//ctxt
func ctxt(js base.SysLoadInfo,server string) *opentsdb.Collect{
	collect := new(opentsdb.Collect)
	collect.Metric = "sys.load.ctxt"
	collect.Value = float64(js.MiscStat.Ctxt)
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = server
	collect.Tags.CtimeStamp = js.TimeStamp
	return collect
}