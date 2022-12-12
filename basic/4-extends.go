package main

import "fmt"

type Person struct {
	age int
}
type CanEat struct {
	name string
}
type Student3 struct {
	Person
	gender string
	CanEat
}

func (c *CanEat) eat() {
	fmt.Println("我在吃饭")
}

func newStu3(name string, age int, gender string) *Student3 {
	return &Student3{
		Person{
			age: age,
		},
		gender,
		CanEat{
			name: name,
		},
	}
}

func main4() {
	stu := newStu3("zs", 22, "难")
	stu.eat()
	fmt.Println(*stu, stu.name, stu.age)
}
