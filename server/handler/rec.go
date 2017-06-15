package handler

import (
	//"encoding/json"
	//"monitor/base"
	//"monitor/server/db"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func ToJson(rec []byte,addr string){
	//var js interface{}
	//json.Unmarshal(rec,&js)
	js, _ := simplejson.NewJson(rec)
	fmt.Println(js)
	//switch js.Type {
	//case "JvmInfo":
	//	go db.InsertJvmDB(js.Data,addr)
	//case "jvminfo":
	//	go db.InsertJvmDB(js.JvmInfo,addr)
	//}
}