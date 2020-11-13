# TODO LIST
我们在这里演示使用GO编写一个银行系统
功能：
- 用户登陆
- 存钱
- 取钱
## cli 内存版本
1. cli命令行方式的
2. 数据存储在内存中
#### 数据结构设计
```go
type User struct{
    UserName string 
    password string
    balance  float 
}
```
#### 为用户添加方法 
user.go
```go
package main

import "sync"

type User struct {
	UserName string
	password string
	balance  float32

	mu sync.Mutex // 互斥锁 保证线程操作数据安全性  (他会自动初始化)
}

func NewUser(username string, password string) *User {
	return &User{
		UserName: username,
		password: password,
	}
}

// 检验password
func (u *User) CheckPassword(password string) bool {
	u.mu.Lock()         // 上锁
	defer u.mu.Unlock() // defer 本方法生命周期结束会调用defer   我们在这里进行释放锁必然会释放 防止死锁

	if u.password == password {
		return true
	}
	return false
}

// 存款
func (u *User) Deposit(money float32) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.balance += money
}

// 取款
func (u *User) Withdrawals(money float32) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.balance < money {
		return false
	}

	u.balance -= money
	return true
}
```
我们写一个银行系统当然不能和只有一个用户啊，我们要设计存储多个用户的方案 我们这里用MAP来存储用户
```go 
type Banks struct {
	users map[string]*User  // 用户必须采用指针 这样才能保证我们修改和访问的数据都是同一块内存

	mu sync.Mutex
}
```
我们现在为Banks添加方法
banks.go
```go
package main

import (
	"errors"
	"fmt"
	"sync"
)

type Banks struct {
	// key username
	users map[string]*User // 用户必须采用指针 这样才能保证我们修改和访问的数据都是同一块内存

	logUser string // 如果用户登陆当前系统
	mu      sync.Mutex
}

var Bank *Banks // 全局函数 未初始化

func init() {
	// 初始化Bank
	Bank = NewBanks()
}

func NewBanks() *Banks {
	return &Banks{
		users: map[string]*User{},
	}
}

func (b *Banks) AddUser(username string, password string) error {
	if b.logUser == "" {
		return errors.New("未登陆系统")
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	_, ex := b.users[username] // 如果这样去取MAP data,ex := 会有两个返回 data为数据,ex为改数据是否存在
	if ex {                    // 如果存在 则返回false
		return errors.New("该用户以存在")
	}

	// 用户不存在就创建用户
	user := NewUser(username, password)
	b.users[username] = user
	return nil // 没有错误返回nil
}

func (b *Banks) Login(username, password string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	user, ex := b.users[username]
	if !ex { // golang的设计理念是 如果出现错误就立即死掉
		return errors.New("改用户不存在")
	}

	if !user.CheckPassword(password) {
		return errors.New("密码不正确")
	}

	b.logUser = username
	return nil
}

func (b *Banks) ObtainingBalance() (float32, error) {
	if b.logUser == "" {
		return 0, errors.New("未登陆系统")
	}

	b.mu.Lock()
	defer b.mu.Unlock()
	user := b.users[b.logUser] // 其实这步可以 如此 return b.users[b.logUser]..ObtainingBalance(), nil
	return user.ObtainingBalance(), nil
}

func (b *Banks) Deposit(money float32) error {
	if b.logUser == "" {
		return errors.New("未登陆系统")
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	b.users[b.logUser].Deposit(money)
	return nil
}

func (b *Banks) Withdrawals(money float32) error {
	if b.logUser == "" {
		return errors.New("未登陆系统")
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.users[b.logUser].Withdrawals(money) {
		return fmt.Errorf("%s", "余额不足") // 抛出错误还可以这样写
	}

	return nil
}
```
#### 然后开始写我们的CLI 把他们结合在一起
banks_cli.go
```go
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
```

### 如果改进这个服务 [todo list2]
- 缺陷
    - 代码没有分层
    - 我们现在存储写的太死 
改进项目:
1.定义存储层接口
定义Interface 
优势:
- 多种实现灵活迁移
- 行为约束
```go
type TodoListStorage interface {
	OpenAccount(user, password string) error                 // 开户
	BalanceEnquiry(user string) (balance float32, err error) // 查询余额
	Withdrawals(user string, money float32) error            // 取款
	Deposit(user string, money float32) error                // 存款
}
```



## WEB 版本
1. 我们先学习如何使用GO编写一个WEB HELLO WORLD

结合前端 配合http库实现一个web的TODO LIST
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