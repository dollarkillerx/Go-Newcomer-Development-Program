package hello

import "fmt"

func init() {
	fmt.Println("Hello World")
}

func HelloWorld() { // 大写开头为公开的方法
	helloWorld()
}

func helloWorld() { // 小写开头为私有方法  只能在同包下才能调用
	fmt.Println("Hello world")
}
