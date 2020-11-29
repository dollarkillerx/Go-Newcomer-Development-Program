package storage

type TodoListStorage interface {
	OpenAccount(account, password string) error // 开户
	CheckAccount(account, password string) error  // check account
	BalanceEnquiry(account string) (balance float32, err error) // 查询余额
	Withdrawals(account string, money float32) error            // 取款
	Deposit(account string, money float32) error                // 存款
}
