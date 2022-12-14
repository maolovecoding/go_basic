package main

import "fmt"

func demo91() {
	// 多个defer 类似于栈结构 后进先出 => ww ls zs
	// defer 是return 后才调用的
	name := "zs"
	defer fmt.Println(name)
	name = "ls"
	defer fmt.Println(name)
	name = "ww"
	defer fmt.Println(name)
}

func main() {
	demo91()
}
