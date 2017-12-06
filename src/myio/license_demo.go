package myio

import (
	"fmt"
	"sync"
	"time"
)

type KeyItem struct {
	Keyid     int
	Keyname   string
	KeyValue  int
	Keytype   string
	UsedValue string
}

type FeatureItem struct {
	FeatureId   int
	FeatureName string
	KeyItems    []*KeyItem
	DeadlineDay string
	PeriodDay   string //宽限期
}

type Licensefile struct {
	LicenSerialNo string
	Content       string
	Revoke        string
	ProductName   string
	VersionName   string
	Feature       []*FeatureItem
}

type WorkKeyItem struct {
	LicenSerialNo string
	Revoke        string
	ProductName   string
	VersionName   string
	Keyid         int
	Keyname       string
	KeyValue      int
	Keytype       string
	UsedValue     string
	DeadlineDay   string
	PeriodDay     string //宽限期
	FeatureName   string
}

type LicenseMgr struct {
	Rwlock      sync.RWMutex
	WorkKeyItem map[string]*WorkKeyItem //key工作Licens的key信息，吊销后，就不在这个map里了
	LicenseSet  map[string]*Licensefile //lsn所有license的文件信息，本文件仅仅用于保存，但是会更新激活状态
}

var LicenMgr *LicenseMgr

//把值cp出去
func GetActiveLicenseFileInfo() []Licensefile {
	var c chan string
	c = make(chan string, 10)
	c <- "sadfa"
	return nil
}

type task struct {
	Tasklen uint
	TaskQue []chan int
}

var t task

func StartRoutine() {
	t = task{
		Tasklen: 32,
		TaskQue: make([]chan int, 32),
	}

	for index := range t.TaskQue {
		t.TaskQue[index] = make(chan int, 10240)
		// go workroutine(index)
	}

	go sendmsg()
	go monitor()
	time.Sleep(time.Second * 3)

	for index := range t.TaskQue {
		// t.TaskQue[index] = make(chan int, 10240)
		go workroutine(index)
	}
	select {}
	// select {
	// case <-tt:
	// 	for k := range t.TaskQue {
	// 		fmt.Println("Routine:", k, "len:", len(t.TaskQue[k]))
	// 	}
	// }

}

func monitor() {
	tt := time.Tick(time.Second)
	var old = [32]int{0}
	var new = [32]int{0}
	for _ = range tt {
		fmt.Println("===========")
		for k := range t.TaskQue {
			new[k] = len(t.TaskQue[k])
			fmt.Println("Routine:", k, "hanlde:", old[k]-new[k], "len:", new[k])
			old[k] = new[k]
		}
	}
}

func sendmsg() {
	tt := time.Tick(time.Second * 10)
	for _ = range tt {
		for i := 0; i < 200*32; i++ {
			select {
			case t.TaskQue[i%32] <- i:
			default:
				fmt.Println("full:", i%32)
			}
		}
	}

}

func workroutine(index int) {
	for k := range t.TaskQue[index] {
		realwork(k)
	}
}

func realwork(str int) {
	time.Sleep(time.Millisecond * 150)
}
