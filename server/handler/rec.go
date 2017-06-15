package handler

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"monitor/server/db"
	"monitor/base"
)
func ToJson(rec []byte,addr string){
	//json.Unmarshal(rec,&js)
	js, _ := simplejson.NewJson(rec)
	js_map,_ := js.Map()
	switch js_map["Type"] {
	case "JvmInfo":
		var info base.JvmInfo
		json.Unmarshal(rec,&info)
		go db.InsertJvmDB(info,addr)
	case "jvminfo":
		var Data base.JvmInfo
		json.Unmarshal(rec,&Data)
		go db.InsertJvmDB(Data,addr)
	case "SysMemInfo":
		var info base.SysMemInfo
		json.Unmarshal(rec,&info)
		go db.InsertMemDB(info,addr)
	case "SysCpuInfo":
		var info base.SysCpuInfo
		json.Unmarshal(rec,&info)
		go db.InsertCpuDB(info,addr)
	case "SysDiskInfo":
		var info base.SysDiskInfo
		json.Unmarshal(rec,&info)
		go db.InsertDiskDB(info,addr)
	case "SysHostInfo":
		var info base.SysHostInfo
		json.Unmarshal(rec,&info)
		go db.InsertHostDB(info,addr)
	case "SysLoadInfo":
		var info base.SysLoadInfo
		json.Unmarshal(rec,&info)
		go db.InsertLoadDB(info,addr)
	case "SysNetInfo":
		var info base.SysNetInfo
		json.Unmarshal(rec,&info)
		go db.InsertNetDB(info,addr)

	}
}