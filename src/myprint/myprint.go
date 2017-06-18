package myprint

import (
	"flag"
	"fmt"
	"myprint"
	"reflect"
	"time"
)

func init() {
	fmt.Println("i am init 1")
}
func init() {
	fmt.Println("i am init 2")
}
func init() {
	fmt.Println("i am init 3")
}

func Myprint() {

	fmt.Println("123")
}

//TestGoRange test go range
func TestGoRange() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

var period = flag.Duration("period", 1*time.Second, "sleep period")

func TestInterface() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

// //channel close
func Mychan() {
	naturals := make(chan int, 1)
	squares := make(chan int, 1)
	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
		fmt.Println("counter", x)
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	// time.Sleep(30 * time.Second)
	for v := range in {
		out <- v * v
		fmt.Println("squarer", v*v)

	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func Mychan2() {
	fmt.Println("mychan2")
	naturals := make(chan int, 0)
	squares := make(chan int)
	// mmap := make(map[int]int,3,5)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func TestNew() {
	fmt.Println("TestNew")
	p := new(int)
	*p = 2
	fmt.Println(p)
	fmt.Println(*p)

}

func TestSlice() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a[:2]=", a[:2])
	fmt.Println("a[2:]=", a[2:5])

}

func MyprintMain() {

	// myprint.TestGoRange()
	// myprint.TestInterface()
	// myprint.Mychan()
	// myprint.Mychan2()
	fmt.Println("good")
	myprint.Myprint()

	myprint.TestNew()
	myprint.TestSlice()
	fmt.Printf("%T\n", 3)
	v := reflect.TypeOf(3)
	fmt.Println(v.String())
	fmt.Println(v)

	t := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(t)          // "3"
	fmt.Printf("%v\n", t)   // "3"
	fmt.Println(t.String()) // NOTE: "<int Value>"
	
}
