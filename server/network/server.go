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
	chk,serverIpInfo := checkIp(add)
	if chk{
		if ServerConfig.OpentsDb.Enable=="true"{
			go handler.ToTsJson(buf[:n],serverIpInfo)
		}else{
			go handler.ToJson(buf[:n],serverIpInfo)
		}
	}
}
//检查ip是否通行
func checkIp(ip string) (bool,base.ServerIpInfo){
	var is_in = bool(false)
	var serverIpInfo base.ServerIpInfo
	for _,ipaddr := range AllowIplist{
		if ipaddr["ip"] == ip {
			is_in = true
			serverIpInfo.ServerId = ipaddr["serverid"].(int64)
			serverIpInfo.Server = ipaddr["server"].(string)
			serverIpInfo.Ip = ipaddr["ip"].(string)
			serverIpInfo.Type = ipaddr["type"].(string)
			return is_in,serverIpInfo
		}
	}
	return is_in,serverIpInfo
}