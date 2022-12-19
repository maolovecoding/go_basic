package main

import (
	"encoding/json"
	"fmt"
)

// tag 的使用
type St1 struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
	Addr string `json:"addr,omitempty"` // 这里的 omitempty 在进行json话的时候 如果该属性的值是默认值（也就是初始值，或者说是这个类型的零值）那么该属性会被剔除
}

func demo131() {
	s1 := St1{
		"张三", 22, "",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("出现错误")
		return
	}
	fmt.Println(s1, string(data))
}

func main() {
	demo131()
}
