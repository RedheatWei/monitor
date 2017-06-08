package handler

import (
	"encoding/json"
	"monitor/server/db"
	"net"
	"monitor/base"
)

func ToJson(rec []byte,addr *net.UDPAddr){
	var js base.JsonInfo
	json.Unmarshal(rec,&js)
	go db.InsertDB(js,addr)
}