package main

import "fmt"

// 语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型
type struct1 struct {
	name, city string // 定义结构体的字段 同样类型的可以写在一行
}

// 只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
// 结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
var s121 struct1

// 在定义一些临时数据结构等场景下还可以使用匿名结构体。
var s122 struct{ name string }

// 我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。
var s123 = new(struct1)

func main12() {
	(*s123).city = "河南" // / *操作优先级低于 .操作 . 会自动解引用
	s123.city = "河南"
	fmt.Println(s123)
	fmt.Println((*s123).city)
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	s124 := &struct1{}
	fmt.Println(s124)
}

type student struct {
	name string
	age  int
}

func main122() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	// 不能使用for range 遍历 for _, stu := range stus
	// &stu每次取得地址都是循环不变量的地址，因为for range变量每次都会把值拷贝到循环不变量上，地址是不变的
	for i := 0; i < len(stus); i++ {
		m[stus[i].name] = &stus[i]
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

type Person2 struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person2) SetDreams(dreams []string) {
	p.dreams = dreams
	// 也可以这里对引用类型 比如切片进行拷贝
	// p.dreams = make([]string, len(dreams))
	// copy(p.dreams, dreams)
}

func main123() {
	p1 := Person2{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	// p1.SetDreams(data)

	//  你真的想要修改 p1.dreams 吗？
	// data[1] = "不睡觉"
	// fmt.Println(p1.dreams) // ?
	copyData := make([]string, len(data))
	copy(copyData, data)
	p1.SetDreams(copyData)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ✅
}
