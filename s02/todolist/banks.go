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
		return fmt.Errorf("%s", "未登陆系统")
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.users[b.logUser].Withdrawals(money) {
		return fmt.Errorf("%s", "余额不足") // 抛出错误还可以这样写
	}

	return nil
}
