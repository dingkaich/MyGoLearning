package mynet

import (
	"log"
	"net"
	"net/http"
	"testing"
)

type defaulthttp struct{}

var defaultvalue = defaulthttp{}

func (d *defaulthttp) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("dingkai"))
}

func TestMyTcplisten(t *testing.T) {

	l, err := net.Listen("tcp", ":123456")
	// l, err := net.ListenPacket("ip", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// 等待下一个连接,如果没有连接,l.Accept会阻塞
		_, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		http.Serve(l, &defaultvalue)
	}

	// for {
	// 	// 等待下一个连接,如果没有连接,l.Accept会阻塞
	// 	conn, err := l.Accept()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// 将新连接放入一个goroute里,然后再等下一个新连接.
	// 	go func(c net.Conn) {
	// 		// 显示连接发送来的内容
	// 		io.Copy(os.Stdout, c)
	// 		// 关闭连接
	// 		c.Close()
	// 	}(conn)
	// }

}
