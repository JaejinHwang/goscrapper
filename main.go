package main

import (
	"fmt"

	"github.com/JaejinHwang/goscrapper/accounts"
)

func main() {
	account := accounts.CreateAccount("Jay")
	account.Deposit(1000)
	account.Deposit(3000)
	account.Withdraw(2000)
	err := account.Withdraw(5000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.CheckBalance())
	account.ChangeOwner("Hwang")
	fmt.Println(account.CheckOwner())
}