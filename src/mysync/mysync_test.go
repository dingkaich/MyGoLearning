package mysync

import (
	"log"
	"sync"
	"testing"
)

func TestMySyncMain(t *testing.T) {
	var once sync.Once
	onceBody := func() {
		log.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
			log.Println("func", i)
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
		log.Println(i)
	}
}
