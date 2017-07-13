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
func InsertLoadDB(js base.SysLoadInfo,serverIpInfo base.ServerIpInfo){
	var loadDB loadDB
	loadDB.AvgStatload1 = avgStatload1(js,serverIpInfo)
	loadDB.AvgStatload5 = avgStatload5(js,serverIpInfo)
	loadDB.AvgStatload15 = avgStatload15(js,serverIpInfo)
	loadDB.MiscStatprocsRunning = miscStatprocsRunning(js,serverIpInfo)
	loadDB.MiscStatprocsBlocked = miscStatprocsBlocked(js,serverIpInfo)
	loadDB.MiscStatctxt = miscStatctxt(js,serverIpInfo)
	v := reflect.ValueOf(loadDB)
	var sli_str []interface{}
	for k := 0; k < v.NumField(); k++{
		if ! v.Field(k).IsNil(){
			val := v.Field(k).Interface()
			sli_str = append(sli_str,val)
		}
	}
	b,err := json.Marshal(sli_str)
	if err!=nil{
		fmt.Println(err)
	}
	opentsdb.SendToTsDb(string(b))
}
//组合数据,便于修改tag
func createCollect(js base.SysLoadInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect{
	collect := new(opentsdb.Collect)
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = serverIpInfo.Server
	collect.Tags.Ip = serverIpInfo.Ip
	collect.Tags.CtimeStamp = js.TimeStamp
	return collect
}
//1分钟load
func avgStatload1(js base.SysLoadInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.load.avgStatload1"
	collect.Value = js.AvgStat.Load1
	return collect
	return collect
}
//5分钟load
func avgStatload5(js base.SysLoadInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.load.avgStatload5"
	collect.Value = js.AvgStat.Load5
	return collect
}
//15分钟load
func avgStatload15(js base.SysLoadInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.load.avgStatload15"
	collect.Value = js.AvgStat.Load15
	return collect
}
//procsRunning
func miscStatprocsRunning(js base.SysLoadInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.load.miscStatprocsRunning"
	collect.Value = float64(js.MiscStat.ProcsRunning)
	return collect
}
//procsBlocked
func miscStatprocsBlocked(js base.SysLoadInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.load.miscStatprocsBlocked"
	collect.Value = float64(js.MiscStat.ProcsBlocked)
	return collect
}
//ctxt
func miscStatctxt(js base.SysLoadInfo,serverIpInfo base.ServerIpInfo) *opentsdb.Collect{
	collect := createCollect(js,serverIpInfo)
	collect.Metric = "sys.load.miscStatctxt"
	collect.Value = float64(js.MiscStat.Ctxt)
	return collect
}