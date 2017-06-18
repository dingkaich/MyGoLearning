package myos

import (
	"fmt"
	"os"
)

func Myhostnanme() {

	// hostname, _ := os.Hostname()
	fmt.Println(os.Hostname())
	fmt.Println(os.Getpagesize())
	fmt.Println(os.Environ())
	os.Clearenv()
	os.Setenv("dingk", "good/boy")
	fmt.Println(os.Environ())
}
