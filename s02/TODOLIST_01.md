# TODO LIST
我们在这里演示使用GO编写一个银行系统
功能：
- 用户登陆
- 存钱
- 取钱
## cli 内存版本
1. cli命令行方式的
2. 数据存储在内存中


## WEB 版本
1. 我们先学习如何使用GO编写一个WEB HELLO WORLD
```go
package main

import (
	"log"

	"net/http"
)

func main() {
	// 添加路由
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	http.HandleFunc("/helloworld", helloWorld)

	// 启动服务器
	if err := http.ListenAndServe("0.0.0.0:8081", nil); err != nil {
		log.Fatalln(err)  // log.Fatalln 发生painc
	}
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))
}
```