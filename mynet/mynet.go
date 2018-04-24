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

	// go Mylistenpackage()
	// go MyTcplisten()
	// go MyUDPlisten()
	go dialtcp()
	go dialudp()
	go mydial()

}

func mydial() {
	d, err := net.Dial("udp", "127.0.0.1:2000")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintln(d, "dial udp")

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

func MyTcplisten() {
	l, err := net.Listen("tcp", ":2001")
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

func dialtcp() {
	tcpaddr, err := net.ResolveTCPAddr("tcp", ":2001")
	log.Println(tcpaddr, err)

	dialtcp, err := net.DialTCP("tcp", nil, tcpaddr)

	if err != nil {
		log.Println(err)
	}

	fmt.Fprintln(dialtcp, "dingkai is a good boy,this is tcp")

}

func MyUDPlisten() {
	l, err := net.Listen("udp", "127.0.0.1:2000")
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

func dialudp() {
	udpaddr, err := net.ResolveUDPAddr("udp", ":2000")
	log.Println(udpaddr, err)

	dialudp, err := net.DialUDP("udp", nil, udpaddr)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintln(dialudp, "dingkai is a good boy,this is udp")

}
