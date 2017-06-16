package network

import (
	"net"
	"fmt"
	"monitor/server/base"
	"monitor/server/handler"
	"monitor/server/db"
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
	var buf [2048]byte
	n, addr, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		return
	}
	add := addr.IP.String()
	fmt.Println(n)
	fmt.Println(add)
	if checkIp(add){
		go handler.ToJson(buf[:n],add)
	}
}
//检查ip是否通行
func checkIp(ip string) bool{
	var is_in = bool(false)
	for _,ipaddr := range AllowIplist{
		if ipaddr["ip"] == ip {
			is_in = true
			return is_in
		}
	}
	return is_in
}