package handler

import (
	"encoding/json"
	//"monitor/base"
	//"monitor/server/db"
	"fmt"
	"github.com/bitly/go-simplejson"
	"monitor/server/db"
	"monitor/base"
)

func ToJson(rec []byte,addr string){
	//json.Unmarshal(rec,&js)
	js, _ := simplejson.NewJson(rec)
	fmt.Println(js)
	js_map,_ := js.Map()
	switch js_map["Type"] {
	case "JvmInfo":
		var Data base.JvmInfo
		json.Unmarshal(rec,&Data)
		fmt.Println(Data)
		go db.InsertJvmDB(Data,addr)
	case "jvminfo":
		var Data base.JvmInfo
		json.Unmarshal(rec,&Data)
		fmt.Println(Data)
		go db.InsertJvmDB(Data,addr)
	}
}