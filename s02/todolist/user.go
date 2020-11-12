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

// 获取余额
func (u *User) ObtainingBalance() float32 {
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.balance
}
