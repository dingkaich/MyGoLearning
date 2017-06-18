package main

import (
	"fmt"
	"log"
	"myos"

	"github.com/beego/logs"
)

func main() {
	fmt.Println("test astaxie's log")
	logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)
	l := logs.GetLogger()
	l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")

	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")

	fmt.Println("test stand log:")
	log.Println("good boy")

	myos.Myhostnanme()

}
