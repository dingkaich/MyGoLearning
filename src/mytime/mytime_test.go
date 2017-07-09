package mytime

import (
	"log"
	"testing"
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
	a := []byte{1, 2, 3, 4, 5}
	log.Println(a[0:0])
	log.Println(a[1:3])
}
