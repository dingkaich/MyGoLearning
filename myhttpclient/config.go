package myhttpclient

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type clientconfig struct {
	ipaddr string `json:"ipaddr"`
	port   string `json:"port"`
	caps   int    `json:"caps"` //保持多少个http连接存在
	tips   int    `json:"tips"` //1s多少个请求出去
}

func getconfig(cfgfile string) *clientconfig {

	if cfgfile == "" {

		return &clientconfig{
			ipaddr: "127.0.0.1",
			port:   "8080",
			caps:   1500,
			tips:   100,
		}
	}

	c, err := ioutil.ReadFile(cfgfile)
	if err != nil {
		log.Println(err)
		return nil
	}
	var cfg = &clientconfig{}
	json.Unmarshal(c, &cfg)
	if err != nil {
		log.Println(err)
		return nil
	}
	return cfg

}
