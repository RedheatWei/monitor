package handler

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"monitor/base"
	"fmt"
	dbjvm "monitor/server/db/jvm"
	dbcpu "monitor/server/db/cpu"
	dbdisk "monitor/server/db/disk"
	dbhost "monitor/server/db/host"
	dbload "monitor/server/db/load"
	dbmem "monitor/server/db/mem"
	dbnet "monitor/server/db/net"
)
func  ToJson(rec []byte,addr string){
	fmt.Println(string(rec))
	js, _ := simplejson.NewJson(rec)
	js_map,_ := js.Map()
	switch js_map["Type"] {
	case "JvmInfo":
		var info base.JvmInfo
		json.Unmarshal(rec,&info)
		go dbjvm.InsertJvmDB(info,addr)
	case "SysMemInfo":
		var info base.SysMemInfo
		json.Unmarshal(rec,&info)
		go dbmem.InsertMemDB(info,addr)
	case "SysCpuInfo":
		var info base.SysCpuInfo
		json.Unmarshal(rec,&info)
		go dbcpu.InsertCpuDB(info,addr)
	case "SysDiskInfo":
		var info base.SysDiskInfo
		json.Unmarshal(rec,&info)
		go dbdisk.InsertDiskDB(info,addr)
	case "SysHostInfo":
		var info base.SysHostInfo
		json.Unmarshal(rec,&info)
		go dbhost.InsertHostDB(info,addr)
	case "SysLoadInfo":
		var info base.SysLoadInfo
		json.Unmarshal(rec,&info)
		go dbload.InsertLoadDB(info,addr)
	case "SysNetInfo":
		var info base.SysNetInfo
		json.Unmarshal(rec,&info)
		go dbnet.InsertNetDB(info,addr)

	}
}