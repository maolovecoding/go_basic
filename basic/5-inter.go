package main

import (
	"fmt"
	"reflect"
)

type fileCache struct {
}

// FetchImage接口实现
func (f *fileCache) FetchImage(url string) string {
	return "从本地缓存中获取图片：" + url
}

// ImageDownloader 图片加载接口
type ImageDownloader interface {
	// FetchImage 获取图片，需要传入图片地址，方法返回图片数据
	FetchImage(url string) string
}

// 定义从网络下载图片的结构体
type netFetch struct {
}

// FetchImage接口实现
func (n *netFetch) FetchImage(url string) string {
	return "从网络下载图片：" + url
}
func demo51() {
	//从本地缓存中获取数据
	var imageLoader ImageDownloader
	imageLoader = new(fileCache)
	data := imageLoader.FetchImage("https://www.example.com/a.png")
	fmt.Println(data)
	if data == "" {
		// 当本地缓存中没有数据时，从网络下载
		var imageLoader2 ImageDownloader
		imageLoader2 = new(netFetch)
		data2 := imageLoader2.FetchImage("https://www.example.com/a.png")
		fmt.Println(data2)
	}
}

// 空接口 任意类型 any
func originOutput(data interface{}) {
	fmt.Println(data, reflect.TypeOf(data))
	d, ok := data.(int)
	if ok {
		fmt.Println(d)
	} else {
		// 不是int 值是int的默认值
		// fmt.Println(d)
	}
}

func main5() {
	originOutput("122")
	originOutput(12)
	// originOutput([]int{1, 2, 3})
	demo52()
}

func demo52() {
	var a1 interface{} = [3]int{1, 2, 3}
	var a2 interface{} = [3]int{1, 2, 3}
	fmt.Println(a1 == a2, &a1, &a2)
}
