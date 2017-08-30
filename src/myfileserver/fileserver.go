package myfileserver

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Index(res http.ResponseWriter, req *http.Request) {
	log.Println("index", req.URL.String())

	// if req.URL.String() != ""

	// defer req.Body.Close()
	tmp, err := template.ParseFiles("./myfileserver/htmlfile/index.html")
	if err != nil {
		res.Write([]byte("get index html error"))
		http.Error(res, "StatusInternalServerError", http.StatusInternalServerError)
		return
	}
	tmp.Execute(res, nil)
	return
}

func ViewFile(res http.ResponseWriter, req *http.Request) {
	// defer req.Body.Close()
	log.Println("view")
	if strings.HasPrefix(req.URL.String(), "/viewfile") {
		http.StripPrefix("/viewfile", http.FileServer(http.Dir("./myfileserver/upload/"))).ServeHTTP(res, req)
	} else {
		http.FileServer(http.Dir("./myfileserver/upload/")).ServeHTTP(res, req)
	}
	return
}

func UploadFile(res http.ResponseWriter, req *http.Request) {
	// defer req.Body.Close()
	switch req.Method {
	//GET
	case http.MethodGet:
		log.Println("get")
		t, err := template.ParseFiles("./myfileserver/htmlfile/file.html")

		if err != nil {
			res.Write([]byte("get upload html error"))
			http.Error(res, "StatusInternalServerError", http.StatusInternalServerError)
			return
		}
		t.Execute(res, "上传文件")
	//POST
	case http.MethodPost:
		log.Println("post")
		req.ParseMultipartForm(1 << 10)         //设置下内存缓存大小，默认为1G
		f, h, err := req.FormFile("uploadfile") //该值有html中定义
		// defer f.Close()
		if err != nil {
			fmt.Fprintln(res, "uploadfile", h.Filename, " failed")
			return
		}

		file, err := os.OpenFile("./myfileserver/upload/"+h.Filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		defer file.Close()
		if err != nil {
			fmt.Fprintln(res, "create ", h.Filename, " in server failed")
			http.Error(res, "StatusInternalServerError", http.StatusInternalServerError)
			return
		}
		_, err = io.Copy(file, f)
		if err != nil {
			fmt.Fprintln(res, "write ", h.Filename, " in server failed")
			http.Error(res, "StatusInternalServerError", http.StatusInternalServerError)
			return
		}
		filedir, _ := filepath.Abs("./myfileserver/upload/" + h.Filename)
		fmt.Fprintf(res, "%v", h.Filename+"上传完成,服务器地址:"+filedir)
	default:
		fmt.Fprintln(res, "only support get and post")
		return

	}

	return
}

func Myftpmain() {

	Mux := http.NewServeMux()
	//根目录是主页
	Mux.HandleFunc("/", Index)
	Mux.HandleFunc("/viewfile", ViewFile)
	Mux.HandleFunc("/viewfile/", ViewFile)
	Mux.HandleFunc("/uploadfile", UploadFile)

	server := http.Server{
		Addr:         ":6060",
		Handler:      Mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
