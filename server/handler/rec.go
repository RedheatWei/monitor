package handler

import (
	"encoding/json"
	"monitor/server/db"
	"monitor/base"
)

func ToJson(rec []byte,addr string){
	var js base.JsonInfo
	json.Unmarshal(rec,&js)
	go db.InsertDB(js,addr)
}