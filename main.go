package main

import "MyGoLearning/myhttpclient"

func main() {
	// log.Printf("%s sadf 100%%s \n", "apple")
	// logs.Info("%s sadf 100%%s", "apple")
	//fmt.Println(os.Args[1:])
	// log.Println("test astaxie's log")
	// logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)
	// l := logs.GetLogger()
	// l.Println("this is a message of http")
	// //an official log.Logger with prefix ORM
	// logs.GetLogger("ORM").Println("this is a message of orm")

	// logs.Debug("my book is bought in the year of ", 2016)
	// logs.Info("this %s cat is %v years old", "yellow", 3)
	// logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	// logs.Error(1024, "is a very", "good game")
	// logs.Critical("oh,crash")

	// log.Println("test stand log:")
	// log.Println("good boy")
	// logs.Info("god")

	//  myos.Myhostnanme()
	//  myos.Myhostnanme()
	//fmt.Println("dsga")
	// // myos.Myosexec()
	// // mystrings.Mystringsmain()
	// // mybytes.MybythesMain()
	//  myio.MyioMain()

	// myio.ChannelLens(10, 20, 30)
	// myioutil.MyIouitlMain()
	// mytime.MyTimeMain()
	// var i int = 0
	// time.After(time.Second*5)
	// for {
	// 	i++
	// 	fmt.Println("good", i)
	// }
	//mylog.MylogMain1()
	//mynet.MyNetMain()
	//myhttp.MyhttpMain()
	// myfileserver.FileserverMain()
	// myredis.MyRedisMain()
	// mycawler.MyMain()
	// select {}
	// mytest()
	//myio.StartRoutine()
	//myprint.Mymapmain()
	// for i := 0; i < 10; i++ {
	// 	a := rand.New(rand.NewSource(1)).Intn(100)
	// 	fmt.Println(a)
	// }
	// mycrypt.MycryptMain()
	// mylogrus.MylogrusMain1()
	//myairbrake.Myairbrakemain()

	//myflags.MyflagsMain()

	//myhttp.MyhttpMain()
	// myhttp.Myclient1()
	// mychannel.MychannelMain()
	// mystruct.Mystructmain()

	// http.Handle("/viewfile", http.StripPrefix("/viewfile", http.FileServer(http.Dir("/Users/dingkai/Golang"))))
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// http.ListenAndServe(":8080", http.Handle(pattern, handler))
	// mytemplate.MytemplateMain()
	// myreflect.MyReflectMain()
	//myjsonrest.MyjsonrestMain()
	myhttpclient.MyhttpclientMain()
}
