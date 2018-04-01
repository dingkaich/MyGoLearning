package myio

import (
	"fmt"
	"math/rand"
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
		for i := 0; i < 100*32; i++ {
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

var a = rand.New(rand.NewSource(23))

func realwork(str int) {
	st := a.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(150+st))
}

//Pop点里的整体值 40
//每个设备的capset 设备数是变化的 1 2 3 4 5 3 2 5 6 7 3 | 2 1
//分配要均匀

//每5秒刷新一次，保持于数据库相同即可
type PopLicense struct {
	Popid  int
	LicKey map[string]*struct {
		MaxVal     int //pop点的信息一般是不动的，可以不加锁
		ReserveVal int
	}
}

var m map[int]*PopLicense

// 注册的时候直接写内存,如果内存中的数据没有，更新内存，同时使用merge命令更新DB。DB可能由于竞争写失败，没有关系
// 直接忽略即可，但是内存中的数据就不要保留了。我们可以等待下一次写入。
// 这里的数据可以每3秒从DBload一次
type DevLicense struct {
	Devid  int
	LicKey map[string]*struct {
		lock     sync.RWMutex
		MaxVal   int
		AllocVal int
		UsedVal  int
	}
}

var n map[int]*DevLicense

func RefeshPopLicense() {

}

func computePoplicense(popid int) {
	// poplicense := getpoplicense(popid)

	// getdevinfo()

}

func getpoplicense(popid int) int {
	return 40
}
