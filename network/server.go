package network

import (
	"fmt"
	"net"
	"bufio"
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
	//buf := make([]byte,512)
	//n, _, err := conn.ReadFromUDP(buf)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(string(message))

	//_, err2 := conn.WriteToUDP([]byte("Received"), rAddr)
	//if err2 != nil {
	//	return
	//}
}