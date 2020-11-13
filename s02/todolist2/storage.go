package todolist2

type TodoListStorage interface {
	OpenAccount(user, password string) error                 // 开户
	BalanceEnquiry(user string) (balance float32, err error) // 查询余额
	Withdrawals(user string, money float32) error            // 取款
	Deposit(user string, money float32) error                // 存款
}
