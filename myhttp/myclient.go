package myhttp

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func Myclient1() {
	res, _ := http.Get("http://www.163.com")
	reqbody, _ := ioutil.ReadAll(res.Body)
	buff := bytes.NewBuffer(reqbody)
	log.Println(buff.String())

}

func Myclient12() {
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	c := &http.Client{Transport: t}
	res, err := c.Get("file:///Users/dingkai/Golang/MyGoLearning")
	if err == nil {
		log.Println(err)
	}
	reqbody, _ := ioutil.ReadAll(res.Body)
	buff := bytes.NewBuffer(reqbody)
	log.Println(buff.String())

}
