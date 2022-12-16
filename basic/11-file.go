package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func demo110() {
	// 打开文件 os.open
	file, err := os.Open("../test.js")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()
	// fmt.Print(file)
	// 读取文件 func (f *File) Read(b []byte) (n int, err error)
	// 它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF
	// 使用Read方法读取数据
	temp := make([]byte, 128)
	content := make([]byte, 128*8)
	for {
		n, err := file.Read(temp)
		if err == io.EOF {
			// 来到这里，说明本次读取的是文件末尾，有效字符已经读取完毕了。本次读取是无效的
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			fmt.Println("文件读取失败")
			return
		}
		fmt.Printf("读取了%d个字符\n", n)
		// 防止读取的内容重复，只能合并到读取的n
		content = append(content, temp[:n]...)
	}
	// fmt.Println(string(temp))
	fmt.Println(string(content))
}

// bufio是在file的基础上封装了一层API，支持更多的功能。 读取文件 	// bufio按行读取
func demo111() {
	file, err := os.Open("../test.js")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // 指定分隔符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。 （过时了）
func demo112() {
	content, err := ioutil.ReadFile("../test.js")
	if err != nil {
		fmt.Println("文件读取失败！")
		return
	}
	fmt.Println(string(content))
}

// os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
// func OpenFile(name string, flag int, perm FileMode) flag 是打开的文件模式 读写追加清空等 perm是文件权限 4读2写1执行
func demo113() {
	// os.O_CREATE|os.O_APPEND 创建 + 追加
	file, err := os.OpenFile("../file-test/write1.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer file.Close()
	file.WriteString("你好，顶顶顶顶")    // 写入字符串
	file.Write([]byte("\n我是字节切片")) // 也可以写入字节切片
}

func demo114() {
	file, err := os.OpenFile("../file-test/write1.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("哈哈哈哈，bufio") //将数据先写入缓存
	writer.Flush()                   // 清空缓存区 写入文件 不清空是没办法写入文件的
}

func main() {
	// demo110()
	// demo111()
	// demo112()
	// demo113()
	demo114()
}
