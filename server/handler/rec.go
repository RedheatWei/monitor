package handler

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"monitor/server/base"
	"fmt"
	dbjvm "monitor/server/db/mysqldb/jvm"
	dbcpu "monitor/server/db/mysqldb/cpu"
	dbdisk "monitor/server/db/mysqldb/disk"
	dbhost "monitor/server/db/mysqldb/host"
	dbload "monitor/server/db/mysqldb/load"
	dbmem "monitor/server/db/mysqldb/mem"
	dbnet "monitor/server/db/mysqldb/net"
	tsload "monitor/server/db/opentsdb/load"
)
//转换json并插入数据库
func  ToJson(rec []byte,serverid int64){
	fmt.Println(string(rec))
	js, _ := simplejson.NewJson(rec)
	js_map,_ := js.Map()
	switch js_map["Type"] {
	case "JvmInfo":
		var info base.JvmInfo
		json.Unmarshal(rec,&info)
		go dbjvm.InsertJvmDB(info,serverid)
	case "SysMemInfo":
		var info base.SysMemInfo
		json.Unmarshal(rec,&info)
		go dbmem.InsertMemDB(info,serverid)
	case "SysCpuInfo":
		var info base.SysCpuInfo
		json.Unmarshal(rec,&info)
		go dbcpu.InsertCpuDB(info,serverid)
	case "SysDiskInfo":
		var info base.SysDiskInfo
		json.Unmarshal(rec,&info)
		go dbdisk.InsertDiskDB(info,serverid)
	case "SysHostInfo":
		var info base.SysHostInfo
		json.Unmarshal(rec,&info)
		go dbhost.InsertHostDB(info,serverid)
	case "SysLoadInfo":
		var info base.SysLoadInfo
		json.Unmarshal(rec,&info)
		go dbload.InsertLoadDB(info,serverid)
	case "SysNetInfo":
		var info base.SysNetInfo
		json.Unmarshal(rec,&info)
		go dbnet.InsertNetDB(info,serverid)
	}
}
func ToTsJson(rec []byte,server string){
	//fmt.Println(string(rec))
	js, _ := simplejson.NewJson(rec)
	js_map,_ := js.Map()
	switch js_map["Type"] {
	//case "JvmInfo":
	//	var info base.JvmInfo
	//	json.Unmarshal(rec,&info)
	//	go dbjvm.InsertJvmDB(info,server)
	//case "SysMemInfo":
	//	var info base.SysMemInfo
	//	json.Unmarshal(rec,&info)
	//	go dbmem.InsertMemDB(info,server)
	//case "SysCpuInfo":
	//	var info base.SysCpuInfo
	//	json.Unmarshal(rec,&info)
	//	go dbcpu.InsertCpuDB(info,server)
	//case "SysDiskInfo":
	//	var info base.SysDiskInfo
	//	json.Unmarshal(rec,&info)
	//	go dbdisk.InsertDiskDB(info,server)
	//case "SysHostInfo":
	//	var info base.SysHostInfo
	//	json.Unmarshal(rec,&info)
	//	go dbhost.InsertHostDB(info,server)
	case "SysLoadInfo":
		var info base.SysLoadInfo
		go tsload.InsertLoadDB(info,server)
	//case "SysNetInfo":
	//	var info base.SysNetInfo
	//	json.Unmarshal(rec,&info)
	//	go dbnet.InsertNetDB(info,server)
	}
}