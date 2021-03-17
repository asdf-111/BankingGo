package main

import (
	"fmt"

	"BankingGo.com/banking"
)

func main() {
	account := banking.Account{Owner: "Happy", Balance: 1000}
	fmt.Println(account)
}
