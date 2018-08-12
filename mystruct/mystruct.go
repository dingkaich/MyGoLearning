package mystruct

import "fmt"

type Mytest struct {
	Name  string
	Value string
}

func getvalue(name, value string, c *Mytest) {
	fmt.Printf("&c %p c %p c %v \n", &c, c, c)
	cc := &Mytest{
		Name:  name,
		Value: value,
	}
	fmt.Printf("&cc %p cc %p cc %v \n", &cc, cc, cc)
	*c = *cc
	fmt.Printf("&c %p c %p c %v \n", &c, c, c)

}

func Mystructmain() {
	c := &Mytest{}
	fmt.Printf("&c %p c %p c %v \n", &c, c, c)
	getvalue("dingkai", "kai", c)
	fmt.Printf("&c %p c %p c %v \n", &c, c, c)
}
