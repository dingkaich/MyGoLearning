package mylog

import (
	"log"
	"os"
)

func MylogMain1() {
	log.Println("mylog")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("file")
	f, _ := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0640)

	log.SetOutput(f)
	log.Println("this is file")
}
