package opentsdb

import (
	"monitor/server/base"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)
func ConnDb() ([]byte,error){
	agentConfig := base.GetConfig()
	host := agentConfig.OpentsDb.Host
	url := host+"/api/stats"
	res,err := httpGet(url)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func SendToTsDb(data string){
	agentConfig := base.GetConfig()
	host := agentConfig.OpentsDb.Host
	url := host+"/api/stats"
	res,err := HttpPost(url,data)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(res))
}
func httpGet(url string) ([]byte,error){
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return body,nil
}

func HttpPost(url,data string) ([]byte,error){
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return body,nil
}