package myhttp

import (
	"log"
	"net/http"
)

func myhttpshandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

//MyhttpsMain
//客户端可以使用以下两个命令来完成双向认证
//curl --cert client.crt --key client.key  --cacert ca.crt https://0.0.0.0:10443
//wget --certificate=client.crt  --private-key=client.key --ca-certificate=ca.crt https://0.0.0.0:10443
func MyhttpsMain() {
	http.HandleFunc("/", myhttpshandler)
	log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
	err := http.ListenAndServeTLS("127.0.0.1:10443", "/Users/dingkai/Golang/CA/server.crt",
		"/Users/dingkai/Golang/CA/server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
