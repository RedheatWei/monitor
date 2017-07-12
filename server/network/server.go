package network

import (
	"net"
	"fmt"
	"monitor/server/base"
	"monitor/server/handler"
	db "monitor/server/db/mysqldb"
)
//读取配置文件
var ServerConfig base.ServerConfig
//读取ip地址
var AllowIplist []map[string]interface{}
func init()  {
	ServerConfig = base.GetConfig()
	AllowIplist = db.GetAllowIpList()
}
//开启udp服务端
func StartServer() {
	service := ":"+ServerConfig.Default.Port
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err!=nil{
		fmt.Println(err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err!=nil{
		fmt.Println(err)
	}
	for {
		handleClient(conn)
	}
}
//处理客户端的消息
func handleClient(conn *net.UDPConn) {
	var buf [4096]byte
	n, addr, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		fmt.Println(err)
	}
	add := addr.IP.String()
	chk,serverid,server := checkIp(add)
	fmt.Println(chk)
	if chk{
		fmt.Println(ServerConfig.OpentsDb.Enable)
		if ServerConfig.OpentsDb.Enable=="true"{
			go handler.ToTsJson(buf[:n],server)
		}else{
			go handler.ToJson(buf[:n],serverid)
		}
	}
}
//检查ip是否通行
func checkIp(ip string) (bool,int64,string){
	var is_in = bool(false)
	var serverid int64
	var server string
	for _,ipaddr := range AllowIplist{
		serverid = ipaddr["serverid"].(int64)
		server = ipaddr["server"].(string)
		if ipaddr["ip"] == ip {
			is_in = true
			return is_in,serverid,server
		}
	}
	serverid = 0
	return is_in,serverid,server
}