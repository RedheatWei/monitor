package sock
import (
	"io"
	"net"
	"monitor/agent/base"
)

func reader(r io.Reader) (n int,res []byte){
	buf := make([]byte, 1024)
	n, err := r.Read(buf[:])
	if err != nil {
		return
	} else {
		return 0,buf[0:n]
	}
}

func SendMsg(msg string) {
	c, err := net.Dial("unix",base.ReadConfig("jvm","sock"))
	if err != nil {
		panic(err.Error())
	}
	//defer c.Close()
	//go reader(c)
	c.Write([]byte(msg))
}
