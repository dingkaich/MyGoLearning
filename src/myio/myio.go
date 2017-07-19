package myio

/*
io包不是具体的实现了。主要是接口之间的调用。


*/

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func MyioMain() {

	//第一个demo
	byt1 := bytes.NewBufferString("dingkai")
	byt2 := bytes.NewBufferString("ddddddd")
	// var byt2 bytes.Buffer
	io.Copy(byt2, byt1)
	log.Println(byt2.String())

	io.WriteString(byt2, "s string")
	log.Println(byt2.String())

	byt1.Reset()
	byt2.Reset()
	byt1.WriteString("dingkai")
	byt2.WriteString("s string")
	bb := make([]byte, 20)
	cc := io.TeeReader(byt1, byt2)
	cc.Read(bb)
	log.Println(string(bb))
	cc.Read(bb)
	log.Println(string(bb))

	byt1.Reset()
	byt2.Reset()
	byt1.WriteString("dingdsadfkai")
	byt2.WriteString("sasdfa string")
	io.MultiReader(byt1, byt2).Read(bb)
	log.Println(string(bb))

	io.MultiReader(byt1, byt2).Read(bb)
	log.Println(string(bb))

	byt1.WriteString("dingdsadfkai")
	io.ReadFull(byt1, bb)
	log.Println(string(bb))

	io.WriteString(byt1, "a1111")
	log.Println(byt1.String())

	// //尝试一下pipe管道的用法
	// pipe_reader, pipe_writer := io.Pipe()
	// pipe_writer.Write(data []byte)
}

func MyReadAtLeast() (n int, err error) {
	n = 1
	err = nil
	return
}

func MyMultiRead() {
	f1, _ := os.Open("dingkai.txt")
	f2, _ := os.Open("dingki.txt")

	multiReader := io.MultiReader(f1, f2)

	p := make([]byte, 50)
	for {
		n, err := multiReader.Read(p)
		log.Println(string(p))
		log.Println(string(p[0:n]))
		if err == io.EOF {

			log.Println("read end")
			break
		}

	}

	MyChanBuffer()
}

func MyPipe() {
	reader, writer := io.Pipe()
	inputData := []byte("1234567890")
	go writer.Write(inputData)
	writer.Close()
	outputData := make([]byte, 10)
	n, _ := reader.Read(outputData)
	log.Println(string(outputData))
	log.Println("read number", n)
	n, _ = reader.Read(outputData)
	log.Println(string(outputData))
	log.Println("read number", n)

}
func MyPipe1() {
	reader, writer := io.Pipe()
	inputData := []byte("1234567890")

	go func() {
		writer.Write(inputData)

	}()

	writer.Close()
	outputData := make([]byte, 10)
	n, err := reader.Read(outputData)
	if err == io.EOF {
		log.Println("executing read return EOF")
		log.Println("executing read reads number", n)
	}
	n, _ = reader.Read(outputData)
	log.Println(string(outputData))
	log.Println("next read number", n)
}

func MyChanBuffer() {
	chanMessage := make(chan string, 2)
	count := 4
	go func() {
		for i := 1; i <= count; i++ {

			// send to chanMessage
			chanMessage <- fmt.Sprintf("message %d", i)
			log.Println("send message")
		}
	}()
	// Pause the main to let the goroutine sends its messages
	time.Sleep(time.Second * 2)
	for i := 1; i <= count; i++ {
		// receive from chanMessage and print
		log.Println(<-chanMessage)
		time.Sleep(time.Second)
	}
}
