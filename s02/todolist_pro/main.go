package main

import (
	"github.com/dollarkillerx/Go-Newcomer-Development-Program/s02/todolist_pro/cli"
	"github.com/dollarkillerx/Go-Newcomer-Development-Program/s02/todolist_pro/conf"
)

func main() {
	conf.InitConfig()

	c := cli.New()
	c.Run()
}
