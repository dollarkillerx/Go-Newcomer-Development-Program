package storage

import (
	"github.com/dollarkillerx/Go-Newcomer-Development-Program/s02/todolist_pro/user"

	"errors"
	"sync"
)

// 对TodoListStorage内存的实现实现, 你也可以用MySQL实现  只要实现TodoListStorage就是直接换上去用
type Simple struct {
	mu sync.Locker
	db map[string]*user.BankAccount
}

func NewSimple() TodoListStorage {
	return &Simple{
		db: map[string]*user.BankAccount{},
	}
}

// 开户
func (s *Simple) OpenAccount(account, password string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.db[account] = user.New(account, password)

	return nil
}

// check account
func (s *Simple) CheckAccount(account, password string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	act, ex := s.db[account]
	if !ex {
		return errors.New("not account")
	}

	if act.CheckPassword(password) {
		return nil
	}

	return errors.New("password error")
}

// 查询余额
func (s *Simple) BalanceEnquiry(account string) (balance float32, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	act, ex := s.db[account]
	if !ex {
		return 0, errors.New("not account")
	}

	return act.BalanceEnquiry(), nil
}

// 取款
func (s *Simple) Withdrawals(account string, money float32) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	act, ex := s.db[account]
	if !ex {
		return errors.New("not account")
	}

	return act.Withdrawals(money)
}

// 存款
func (s *Simple) Deposit(account string, money float32) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	act, ex := s.db[account]
	if !ex {
		return errors.New("not account")
	}

	act.Deposit(money)
	return nil
}
