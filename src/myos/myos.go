package myos

import (
	"log"
	"os"
)

func Myhostnanme() {

	// hostname, _ := os.Hostname()
	// log.Println(os.Hostname())
	// log.Println(os.Getpagesize())
	log.Println(os.Environ())
	// //os.Clearenv()
	// os.Setenv("dingkai", "good/boy")
	// log.Println(os.Environ())
	// log.Println(os.ExpandEnv("$dingkai=11"))
	// log.Println(os.Getegid())
	// log.Println(os.Getuid())
	// f, _ := os.Stat("test.log")
	// log.Println(f)
	// f, _ = os.Lstat("test.log")
	// log.Println(f)
	// log.Println(f.Mode())
	log.Println(os.IsPathSeparator(','))
	os.MkdirAll("dingkai", 0777)
	os.Remove("dingkai")
	log.Println(os.TempDir())
	// file, _ := os.Create("dingki.txt")
	// file.Close()
	file1, err := os.OpenFile("dingki.txt", os.O_CREATE|os.O_RDWR, 0777)
	file1.Write([]byte("good boy"))
	if err != nil {
		log.Print(err)
	}
	log.Println(file1.Fd())
	// file1.Sync()
	file1.Seek(0, 0)
	// var str [1024]byte
	str := make([]byte, 20)
	count, err := file1.Read(str)
	if err != nil {
		log.Println(count, err)
	}

	log.Println(string(str))

	// f1, _ := file.Stat()
	// log.Println(f1.Mode())
	f2, _ := file1.Stat()
	log.Println(f2.Mode())
	file1.Close()

	dir, _ := os.Getwd()

	file1, _ = os.Open(dir)

	f, _ := file1.Readdir(0)
	for _, j := range f {
		log.Println(j.Name(), "\t", j.Mode(), j.ModTime(), j.Size())
		// log.Println(i, j)

	}

}
