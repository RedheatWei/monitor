package main
import (
	"fmt"
	"net"
	"os"
	"github.com/shirou/gopsutil/mem"
	"reflect"
	"encoding/binary"
	"strconv"
	"unsafe"
)
var(
	server,input string
)
func udpSend(serverip ,port string,msg *mem.VirtualMemoryStat){
	var buf [512]byte
	server += serverip+":"+port
	fmt.Printf("send to %s \n", server)
	udpAddr, err := net.ResolveUDPAddr("udp4", server)
	checkErr(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	defer conn.Close()
	checkErr(err)
	rAddr := conn.RemoteAddr()

	//total := Int64ToBytes(msg.Total)
	total := strconv.FormatUint(msg.Total,10)
	n, err := conn.Write(S2B(&total))

	checkErr(err)
	n, err = conn.Read(buf[0:])
	checkErr(err)
	fmt.Println("Reply:", rAddr.String(), string(buf[0:n]))
	os.Exit(0)
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
func Int64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}
func main() {
	v, _ := mem.VirtualMemory()
	fmt.Println("type:", reflect.TypeOf(Int64ToBytes(v.Total)))
	// convert to JSON. String() is also implemented
	fmt.Println(v.Total)
	//udpSend("123.56.92.243","33990",v)
	udpSend("127.0.0.1","33990",v)
}