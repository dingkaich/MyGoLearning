package main

import (
	"mylogrus"
)

// //	"github.com/beego/logs"
// var testmap map[string]int

// //slice的坑要牢记
// func mytest() {
// 	var a = make([]int, 0, 4)
// 	fmt.Println(len(a), cap(a))

// 	// fmt.Printf("%p\n", &a[0])
// 	a = append(a, 1, 2)
// 	fmt.Println(len(a), cap(a), a)

// 	fmt.Printf("%p\n", &a[0])
// 	b := append(a, 1, 2, 3, 12, 31, 21, 31, 31, 13, 13, 31, 31, 31)
// 	fmt.Println(len(a), cap(a), a)

// 	fmt.Printf("%p\n", &a[0])
// 	fmt.Println(len(b), cap(b), b)
// 	fmt.Printf("%p\n", &b[0])
// 	a = append(a, 0)
// 	fmt.Println(a, b)

// 	if testmap == nil {
// 		fmt.Println("make map")
// 		testmap = make(map[string]int)
// 	}

// 	testmap["adsf"] = 1
// 	aaaa := myio.GetActiveLicenseKeyItem()
// 	if aaaa == nil {
// 		fmt.Println("nil")
// 	}
// 	fmt.Println(aaaa, "==", len(aaaa))
// 	// var aa [10]int
// 	// aa = append(aa, 1)
// 	// fmt.Println(aa)
// }

func main() {
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
	mylogrus.MylogrusMain()

	// myflags.MyflagsMain()

}
