package cli

import (
	"github.com/dollarkillerx/Go-Newcomer-Development-Program/s02/todolist_pro/conf"
	"github.com/dollarkillerx/Go-Newcomer-Development-Program/s02/todolist_pro/storage"
	"strings"

	"fmt"
	"log"
)

type Cli struct {
	db storage.TodoListStorage
}

func New() *Cli {
	var db storage.TodoListStorage
	if strings.Index(conf.Config.Storage, "MySQL") != -1 {
		db = storage.NewMySQLAccount()
	} else {
		db = storage.NewSimple()
	}

	return &Cli{
		db: db,
	}
}

func (c *Cli) Run() {
	fmt.Println("欢迎来到大雄银行: ")

loop:
	for {
		printMenu()
		fmt.Print("请输入: ")
		var i string
		_, err := fmt.Scan(&i)
		if err != nil {
			log.Fatalln(err)
		}
		switch i {
		case "Ext":
			break loop // 如果只有以break 只会退出switch  还会一直执行for循环
		case "1":
			fmt.Printf("请输入用户名:")
			var username string
			var password string
			_, err := fmt.Scan(&username)
			if err != nil {
				log.Println("输入错误")
				continue
			}
			fmt.Printf("请输入密码:")
			_, err = fmt.Scan(&password)
			if err != nil {
				log.Println("输入错误")
				continue
			}

			if err := c.db.CheckAccount(username, password); err != nil {
				log.Println(err)
				continue
			}
			c.bank(username)
		case "2":
			fmt.Printf("请输入用户名:")
			var username string
			var password string
			_, err := fmt.Scan(&username)
			if err != nil {
				log.Println("输入错误")
				continue
			}
			fmt.Printf("请输入密码:")
			_, err = fmt.Scan(&password)
			if err != nil {
				log.Println("输入错误")
				continue
			}

			if err := c.db.OpenAccount(username, password); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("注册成功")
		default:
			fmt.Println("What fuck？ 你输入个什么玩意?")
		}
	}
	fmt.Println("祝你生活愉快！！")
}

func (c *Cli) bank(account string) {
	for {
		printMenu2()
		fmt.Print("请输入: ")
		var i string
		_, err := fmt.Scan(&i)
		if err != nil {
			log.Fatalln(err)
		}

		switch i {
		case "1":
			balance, err := c.db.BalanceEnquiry(account)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("用户: %s 余额: %f \n", account, balance)
		case "2":
			fmt.Print("请输入需要存多少: ")
			var money float32
			_, err := fmt.Scan(&money)
			if err != nil {
				log.Println(err)
				continue
			}
			if err := c.db.Deposit(account, money); err != nil {
				log.Println(err)
				return
			}
			fmt.Println("存款成功!!!")
		case "3":
			fmt.Print("请输入需要取多少: ")
			var money float32
			_, err := fmt.Scan(&money)
			if err != nil {
				log.Println(err)
				continue
			}
			if err := c.db.Withdrawals(account, money); err != nil {
				if err == fmt.Errorf("余额不足") {

				}
				log.Println(err)
				return
			}
			fmt.Println("存款成功!!!")
		case "Ext":
			return // return直接退出方法
		default:
			fmt.Println("What fuck？ 你输入个什么玩意?")
		}
	}
}

// 打印目录
func printMenu() {
	fmt.Println()
	fmt.Println("=========================")
	fmt.Println("登陆系统请输入1: ")
	fmt.Println("注册请输入2: ")
	fmt.Println("退出请输入Ext: ")
	fmt.Println("=========================")
}

func printMenu2() {
	fmt.Println()
	fmt.Println("=========================")
	fmt.Println("查询余额1: ")
	fmt.Println("存钱2: ")
	fmt.Println("取钱3: ")
	fmt.Println("退出请输入Ext: ")
	fmt.Println("=========================")
}
