package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	body, err := http.Get("http://www.baidu.com")
	// body 会baidu文件句柄
	if err != nil {
		log.Fatalln(err)
	}
	defer body.Body.Close() //  处理完毕 关闭文件句柄 (切忌)
	// 针对小文件读取 我们可以使用ioutil
	bytes, err := ioutil.ReadAll(body.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(bytes)) // 打印我们下载到的文件
}
