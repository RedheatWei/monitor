package sock

import (
	"net"
	"monitor/agent/base"
	//"io/ioutil"
	"io/ioutil"
)
func echoServer(c net.Conn) {
	buf := make([]byte, 512)
	nr, err := c.Read(buf)
	if err != nil {
		return
	}
	data := buf[0:nr]
	println("Server got:", string(data))
	//_, err = c.Write(data)
	//if err != nil {
	//	panic("Write:"+ err.Error())
	//}
}
func ListenStart() {
	l, err := net.Listen("unix",base.ReadAgentConfig("jvm","sock"))
	//defer l.Close()
	if err != nil {
		println("listen error", err.Error())
		return
	}
	for{
		fd, err := l.Accept()
		if err != nil {
			println("accept error", err.Error())
			return
		}
		//go echoServer(fd)
		//buf,err := ioutil.ReadAll(fd)
		for{
			buf := make([]byte, 512)
			nr ,_ := fd.Read(buf)
			if err != nil {
				return
			}
			data := buf[0:nr]
			println("Server got:", string(data))
		}


	}
}