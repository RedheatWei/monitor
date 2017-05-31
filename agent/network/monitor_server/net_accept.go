package main

import (
	"fmt"
	"net"
	"os"
	//"reflect"
	//"unsafe"
	//"strconv"
	//"strings"
	//"math"
	//"encoding/binary"
)

func main() {
	service := ":33990"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)
	for {
		handleClient(conn)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
//func BytesString(b []byte) String {
//	return *(*String)(unsafe.Pointer(&b))
//}
//func convert( b []byte ) string {
//	s := make([]string,len(b))
//	for i := range b {
//		s[i] = strconv.Itoa(int(b[i]))
//	}
//	return strings.Join(s,"")
//}

func handleClient(conn *net.UDPConn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, rAddr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			return
		}
		fmt.Println("client", rAddr.String(),string(buf[0:n]))

		_, err2 := conn.WriteToUDP([]byte("connect success"), rAddr)
		if err2 != nil {
			return
		}
	}
}