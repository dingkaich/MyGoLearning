package mystrings

import (
	"log"
	"strings"
)

func Mypnt(v *string) {
	log.Println(v)
	log.Println(*v)
	*v = "kaiding"
}
func Mypnt1(v string) {
	log.Println(v)
	v = "kaiding1"
	log.Println(v)
}

func Mystringsmain() {
	log.Println(strings.EqualFold("s string", "S String"))
	log.Println(strings.HasPrefix("ss string", " s1 "))
	log.Println(strings.Contains("seafood", "foo"))
	log.Println(strings.Contains("seafood", "bar"))
	log.Println(strings.Contains("seafood", ""))
	log.Println(strings.Contains("", ""))
	log.Println(strings.ContainsAny("dingkai", "DINGKaI"))
	log.Println(strings.Count("disaaaangkai", "aa"))
	log.Println(strings.IndexByte("disaaaangkai", 'a'))
	log.Println(strings.IndexRune("丁凯", '凯'))

	dingkai := "丁凯"
	r := []rune(dingkai)
	log.Printf("%c", r[1])

	//以下代码可以证明golang的函数传参是值传递，并非引用
	str := "dingkai"
	Mypnt(&str)
	log.Println(str)

	Mypnt1(str)
	log.Println(str)

	log.Println(strings.ToTitle("loud noises"))
	log.Println("LOUD NOISES")
	log.Println("ba" + strings.Repeat("na", 2))
	log.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	log.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	log.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	log.Printf("%q (nil = %v)\n", z, z == nil)

	s := []string{"foo", "bar", "baz"}
	log.Println(strings.Join(s, ", "))

	reader := strings.NewReader("dingkaich is a good boy")
	log.Println(reader.Len())

	var out2 []byte
	out2 = make([]byte, 100)
	n, err := reader.Read(out2)
	log.Println(n, err, string(out2))

	repalce := strings.NewReplacer("dingkaich", "good")
	log.Println(repalce)
	log.Println(repalce.Replace("ding"))
	log.Println(repalce)
}
