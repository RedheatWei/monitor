package net

import (
	"monitor/server/db/opentsdb"
	"time"
	"reflect"
	"encoding/json"
	"fmt"
	"monitor/server/base"
	"github.com/shirou/gopsutil/net"
)
//定义收集的项目
type netDB struct {
	IOCountersStatbytesSent *opentsdb.Collect_net
	IOCountersStatbytesRecv *opentsdb.Collect_net
	IOCountersStatpacketsSent *opentsdb.Collect_net
	IOCountersStatpacketsRecv *opentsdb.Collect_net
}
//主函数
func InsertNetDB(js base.SysNetInfo,serverIpInfo base.ServerIpInfo){
	var netDB netDB
	var sli_str []interface{}
	for _,usage_stat := range js.IOCountersStat{
		netDB.IOCountersStatbytesSent = iOCountersStatbytesSent(js,serverIpInfo,usage_stat)
		netDB.IOCountersStatbytesRecv = iOCountersStatbytesRecv(js,serverIpInfo,usage_stat)
		netDB.IOCountersStatpacketsSent = iOCountersStatpacketsSent(js,serverIpInfo,usage_stat)
		netDB.IOCountersStatpacketsRecv = iOCountersStatpacketsRecv(js,serverIpInfo,usage_stat)
		v := reflect.ValueOf(netDB)
		for k := 0; k < v.NumField(); k++{
			if ! v.Field(k).IsNil(){
				val := v.Field(k).Interface()
				sli_str = append(sli_str,val)
			}
		}
	}
	b,err := json.Marshal(sli_str)
	if err!=nil{
		fmt.Println(err)
	}
	opentsdb.SendToTsDb(string(b))
}
//组合数据,便于修改tag
func createCollect(js base.SysNetInfo,serverIpInfo base.ServerIpInfo,usage_stat net.IOCountersStat) *opentsdb.Collect_net{
	collect := new(opentsdb.Collect_net)
	collect.TimeStamp = time.Now().Unix()
	collect.Tags.Server = serverIpInfo.Server
	collect.Tags.Ip = serverIpInfo.Ip
	collect.Tags.CtimeStamp = js.TimeStamp
	collect.Tags.IOCountersStatname = usage_stat.Name
	return collect
}

func iOCountersStatbytesSent(js base.SysNetInfo,serverIpInfo base.ServerIpInfo,usage_stat net.IOCountersStat) *opentsdb.Collect_net{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.net.iOCountersStatbytesSent"
	collect.Value = float64(usage_stat.BytesSent)
	return collect
}
func iOCountersStatbytesRecv(js base.SysNetInfo,serverIpInfo base.ServerIpInfo,usage_stat net.IOCountersStat) *opentsdb.Collect_net{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.net.iOCountersStatbytesRecv"
	collect.Value = float64(usage_stat.BytesRecv)
	return collect
}
func iOCountersStatpacketsSent(js base.SysNetInfo,serverIpInfo base.ServerIpInfo,usage_stat net.IOCountersStat) *opentsdb.Collect_net{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.net.iOCountersStatpacketsSent"
	collect.Value = float64(usage_stat.PacketsSent)
	return collect
}
func iOCountersStatpacketsRecv(js base.SysNetInfo,serverIpInfo base.ServerIpInfo,usage_stat net.IOCountersStat) *opentsdb.Collect_net{
	collect := createCollect(js,serverIpInfo,usage_stat)
	collect.Metric = "sys.net.iOCountersStatpacketsRecv"
	collect.Value = float64(usage_stat.PacketsRecv)
	return collect
}