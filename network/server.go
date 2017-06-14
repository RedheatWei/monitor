package network

import (
	"net"
	"monitor/base"
	"monitor/server/handler"
	"strings"
)

var allow_iplist = strings.Split(base.ReadServerConfig("default","allow_iplist"),",")

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
	var buf [2048]byte
	n, addr, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		return
	}
	add := addr.IP.String()
	if checkIp(allow_iplist,add){
		go handler.ToJson(buf[:n],add)
	}
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