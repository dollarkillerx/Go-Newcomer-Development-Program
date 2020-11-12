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
