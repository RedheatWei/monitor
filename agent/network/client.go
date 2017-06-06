package network
import (
	"fmt"
	"net"
	"os"
	"reflect"
	"unsafe"
)
func UdpSend(server string,msg []byte){
	//var buf [512]byte
	fmt.Printf("send to %s \n", server)
	udpAddr, err := net.ResolveUDPAddr("udp4", server)
	checkErr(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	defer conn.Close()
	checkErr(err)
	//rAddr := conn.RemoteAddr()
	//total := Int64ToBytes(msg.Total)
	//total := strconv.FormatUint(msg.Total,10)
	//n, err := conn.Write(S2B(&total))
	conn.Write(msg)
	//checkErr(err)
	//n, err = conn.Read(buf[0:])
	//checkErr(err)
	//fmt.Println("Reply:", rAddr.String(), string(buf[0:n]))
}
func S2B(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(s))))
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
//func Int64ToBytes(i uint64) []byte {
//	var buf = make([]byte, 8)
//	binary.BigEndian.PutUint64(buf, i)
//	return buf
//}
//func main() {
//	//v, _ := mem.VirtualMemory()
//	//fmt.Println("type:", reflect.TypeOf(Int64ToBytes(v.Total)))
//	//fmt.Println(v.Total)
//	//udpSend("123.56.92.243","33990",v)
//	udpSend("127.0.0.1:33990",v)
//}