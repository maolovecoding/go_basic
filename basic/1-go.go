package main

import "fmt"

func demo1() {
	var age int = 10
	fmt.Println(age)
	// & 指针
	a := new(int)
	*a = 10
	fmt.Println(a, *a)
	len := 7
	for i := 0; i <= len; i++ {
		for j := len - i; j >= 0; j-- {
			fmt.Print(" ")
		}
		for j := 0; j <= 2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func demo2() {
	test := "abc"
	// 不需要break
	switch test {
	case "abc", "a":
		fmt.Println(test)
	case "ab":
		fmt.Println(test)
	}
	switch {
	case test == "abc":
		println(test)
	}
}

func demo3() {
	for i := 0; i <= 10; i++ {
		for j := i; j <= 10; j++ {
			println(j)
			if j == 8 {
				goto myBreak
			}
		}
	}
	return
myBreak:
	fmt.Println("结束了")
}

func demo4() {
	var res [10]int
	res[0] = 10
	fmt.Println(res)
	var s = make(map[string]string)
	s["name"] = "zs"
	fmt.Println(s)
	str := fmt.Sprintf("abc=%d", 10)
	fmt.Println(str)
}

func demo5() {
	var res []int
	for i := 0; i <= 1000; i++ {
		res = append(res, i)
	}
	fmt.Println("元素的个数为：", len(res))
}

func demo6() {
	var res = make(map[string]string)
	res["a"] = "abc"
	res["b"] = "abc"
	res["c"] = "abc"
	for key, value := range res {
		fmt.Println(key, value)
	}
}

func main1() {
	demo6()
}
