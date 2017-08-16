package myhttp

import (
	"bytes"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func MyhttpMain() {
	// listen0()
	// listen1()
	listen2()
}

///////////////////////////////////////////////////////

func listen0() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/hello/dingkai", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("good boy"))
	})
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

///////////////////////////////////////////////////////

type defaulthttp struct{}

var defaultvalue = defaulthttp{}

func (d *defaulthttp) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("dingkai"))
}
func listen1() {
	log.Println("go to listen1")
	l, err := net.Listen("tcp", ":12321")
	// l, err := net.ListenPacket("ip", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	// for {
	// 	// 等待下一个连接,如果没有连接,l.Accept会阻塞
	// 	_, err := l.Accept()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	http.Serve(l, &defaultvalue)
	// }
	http.Serve(l, &defaultvalue)

}

///////////////////////////////////////////////////////

func listen2() {
	res, err := http.Get("http://www.baidu.com/s?wd=golang")
	if err != nil {
		log.Println(err)
		return
	}

	io.Copy(os.Stderr, res.Body)

	res.Body.Close()
	res, err = http.Post("http://www.baidu.com", "text/plain", bytes.NewReader([]byte("sb baidu")))
	log.Println("post")
	if err != nil {
		log.Println(err)
	}
	io.Copy(os.Stderr, res.Body)
	log.Println(res.Body)

}
