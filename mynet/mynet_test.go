package mynet

import (
	"fmt"
	"html"
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

func TestMyHttpListen(t *testing.T) {
	http.HandleFunc("/hijack", func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Don't forget to close the connection:
		defer conn.Close()
		bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
		bufrw.Flush()
		s, err := bufrw.ReadString('\n')
		if err != nil {
			log.Printf("error reading string: %v", err)
			return
		}
		fmt.Fprintf(bufrw, "You said: %q\nBye.\n", s)
		bufrw.Flush()
	})
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
