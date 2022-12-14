package main

import (
	"errors"
	"fmt"
)

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

func main8() {
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

	// var a = make([]string, 5, 10)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// }
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }
	// demo88()
	fmt.Println(demo89())
	fmt.Println(demo892())
}

func demo83(slice []int) {
	slice[0] = 10
}
func demo84() {
	// =======
	s1 := fmt.Sprint("沙河小王子")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)
	// ========
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Println(w)
}

func demo85() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	// 地址不同 只要数组大小一样，元素一样，也是相同的
	fmt.Println(a == b)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
}

func demo86() {
	a := make([]int, 5, 10)
	fmt.Println(a)
	// 追加元素 是在切片已经用默认值填充后开始的，也就是索引5开始追加
	a = append(a, 1, 2)
	fmt.Println(a)
}

func demo87() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 6, 10)
	c[0] = 100
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}

func demo88() {
	name := "go"
	// 类似对这个函数做了快照，当前函数执行完 在执行defer后面的函数 或者语句
	defer fmt.Println(name) // 输出: go
	name = "python"
	fmt.Println(name) // 输出: python
}
func demo89() (i int) {
	// 因为返回值是i 这个i会在堆中，因为形成的闭包（defer），i不会销毁了 所以最终返回值是 11了
	i = 10
	defer func() {
		i = 11
	}()
	return 5
}
func demo892() int {
	i := 10
	defer func() {
		i = 11
	}()
	// 返回值就是 10 因为返回值i上在一个新的栈空间中
	return i
}
