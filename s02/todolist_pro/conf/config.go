package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
	Storage string // MySQL,Mem  存储层使用 mysql还是内存

	MySQLConf MySQLConf
}

type MySQLConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
	Address  string `json:"address"`
}

var Config *config

func InitConfig() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalln(err)
	}
	var cfg config
	if err := json.Unmarshal(file, &cfg); err != nil {
		log.Fatalln(err)
	}

	Config = &cfg
}