package mytime

import (
	"fmt"
	"testing"
	"time"
)

func Test_timer1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"dingkai"},
	}
	for range tests {

		timer1()
	}
}

func Test_timer2(t *testing.T) {
	c := make(chan bool)
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("run")
		c <- true
		fmt.Println("end")
	}()

	select {
	case <-c:
		fmt.Println("get")
	case <-time.After(time.Second * 10):
		fmt.Println("timeout1")
	}

}

func Test_timer3(t *testing.T) {
	// c := make(chan bool)
	c := make(chan bool, 1)
	c <- true
	fmt.Println("good")

}
