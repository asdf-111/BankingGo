package main

import (
	"fmt"
)

func main() {
	account := banking.Account{Owner: "Happy", Balance: 1000}
	fmt.Println(account)
}
