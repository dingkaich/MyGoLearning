package myfileserver

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Deafult(res http.ResponseWriter, req *http.Request) {
	log.Println("Deafult", req.URL.String())
	if req.URL.Path == "" || req.URL.Path == "/" {
		Index(res, req)
	} else {
		ViewFile(res, req)
	}

	return
}

func Index(res http.ResponseWriter, req *http.Request) {

	tmp, err := template.ParseFiles("./myfileserver/htmlfile/index.html")
	if err != nil {
		res.Write([]byte("get index html error"))
		http.Error(res, "StatusInternalServerError", http.StatusInternalServerError)
		return
	}
	tmp.Execute(res, nil)

}

func ViewFile(res http.ResponseWriter, req *http.Request) {
	// defer req.Body.Close()
	log.Println("view")
	if strings.HasPrefix(req.URL.String(), "/viewfile") {
		http.StripPrefix("/viewfile", http.FileServer(http.Dir("./myfileserver/upload/"))).ServeHTTP(res, req)
	} else {
		res.Header().Set("Content-Type", "application/octet-stream") //设置文件下载类型
		http.FileServer(http.Dir("./myfileserver/upload/")).ServeHTTP(res, req)
		log.Println(req.Header.Get("Content-Type"), "|", res.Header().Get("Content-Type"))
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
		req.ParseMultipartForm(1 << 10) //设置下内存缓存大小，默认为1G
		log.Println(req)

		f, h, err := req.FormFile("uploadfile") //该值有html中定义
		// defer f.Close()
		if err != nil || f == nil || h == nil {
			// if req.Form
			fmt.Fprintln(res, "uploadfile", func(i *multipart.FileHeader) string {
				if i == nil {
					return "N/A"
				} else {
					return i.Filename
				}
			}(h), " failed")
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
		fmt.Fprintf(res, "%v\n", h.Filename+"上传完成,服务器地址:"+filedir)
	default:
		fmt.Fprintln(res, "only support get and post")
		return

	}

	return
}

func Myftpmain() {

	Mux := http.NewServeMux()
	//根目录是主页
	Mux.HandleFunc("/", Deafult)
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
