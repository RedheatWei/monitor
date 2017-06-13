package handler

import (
	"encoding/json"
	"monitor/base"
	"monitor/server/db"
)

func ToJson(rec []byte,addr string){
	var js base.Senddata
	json.Unmarshal(rec,&js)
	//fmt.Println()
	go db.InsertDB(js.JvmData,addr)
}