package main

import "fmt"

// Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
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

func main9() {
	demo91()
}
