package jvm
import (
	"fmt"
	"strconv"
	"monitor/base"
)
//检查端口是否开启
func checkPort(port string) int{
	url := "http://127.0.0.1:"+port+"/jolokia/"
	code,msg := base.HttpGet(url)
	if code == 0 {
		return 0
	}else{
		fmt.Println(string(msg))
		return 1
	}
}
//获取绑定jok的端口
func GetPort(jok *Jolokia,pid_slice []string) []string{
	var port_list []string
	portstart,_ := strconv.Atoi(jok.Portstart)
	for port := portstart;port<=portstart+len(pid_slice) ;port++  {
		strport := strconv.Itoa(port)
		if checkPort(strport) == 0{
			port_list = append(port_list, strport)
		}
	}
	return  port_list
}
