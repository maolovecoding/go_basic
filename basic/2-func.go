package main

import (
	"fmt"
	"math"
)

func demo21(str string) string {
	return str + "abc"
}

func demo22(res *[]string) {
	*res = append(*res, "str")
}
func demo23() {
	const (
		a = 1 << iota // 1
		b             // 2
		c             // 4
		d             // 8
		e = 3 << iota // 3 * 2 ^ 4
		f
		g
	)
	fmt.Println(a, b, c, d, e, f, g)
}

func demo24() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

func demo25() {
	// 自动推断数组的长度 nice
	var arr = [...]int{1, 2, 3}
	fmt.Println(arr, len(arr), math.Max(float64(arr[0]), float64(arr[1])))
}
func demo26() {
	var arr = new(int)
	*arr = 10
	fmt.Println(arr, &arr, *arr)
}

type fn func(int) int

func demo27(cb fn) int {
	return cb(10)
}

func main2() {
	// fmt.Print(demo21("你好："))

	// defer fmt.Print("素数")
	// defer fmt.Print("查找")
	// var res []string
	// demo22(&res)
	// fmt.Println(res)

	// demo23()
	// demo24()
	// demo25()
	// demo26()

	// fmt.Println(demo27(func(i int) int {
	// 	return i + 10
	// }))
	// stu := Student{name: "zs", age: 22}
	// var stu2 Student = Student{name: "z", age: 33}
	// stu.next = &stu2
	// fmt.Print(stu)
	demo28()
}
func demo28() {
	var arr = []int{1, 2, 3, 4}
	for _, value := range arr {
		fmt.Println(value)
	}
}

type Student struct {
	name string
	age  int
	next *Student
}
