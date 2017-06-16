package network
import (
	"fmt"
	"net"
	"os"
)
//用udp发送数据
func UdpSend(server string,msg []byte){
	//var buf [512]byte
	udpAddr, err := net.ResolveUDPAddr("udp4", server)
	checkErr(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	defer conn.Close()
	checkErr(err)
	fmt.Println(server)
	fmt.Println(string(msg))
	conn.Write(msg)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}