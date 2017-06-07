package network

import (
	"fmt"
	"net"
	"bufio"
	"monitor/base"
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
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(string(message))
}