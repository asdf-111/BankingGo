package main

import (
	"fmt"

	"BankingGo.com/account"
)

func main() {
	account := account.NewAccount("Happy")
	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
