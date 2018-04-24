package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	ln, err := net.Listen("tcp", "0.0.0.0:6060")

	if err != nil {
		log.Println("listen failed", err)
	}

	for {
		con, err := ln.Accept()
		if err != nil {
			log.Println("accept failed", err)
		}

		go handle(con)

	}

}

func handle(con net.Conn) {
	defer con.Close()
	remoteip := con.RemoteAddr().String()
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, con)
	log.Println(buf.String())
	var id, ip string
	for _, ke := range strings.Split(buf.String(), "?") {
		if strings.Contains(ke, "myid:") {
			id = strings.TrimPrefix(ke, "myid:")
		}
		if strings.Contains(ke, "myip::") {
			ip = strings.TrimPrefix(ke, "myid:")
		}
	}
	if nil == net.ParseIP(ip) {
		fmt.Fprintf(con, "your ip[%s] cannot be  parsed! please resend right ip!", ip)
		return
	}

	f, err := os.OpenFile("/etc/hosts.allow", os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		log.Printf("RemoteIP[%s],myid[%s], Myip[%s], open file /etc/hosts.allow failed,%v", remoteip, id, ip, err)
		return
	}
	f.WriteString(ip)
	return

}
