package myhttp

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func MyhttpMain() {
	// listen0()
	// listen1()
	// listen2()
	// learn3()
	// learn4()
	learn5()
}

///////////////////////////////////////////////////////

func listen0() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/hello/dingkai", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("good boy"))
	})

	http.Handle("/", http.FileServer(http.Dir("F:\\百度云盘")))
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("F:\\百度云盘\\魏小玲891126"))))
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
	f, err := os.OpenFile("dingkai.html", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer f.Close()
	if err != nil {
		io.Copy(os.Stderr, res.Body)
	} else {
		io.Copy(f, res.Body)
	}
	var head string
	// log.Println(res.Header)
	head = fmt.Sprintf("%v", res.Header)
	head = http.CanonicalHeaderKey(head)
	log.Println(head)
	log.Println(http.DetectContentType([]byte(head)))

	res.Body.Close()
	res, err = http.Post("http://www.baidu.com", "text/plain", bytes.NewReader([]byte("sb baidu")))
	log.Println("post")
	if err != nil {
		log.Println(err)
	}
	if f != nil {
		io.Copy(f, res.Body)
	} else {
		io.Copy(os.Stderr, res.Body)
	}

	head = ""
	// log.Println(res.Header)
	head = fmt.Sprintln(head, res.Header)
	head = http.CanonicalHeaderKey(head)
	log.Println(head)
	log.Println(http.DetectContentType([]byte(head)))
	res.Body.Close()
	// log.Println(head)

}

///////////////////////////////////////////////////////

func learn3() {
	// buf := bytes.NewBuffer(nil)
	// buf.WriteString("dingkai")

	req, err := http.Get("http://www.baidu.com")
	// defer req.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(req)
	f, _ := os.OpenFile("dingkai.html", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer f.Close()
	io.Copy(f, req.Body)
	f.WriteString("============cookies===================\n")

	f.WriteString("cookielen=" + strconv.Itoa(len(req.Cookies())) + "\n")
	for k, v := range req.Cookies() {
		f.WriteString(strconv.Itoa(k) + "::\n" + "nameis" + v.Name + "\n" + v.String() + "\n")
	}

}

//////////////////////////////////////////////////////////////////////////////////

func learn4() {

	http.Handle("/", http.FileServer(http.Dir("E:\\")))
	http.ListenAndServe(":12345", nil)
	// http.Handle("/LinuxIMG/", http.StripPrefix("/LinuxIMG/", http.FileServer(http.Dir("F:\\LinuxIMG"))))

}

//////////////////////////////////////////////////////////////////////////////////
//isfileexist
//返回true 文件存在，返回false文件不存在
func isfileexist(name string) bool {
	f, err := os.Open(name)
	defer f.Close()

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true

}

func readfile(res http.ResponseWriter, req *http.Request) {
	//提取请求的URL
	filename := "G:/" + req.URL.Path
	filename, _ = filepath.Abs(filename)
	strings.Replace(filename, "/", "\\", -1)

	if isfileexist(filename) {

		http.ServeFile(res, req, filename)
	} else {
		log.Println(filename, "not exist")
	}

	return

}

func learn5() {

	http.HandleFunc("/", readfile)
	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("good boy"))
	})

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

////////////////////////////////////////////////////////////////////////////////////////////
func learn6() {
	addr := "localhost:8080"

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		str := "if you want to view files,please go to  " + addr + "/viewfile\n"
		str1 := "if you want to upload files,please go to  " + addr + "/upload\n"

		res.Write([]byte(str + str1))

	})

	http.ListenAndServe(addr, nil)

}
