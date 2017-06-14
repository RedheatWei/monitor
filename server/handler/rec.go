package handler

import (
	"encoding/json"
	"monitor/base"
	"monitor/server/db"
)

func ToJson(rec []byte,addr string){
	var js base.Senddata
	json.Unmarshal(rec,&js)
	switch js.Type {
	case "JvmInfo":
		go db.InsertJvmDB(js.JvmInfo,addr)
	case "jvminfo":
		go db.InsertJvmDB(js.JvmInfo,addr)
	}
}