package handler

import (
	"encoding/json"
	//"monitor/server/db"
	"monitor/base"
	"fmt"
)

func ToJson(rec []byte,addr string){
	var js base.Senddata
	json.Unmarshal(rec,&js)
	fmt.Println(js)
	//go db.InsertDB(js,addr)
}