package main

import "fmt"

func demo81() {
	var a int8 = 127  // 最大127
	var b uint8 = 255 // 255
	fmt.Println(a, b)
}

func main() {
	demo81()
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice为 %d, 其长度为: %d，容量为：%d\n", myslice, len(myslice), cap(myslice))
	// fmt.Println(myslice[3]) // 为什么这里不能访问？
	myslice = myslice[0:cap(myslice)]
	fmt.Printf("myslice的第四个元素为: %d", myslice[3]) // 这里可以访问
}
