package base

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"monitor/agent/conf"
)

//get方法
func HttpGet(url string) (n int,res []byte){
	response,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return 1,[]byte{}
	}else{
		defer response.Body.Close()
		body,_ := ioutil.ReadAll(response.Body)
		return 0,body
	}
}
//读取配置文件
func ReadConfig(mod,par string) string{
	config:= new(conf.Config)
	config.InitConfig("conf/agent.conf")
	return config.Read(mod,par)
}