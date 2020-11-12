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
