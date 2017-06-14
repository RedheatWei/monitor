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
	fmt.Println(string(js))
	switch js.Type {
	case "JvmInfo":
		go db.InsertJvmDB(js.JvmInfo,addr)
	case "jvminfo":
		go db.InsertJvmDB(js.JvmInfo,addr)
	}
}