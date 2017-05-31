package jvm
import (
	"fmt"
	"net/http"
	"io/ioutil"
)
func getJokPort(jok *Jolokia) string{
	return jok.Portstart
}
func GetUrlRes(url string){
	//url := "http://127.0.0.1:"+port+"/jolokia/"
	response,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}else{
		defer response.Body.Close()
		body,_ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
		//return body
	}

}