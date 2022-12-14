package main

import "fmt"

// 函数的参数中如果相邻变量的类型相同，则可以省略类型，例如：
func demo101(a, b int) int {
	return a + b
}

// 可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。

// 注意：可变参数通常要作为函数的最后一个参数。
func demo102(args ...int) int {
	// 可变参数 是一个切片
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

// Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。
func demo103(a, b int) (int, int) {
	return a + b, a - b
}

// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
func demo104(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return sum, sub
}

// 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
func demo105(x string) []rune {
	if x == "" {
		fmt.Println("-----")
		return nil
	}
	str := make([]rune, len(x))
	for _, char := range x {
		str = append(str, char)
	}
	return str
}

// 高阶函数 参数是函数 或者返回值是一个函数 可以实现回调函数 以及闭包行为等
func demo106(init int) func(x int) int {
	sum := init
	return func(x int) int {
		sum += x
		return sum
	}
}

// 在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
func demo107() {}

func main10() {
	// fmt.Println(demo101(10, 20))
	// fmt.Println(demo102(10, 20, 30, 40))
	// fmt.Println(demo103(20, 10))
	// fmt.Println(demo104(20, 10))
	// fmt.Println(demo105("abc"), demo105(""))
	// add := demo106(10)
	// fmt.Println(add(20))
	// fmt.Println(add(30))
	// fmt.Println(add(40))
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 5
}
func main101() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	// defer 虽然延迟执行 但是会保存现场 记录defer之前变量的值 然后向下执行
	x := 1
	y := 2
	// x = 1 y = 2
	defer calc("AA", x, calc("A", x, y)) // 1 + 3 = 4
	x = 10
	defer calc("BB", x, calc("B", x, y)) // 10 + 10 + 2 = 22
	y = 20
}
