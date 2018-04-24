package mytime

import (
	"log"
	"time"
)

func timer1() {
	log.Println("i am timer1")
}

var C chan struct{}

func MyTimeMain() {
	C = make(chan struct{})
	local, _ := time.LoadLocation("Local")

	log.Println(local.String())

	//时间的格式化
	t := time.Now()
	a, b := t.Zone()
	log.Println(a, b)
	log.Println(t.Year())
	log.Println(t.YearDay())
	log.Println(t.Location())
	log.Println(t.Unix())

	s := t.Format("20060102150405 CST")
	log.Println(s)
	t, _ = time.Parse("20060102150405 CST", s)
	log.Println(t)

	// 定时器
	tr := time.NewTimer(time.Second * 5)
	time.AfterFunc(time.Second*2, timer1)
	log.Println("wait 5s")
	select {
	case <-tr.C:
		log.Println("get 5s")
	case <-time.After(3 * time.Second):
		log.Println("timed out")
	}

	tt := time.NewTicker(time.Second * 1)

	go func(tt *time.Ticker) {
		i := 1
		for {
			select {
			case <-tt.C:
				log.Printf("%d\n", i)
				if i++; i > 3 {
					i = 0
					goto end
				}

			}
		}

	end:
		log.Println("end func")
		C <- struct{}{}

	}(tt)

	ttt := time.Tick(time.Second)

	for now := range ttt {
		log.Println(now)
		select {
		case <-C:
			log.Println("get channl")
			goto finish
		default:
		}
	}
finish:
}
