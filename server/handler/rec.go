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
	//fmt.Println()
	fmt.Println(string(js.SysMemData))
	go db.InsertDB(js.JvmData,addr)
}