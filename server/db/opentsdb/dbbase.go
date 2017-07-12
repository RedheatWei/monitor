package opentsdb

import (
	"monitor/server/network"
	"monitor/agent/base"
	"fmt"
)
func ConnDb() ([]byte,error){
	agentConfig := base.GetConfig()
	host := agentConfig.OpentsDb.Host
	url := host+"/api/stats"
	res,err := network.HttpGet(url)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func SendToTsDb(data string){
	agentConfig := base.GetConfig()
	host := agentConfig.OpentsDb.Host
	url := host+"/api/stats"
	res,err := network.HttpPost(url,data)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(res))
}