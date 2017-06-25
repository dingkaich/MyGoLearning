package mybytes

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func MybythesMain() {

	//看下byte的单位是不是字符串、
	str := "丁凯str"
	bytstr := []byte(str)
	rustr := []rune(str)
	for k, v := range bytstr {
		log.Println(k, v, string(v))
	}
	log.Println(len(bytstr))
	log.Println(len(rustr))

	for k, v := range rustr {
		log.Println(k, v, string(v))
	}
	//通过以上是不是可以说明rune 支持任何字符串的切片而byte不支持呢。

	bytstr = []byte("dingdingding")
	bytstr1 := make([]byte, 6)
	bytstr1[0] = 'd'
	bytstr1[1] = 'i'
	bytstr1[2] = 'g'
	bytstr1[3] = 'd'
	bytstr1[4] = 'g'
	bytstr1[5] = '\n'
	log.Println(bytstr)
	log.Println(bytstr1)
	log.Println(bytes.Compare(bytstr, bytstr1))
	log.Println(bytes.Equal(bytstr, bytstr1))

	byspit := bytes.Split(bytstr, []byte("ng"))
	log.Println(byspit)

	//bytes's reader
	byreader := bytes.NewReader(bytstr1)
	byreader.WriteTo(os.Stderr)
	log.Println(byreader.Len())
	nn, _ := byreader.Read(bytstr)
	log.Println(nn, bytstr)
	err := byreader.UnreadByte()
	log.Println(err)
	log.Println("err")
	byreader.WriteTo(os.Stderr)

	var b bytes.Buffer
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stderr)

	b1 := bytes.NewBufferString("s string")
	log.Println(b1.Len())
	log.Println(b1.Bytes())
	log.Println(b1.String())
	var bbb [100]byte
	bb := bbb[0:]
	bb1 := bbb[10:]
	log.Println(b1.ReadString('t'))
	b1.Read(bb)
	log.Println(string(bb))
	b1.UnreadByte()
	b1.Read(bb1)
	log.Println(string(bb1))
	// b1.Write([]byte("dddd"))
	b1.WriteString("ddddsssssssssssssssssssssssssssss")
	log.Println(b1.String())
	b1.WriteString("eeeeee")
	log.Println(b1.String())
	// b1.Read(bb)
	// log.Println(string(bb))
	b1.WriteTo(os.Stderr)

}
