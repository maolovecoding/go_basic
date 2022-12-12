package main

import "fmt"

func (s *Student2) toString() string {
	return s.name
}

type Student2 struct {
	name   string
	age    int
	gender bool
}

func newStu(name string, age int, gender bool) *Student2 {
	// 因为是对象  所以我们不是返回对象的副本，而是返回创建好的对象
	return &Student2{
		name: name, age: age, gender: gender,
	}
}

func fn1() {

	s1 := Student2{
		name: "zs", age: 22, gender: true,
	}
	fmt.Println(s1)
}

func fn2() {
	stu := newStu("zs", 22, true)
	fmt.Println(stu)
}

func main3() {
	fn2()
}
