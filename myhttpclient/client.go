package myhttpclient

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var client *http.Client

var cfg *clientconfig
var runflag = make(chan struct{}, 1024)

var cnt int64

func get(i int, url string) {

	for {
		<-runflag
		rsp, err := client.Get(url)
		if err != nil {
			log.Println(err)
			return
		}
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, rsp.Body)
		if atomic.AddInt64(&cnt, 1)%100 == 0 {
			log.Println(time.Now())
		}
		//log.Printf("i=%d %s", i, buf.String())
	}

}

func tipsctl() {
	t := time.Tick(time.Second)
	for {
		select {
		case <-t:
			for index := 0; index < cfg.tips; index++ {
				runflag <- struct{}{}
			}

		}
	}
}

func MyhttpclientMain() {
	cfg = getconfig("")
	client = &http.Client{}
	for i := 0; i < cfg.caps; i++ {
		go get(i, "http://localhost:8080/api/message")
	}
	go tipsctl()
	select {}

}
