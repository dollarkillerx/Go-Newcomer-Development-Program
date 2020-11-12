## http client 
### client
- 爬取百度首页
```go
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
    // ioutil.ReadAll 会直接读取到内存当中 如果是大文件 会导致内存突然占用过大
	bytes, err := ioutil.ReadAll(body.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(bytes)) // 打印我们下载到的文件
}
```
- 下载大文件
```go
package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

// 下载go1.15到文件 go1.15.4.linux-amd64.tar.gz中
func main() {
	filaUrl := "https://studygolang.com/dl/golang/go1.15.4.linux-amd64.tar.gz"

	body, err := http.Get(filaUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer body.Body.Close()

	// 创建目标存储文件
	// 返回文件句柄和error
	outfile, err := os.Create("go1.15.4.linux-amd64.tar.gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer outfile.Close()

	reader := bufio.NewReader(body.Body) // 创建读取迭代器
	buf := make([]byte, 512)             // 创建512字节的缓冲区

	for {
		i, err := reader.Read(buf) // 我们次读取512字节   i为游标(装满返回512,反之返回装到了的游标)
		if err != nil {
			if err == io.EOF { // 返回io.EOF 代表读取完毕
				break
			}
			// 反之就是错误
			log.Fatalln(err)
		}

		outfile.Write(buf[:i])
	}

	log.Println("Download Success")
}
```
- 进阶
阅读： https://i6448038.github.io/2017/11/11/httpAndGolang/
阅读源码： https://github.com/dollarkillerx/urllib