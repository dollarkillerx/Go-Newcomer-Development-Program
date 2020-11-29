package user

import "errors"

// 银行账户
type BankAccount struct {
	account  string // 私有
	password string
	balance  float32
}

func New(account, password string) *BankAccount {
	return &BankAccount{account: account, password: password}
}

func (b *BankAccount) CheckPassword(passwd string) bool {
	if passwd == b.password {
		return true
	}

	return false
}

func (b *BankAccount) Withdrawals(money float32) error {
	if b.balance < money {
		return errors.New("Insufficient balance")
	}

	b.balance -= money
	return nil
}

func (b *BankAccount) Deposit(money float32) {
	b.balance += money
}

func (b *BankAccount) BalanceEnquiry() float32 {
	return b.balance
}
