package network

import (
	"fmt"
	"net"
)

func StartServer() {
	service := ":33990"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)
	for {
		handleClient(conn)
	}
}


func handleClient(conn *net.UDPConn) {
	//defer conn.Close()
	buf := make([]byte,512)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		return
	}
	fmt.Println(string(buf[0:n]))

	//_, err2 := conn.WriteToUDP([]byte("Received"), rAddr)
	//if err2 != nil {
	//	return
	//}
}