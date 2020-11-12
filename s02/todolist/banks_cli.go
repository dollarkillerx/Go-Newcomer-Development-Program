package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("欢迎来到大雄银行: ")

loop: // 标识符
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

			if err := Bank.Login(username, password); err != nil {
				log.Println(err)
				continue
			}
			bank(username)
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

			if err := Bank.AddUser(username, password); err != nil {
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

func bank(username string) {
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
			balance, err := Bank.ObtainingBalance()
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("用户: %s 余额: %f \n", username, balance)
		case "2":
			fmt.Print("请输入需要存多少: ")
			var count float32
			_, err := fmt.Scan(&count)
			if err != nil {
				log.Println(err)
				continue
			}
			if err := Bank.Deposit(count); err != nil {
				log.Println(err)
				return
			}
			fmt.Println("存款成功!!!")
		case "3":
			fmt.Print("请输入需要取多少: ")
			var count float32
			_, err := fmt.Scan(&count)
			if err != nil {
				log.Println(err)
				continue
			}
			if err := Bank.Withdrawals(count); err != nil {
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
