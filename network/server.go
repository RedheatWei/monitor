package network

import (
	"net"
	"monitor/base"
	"monitor/server/handler"
)

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
	_, _, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		return
	}
	//tmp := buf[:n]
	handler.ToJson(buf[:])
	//fmt.Println(addr)
	//fmt.Println(tmp)
	//fmt.Println(string(buf[:]))
}