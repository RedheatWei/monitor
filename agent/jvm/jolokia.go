package jvm
import(
	"fmt"
	"os/exec"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

type Jolokia struct {
	Jolokiapath,Jolokianame,Portstart string
}
//获取Java程序pid
func GetPid(jok *Jolokia) []string{
	jolokia := "java -jar "+jok.Jolokiapath+jok.Jolokianame+" list | grep -v 'jolokia' | cut -d' ' -f1"
	//cmd := exec.Command("/bin/sh", "-c", `java -jar /usr/local/jolokia/jolokia-jvm-1.3.6-agent.jar list | grep -v "jolokia" | cut -d' ' -f1`)
	opBytes := execShell(jolokia)
	pid_slice := strings.Split(string(opBytes),"\n")
	return pid_slice
}
//开启jolokia
func StartJok(jok *Jolokia,pid_slice []string){
	for _,pid:=range pid_slice[:len(pid_slice)-1]{
		if(bingPort(jok,pid,len(pid_slice))==0){
			continue
		}else{
			break
		}
	}
}
//绑定端口
func bingPort(jok *Jolokia,pid string,pid_num int) int{
	portstart,_ := strconv.Atoi(jok.Portstart)
	for port := portstart; port < portstart+pid_num; port++ {
		fmt.Println(pid)
		jolokia := "java -jar "+jok.Jolokiapath+jok.Jolokianame+" --host 127.0.0.1 --port="+strconv.Itoa(port)+" start "+pid
		fmt.Println(jolokia)
		opBytes := execShell(jolokia)
		fmt.Println(string(opBytes))
		if strings.Contains(string(opBytes),"127.0.0.1"){
			return 0
		}else{
			continue
		}
	}
	return 1
}

//停止监听
func StopJok(jok *Jolokia,pid_slice []string){
	for _,pid:=range pid_slice[:len(pid_slice)-1]{
		jolokia := "java -jar "+jok.Jolokiapath+jok.Jolokianame+" stop "+pid
		opBytes := execShell(jolokia)
		fmt.Println(string(opBytes))
	}
}
//执行shell
func execShell(shell string) []byte{
	cmd := exec.Command("/bin/sh", "-c", shell)
	stdout, err := cmd.StdoutPipe()
	checkErr(err)
	defer stdout.Close()
	checkErr(cmd.Start())
	opBytes, err := ioutil.ReadAll(stdout)
	checkErr(err)
	return opBytes
}

//func main(){
//	jok := &Jolokia{"/usr/local/jolokia/","jolokia-jvm-1.3.6-agent.jar"}
//	pid_slice := GetPid(jok)
//	StartJok(jok,pid_slice)
//}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}