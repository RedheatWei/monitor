package network

import (
	"net"
	"monitor/base"
	"monitor/server/handler"
	"monitor/server/db"
	"fmt"
)

var ServerConfig base.ServerConfig
var AllowIplist []map[string]interface{}
func init()  {
	ServerConfig = base.GetConfig()
	AllowIplist = db.GetAllowIpList()
	for j,k := range AllowIplist{
		//fmt.Println(j)
		fmt.Println(k["ip"])
	}
}
func StartServer() {
	service := ":"+ServerConfig.Default.Port
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)
	for {
		handleClient(conn)
	}
}



func handleClient(conn *net.UDPConn) {
	var buf [2048]byte
	n, addr, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		return
	}
	add := addr.IP.String()

	//if checkIp(add){
	go handler.ToJson(buf[:n],add)
	//}
}

//func checkIp(ip string) bool{
//	var is_in = bool(false)
//	for _,ipaddr := range AllowIplist{
//		if ipaddr == ip {
//			is_in = true
//			return is_in
//		}
//	}
//	return is_in
//}