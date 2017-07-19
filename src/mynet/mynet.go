package mynet

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func MyNetMain() {

	eth, _ := net.Interfaces()
	for i, k := range eth {
		log.Println(i, k)
	}

	ipaddr, err := net.ResolveIPAddr("ip", "www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ipaddr)
	Mylistenpackage()

}
func Mylistenpackage() {
	l, err := net.ListenPacket("udp", "127.0.0.1:2000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		b := make([]byte, 1000)
		l.ReadFrom(b)
		log.Println(string(b))
	}

}

func Mylisten() {
	l, err := net.Listen("tcp", ":2000")
	// l, err := net.ListenPacket("ip", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// 等待下一个连接,如果没有连接,l.Accept会阻塞
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 将新连接放入一个goroute里,然后再等下一个新连接.
		go func(c net.Conn) {
			// 显示连接发送来的内容
			io.Copy(os.Stdout, c)
			// 关闭连接
			c.Close()
		}(conn)
	}
}
