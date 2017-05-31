package jvm
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	//"net/url"
)
func getUrlRes(url string) (n int,msg []byte){
	//url := "http://127.0.0.1:"+port+"/jolokia/"
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
func checkPort(port string) int{
	url := "http://127.0.0.1:"+port+"/jolokia/"
	code,msg := getUrlRes(url)
	if code == 0 {
		return 0
	}else{
		fmt.Println(string(msg))
		return 1
	}
}
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