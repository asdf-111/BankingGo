package main

import (
	"fmt"

	"./banking"
)

func main() {
	account := banking.Account{Owner: "Happy", Balance: 1000}
	fmt.Println(account)
}
