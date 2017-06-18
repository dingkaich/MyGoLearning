package myos

import (
	"fmt"
	"os"
)

func Myhostnanme() {

	// hostname, _ := os.Hostname()
	// fmt.Println(os.Hostname())
	// fmt.Println(os.Getpagesize())
	fmt.Println(os.Environ())
	// //os.Clearenv()
	// os.Setenv("dingkai", "good/boy")
	// fmt.Println(os.Environ())
	// fmt.Println(os.ExpandEnv("$dingkai=11"))
	// fmt.Println(os.Getegid())
	// fmt.Println(os.Getuid())
	// f, _ := os.Stat("test.log")
	// fmt.Println(f)
	// f, _ = os.Lstat("test.log")
	// fmt.Println(f)
	// fmt.Println(f.Mode())
	fmt.Println(os.IsPathSeparator(','))
	os.MkdirAll("dingkai", 0777)
	os.Remove("dingkai")
	fmt.Println(os.TempDir())
	// file, _ := os.Create("dingki.txt")
	// file.Close()
	file1, err := os.OpenFile("dingki.txt", os.O_CREATE|os.O_RDWR, 0777)
	file1.Write([]byte("good boy"))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(file1.Fd())
	// file1.Sync()
	file1.Seek(0, 0)
	// var str [1024]byte
	str := make([]byte, 20)
	count, err := file1.Read(str)
	if err != nil {
		fmt.Println(count, err)
	}

	fmt.Println(string(str))

	// f1, _ := file.Stat()
	// fmt.Println(f1.Mode())
	f2, _ := file1.Stat()
	fmt.Println(f2.Mode())
	file1.Close()

	dir, _ := os.Getwd()

	file1, _ = os.Open(dir)

	f, _ := file1.Readdir(0)
	for _, j := range f {
		fmt.Println(j.Name(), "\t", j.Mode(), j.ModTime(), j.Size())
		// fmt.Println(i, j)

	}

}
