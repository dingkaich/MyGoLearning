package mychannel

import (
	"fmt"
	"time"
)

var g chan struct{} = make(chan struct{}, 5)

func MychannelMain() {
	var c chan struct{}
	c = make(chan struct{}, 1)

	for i := 0; i < 5; i++ {
		go gorounte(i, c)
	}
	time.Sleep(time.Second * 5)
	c <- struct{}{}
	g <- struct{}{}
	g <- struct{}{}
	g <- struct{}{}
	g <- struct{}{}

	go gorounte1(10, c)
	fmt.Println("end program")
	select {}

}

func gorounte(i int, c chan struct{}) {
	fmt.Println("this is gorouine:", i)
	for {
		select {
		case <-c:
			fmt.Println("begin goroutine:", i)
			c <- struct{}{}
			fmt.Println("end goroutine:", i)
			return
		default:
			time.Sleep(time.Second * 1)
			//work
		}

	}

}

func gorounte1(i int, c chan struct{}) {
	fmt.Println("this is gorouine:", i)

	var gg chan struct{} = make(chan struct{}, 5)
	var ggg chan struct{} = make(chan struct{}, 5)

	gg <- struct{}{}

	ggg <- struct{}{}
	ggg <- struct{}{}
	ggg <- struct{}{}
	ggg <- struct{}{}
	for {
		select {
		case <-g:
			fmt.Println("get g", i)
		case <-gg:
			fmt.Println("get gg", i)
		case <-ggg:
			fmt.Println("get ggg", i)

		default:
			time.Sleep(time.Second * 1)
			//work
		}

	}

}
