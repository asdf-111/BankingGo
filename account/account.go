package account

import "errors"

//Account Struct
type Account struct {
	owner   string
	balance int
}

//Create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit(+)
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance
func (a *Account) Balance() int {
	return a.balance
}

//Withdraw(-)
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw")
	}
	a.balance -= amount
	return nil
}
