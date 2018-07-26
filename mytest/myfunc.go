package mytest

import (
	"errors"
	"fmt"
	"math"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return math.Trunc(a / b), nil
}

func Mytestmain() {
	arr := make([]int, 0, 8)
	arr = []int{1, 2, 3}
	fmt.Println(arr)
	for _, v := range arr {
		arr[3] = 5
		arr = append(arr, v)
	}
	fmt.Println(arr)

}

func MyTestmain1() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)
	fmt.Println("slice addr:", &slice)
	for index, value := range slice {
		fmt.Println(&index, &value, &slice[index])
		// num := value
		myMap[index] = &slice[index]
	}
	fmt.Println("=====new map=====")
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%d,%v\n", key, *value, value)
		fmt.Println(key, "=", *value, value)
	}
}
