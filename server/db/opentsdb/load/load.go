package load

import (
	"monitor/server/base"
	"fmt"
	"time"
	"encoding/json"
	"monitor/server/db/opentsdb"
)

func InsertLoadDB(js base.SysLoadInfo,server string){
	_,err := opentsdb.ConnDb()
	if err != nil{
		fmt.Println(err)
	}
	load := new(Collect_load)
	load.Metric = "sys.load.1m"
	load.Value = js.AvgStat.Load1
	load.TimeStamp = time.Now().Unix()
	load.Tags.Server = server
	b,err := json.Marshal(load)
	if err!=nil{
		fmt.Println(err)
	}
	//fmt.Println()
	//var load_js Collect_load
	//json.Unmarshal(b,&load_js)
	fmt.Println(string(b))
	opentsdb.SendToTsDb(string(b))
}