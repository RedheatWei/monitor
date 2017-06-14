package handler

import (
	"encoding/json"
	"monitor/base"
	"monitor/server/db"
	"fmt"
)

func ToJson(rec []byte,addr string){
	var js base.Senddata
	json.Unmarshal(rec,&js)
	fmt.Println(js.Type)

	switch js.Type {
	case "JvmInfo":
		go db.InsertJvmDB(js.JvmData,addr)
	//case "SysMemInfo":
	//	go db.InsertMemDB(js.SysMemData,addr)
	}
}