package network

import (
	"fmt"
	"net"
	"syscall"
)

func StartServer() {
	service := ":33990"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)
	fmt.Println(conn)
	for {
		handleClient(conn)
	}
}


func handleClient(conn *net.UDPConn) {
	//defer conn.Close()

	oob := make([]byte, 512)
	buff := make([]byte, 512)
	_, _, flags, _, _ := conn.ReadMsgUDP(buff, oob)
	if flags & syscall.MSG_TRUNC != 0 {
		fmt.Println("truncated read")
	}
	fmt.Println(string(buff),string(oob))

	//buf := make([]byte,512)
	//n, _, err := conn.ReadFromUDP(buf)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(buf[0:n]))

	//_, err2 := conn.WriteToUDP([]byte("Received"), rAddr)
	//if err2 != nil {
	//	return
	//}
}