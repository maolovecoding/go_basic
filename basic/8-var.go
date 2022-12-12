package main

import "fmt"

func demo81() {
	var a int8 = 127  // 最大127
	var b uint8 = 255 // 255
	fmt.Println(a, b)
}

func demo82() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
}

func main() {
	// demo81()
	// var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// myslice := numbers4[4:6:8]
	// fmt.Printf("myslice为 %d, 其长度为: %d，容量为：%d\n", myslice, len(myslice), cap(myslice))
	// // fmt.Println(myslice[3]) // 为什么这里不能访问？
	// myslice = myslice[0:cap(myslice)]
	// fmt.Printf("myslice的第四个元素为: %d", myslice[3]) // 这里可以访问
	// demo82()

	// slice := [3]int{1, 2, 3}
	// demo83(slice[0:1])
	// fmt.Println(slice)

	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func demo83(slice []int) {
	slice[0] = 10
}
