package network

import (
	"net"
	"monitor/base"
	"fmt"
)

var allow_iplist = []string{"192.168.1.238"}

func StartServer() {
	service := ":"+base.ReadServerConfig("default","port")
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [1024]byte
	n, _, err := conn.ReadFromUDP(buf[:])

	if err != nil {
		return
	}
	fmt.Println(conn.LocalAddr())
	//if checkIp(allow_iplist,conn.LocalAddr().Network()){
	//	go handler.ToJson(buf[:n],conn.LocalAddr())
	//}
}

func checkIp(allow_iplist []string,ip string) bool{
	var is_in = bool(false)
	for _,ipaddr := range allow_iplist{
		if ipaddr == ip {
			is_in = true
			return is_in
		}
	}
	return is_in
}