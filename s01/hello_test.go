package s01

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"testing"

	"github.com/dollarkillerx/Go-Newcomer-Development-Program/s01/hello"
)

func TestHello(t *testing.T) {
	hello.HelloWorld()
}

type config struct {
	Name string `yaml:"name"` // 后面的是tag  表示符 让yaml解析器知道 这个key对应文件那个val(通过反射实现的)
	Age  int    `yaml:"age"`
}

func TestYaml(t *testing.T) {
	fileBytes, err := ioutil.ReadFile("./config.yaml") // 这里使用ioutil来去读这个文件  ioutil.ReadFile会将此文件完整的读取到内存中 ,如果是大文件切忌使用改API
	if err != nil {
		log.Fatalln(err)
	}
	cfg := &config{}
	if err := yaml.Unmarshal(fileBytes, cfg); err != nil { // 调用Unmarshal方法进行解析  传入 yaml文件,要解析成的struct的指针
		log.Fatalln(err)
	}

	log.Println(cfg) // 打印我们解析的结构体
}
