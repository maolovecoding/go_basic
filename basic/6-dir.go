package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func demo61() {
	str := "我要写入数据了"
	chars := []byte(str)
	os.WriteFile("a.txt", chars, 01411)
	os.Stdout.Write(chars)
}

func demo62() {
	file, err := os.Create("a.txt")
	if err != nil {
		// 创建文件失败
	}
	defer file.Close()
	file.Write([]byte("你好，我是哈哈哈哈"))
	file.Sync()
}

func demo63() string {
	resp, err := http.Get("http://127.0.0.1:8888/")
	//使用断言关闭网络请求
	defer resp.Body.Close()
	//使用ioutil工具包获取服务端响应数据
	// buf := make([]byte, 6)
	// resp.Body.Read(buf)
	// fmt.Println(string(buf))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取网络响应失败：", err)
		panic(err)
	}
	return string(body)
}

func main6() {
	str := demo63()
	fmt.Println(str)
}
