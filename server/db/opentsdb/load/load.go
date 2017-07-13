package load

import (
	"monitor/server/base"
	"fmt"
	"time"
	"reflect"
	"encoding/json"
	//"monitor/server/db/opentsdb"
)

type loadDB struct {
	AvgStatload1 *Collect_load
	AvgStatload5 *Collect_load
	AvgStatload15 *Collect_load
	MiscStatprocsRunning *Collect_load
	MiscStatprocsBlocked *Collect_load
	MiscStatctxt *Collect_load
}

func InsertLoadDB(js base.SysLoadInfo,server string){
	load := new(Collect_load)
	var loadDB loadDB
	loadDB.AvgStatload1 = load1(load,js,server)
	t := reflect.TypeOf(loadDB)
	v := reflect.ValueOf(loadDB)
	for k := 0; k < t.NumField(); k++{
		b,err := json.Marshal(v.Field(k).Interface())
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(string(b))
		//opentsdb.SendToTsDb(string(b))
	}
}
func load1(load *Collect_load,js base.SysLoadInfo,server string) *Collect_load{
	load.Metric = "sys.load.1m"
	load.Value = js.AvgStat.Load1
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	load.Tags.CtimeStamp = js.TimeStamp
	return load
}
