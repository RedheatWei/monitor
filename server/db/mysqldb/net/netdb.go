package net

import (
	"monitor/server/base"
	"fmt"
	db "monitor/server/db/mysqldb"

	"github.com/shirou/gopsutil/net"

)

func InsertNetDB(js base.SysNetInfo,serverid int64){
	db := db.ConnDB()
	net := new(Collect_net)
	net.ServerId = serverid
	net.TimeStamp = js.TimeStamp
	affected, err := db.Insert(net)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
	stat_len := len(js.IOCountersStat)
	if stat_len > 0{
		insertIOStatDB(js.IOCountersStat,serverid,js.TimeStamp,stat_len)
	}
	stat_len2 := len(js.InterfaceStat)
	if stat_len2 > 0{
		insertInterfaceDB(js.InterfaceStat,serverid,js.TimeStamp,stat_len2)
	}
}
func insertIOStatDB(io []net.IOCountersStat,serverid int64,timestamp int64,stat_len int){
	db := db.ConnDB()
	userstats := make([]*Collect_net_iocountersstat,stat_len)
	for key,userstat_s := range io{
		userstats[key] = new(Collect_net_iocountersstat)
		userstats[key].ServerId = serverid
		userstats[key].IOCountersStatname = userstat_s.Name
		userstats[key].IOCountersStatbytesSent = userstat_s.BytesSent
		userstats[key].IOCountersStatbytesRecv = userstat_s.BytesRecv
		userstats[key].IOCountersStatpacketsSent = userstat_s.PacketsSent
		userstats[key].IOCountersStatpacketsRecv = userstat_s.PacketsRecv
		userstats[key].IOCountersStaterrin = userstat_s.Errin
		userstats[key].IOCountersStaterrout = userstat_s.Errout
		userstats[key].IOCountersStatdropin = userstat_s.Dropin
		userstats[key].IOCountersStatdropout = userstat_s.Dropout
		userstats[key].IOCountersStatfifoin = userstat_s.Fifoin
		userstats[key].IOCountersStatfifoout = userstat_s.Fifoout
		userstats[key].TimeStamp = timestamp
	}
	affected, err := db.Insert(userstats)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
}
func insertInterfaceDB(ifc []net.InterfaceStat,serverid int64,timestamp int64,stat_len int){
	db := db.ConnDB()
	userstats := make([]*Collect_net_interfacestat,stat_len)
	for key,userstat_s := range ifc{
		userstats[key] = new(Collect_net_interfacestat)
		userstats[key].ServerId = serverid
		userstats[key].InterfaceStatmtu = userstat_s.MTU
		userstats[key].InterfaceStatname = userstat_s.Name
		userstats[key].InterfaceStathardwareaddr = userstat_s.HardwareAddr
		userstats[key].TimeStamp = timestamp
		stat_len2 := len(userstat_s.Addrs)
		if stat_len2 > 0{
			insertInterfaceAddrDB(userstat_s.Addrs,serverid,timestamp,stat_len2)
		}
	}
	affected, err := db.Insert(userstats)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
}
func insertInterfaceAddrDB(ipaddr []net.InterfaceAddr,serverid int64,timestamp int64,stat_len int){
	db := db.ConnDB()
	userstats := make([]*Collect_net_interfacestat_addrs,stat_len)
	for key,userstat_s := range ipaddr{
		userstats[key] = new(Collect_net_interfacestat_addrs)
		userstats[key].ServerId = serverid
		userstats[key].Addrsaddr = userstat_s.Addr
		userstats[key].TimeStamp = timestamp
	}
	affected, err := db.Insert(userstats)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(affected)
}