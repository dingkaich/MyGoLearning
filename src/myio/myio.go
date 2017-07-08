package myio

/*
io包不是具体的实现了。主要是接口之间的调用。


*/

import (
	"bytes"
	"io"
	"log"
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
