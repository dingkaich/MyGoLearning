package main

import (
	"log"
	"myio"

	"github.com/beego/logs"
	"myos"
)

//	"github.com/beego/logs"

func main() {
	log.Println("test astaxie's log")
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

	log.Println("test stand log:")
	log.Println("good boy")
	logs.Info("god")

	 myos.Myhostnanme()
	// myos.Myosexec()
	// mystrings.Mystringsmain()
	// mybytes.MybythesMain()
	// myio.MyioMain()
	myio.ChannelLens(10, 20, 30)
	// myioutil.MyIouitlMain()
	// mytime.MyTimeMain()
	// var i int = 0
	// time.After(time.Second*5)
	// for {
	// 	i++
	// 	fmt.Println("good", i)
	// }

	// mynet.MyNetMain()

	select {}
}
