package myioutil

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func MyIouitlMain() {

	dir, _ := ioutil.ReadDir("github.com/beego/logs")

	for _, v := range dir {
		switch {
		case v.IsDir():
			log.Println("dir", v.Name(), "\t", v.Size())
		case !v.IsDir():
			log.Println("file", v.Name(), "\t", v.Size())
		default:
			log.Println("unknow file")
		}
	}

	f, _ := os.OpenFile("dingkai.txt", os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	a, _ := ioutil.ReadAll(f)
	log.Println(string(a))
	f.Seek(0, 0)

	b := bufio.NewReader(f)

	for {
		aa, aa1, aa2 := b.ReadLine()
		if aa2 != nil {
			log.Println(aa2)
			break
		}
		log.Println(string(aa), aa1, aa2)
	}

	ioutil.WriteFile("dingkai.txt", []byte("dingkaichdasfasdfsadfdas"), 0777)
	a, _ = ioutil.ReadFile("dingkai.txt")
	log.Println(string(a))
}
