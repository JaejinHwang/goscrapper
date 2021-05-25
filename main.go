package main

import (
	"fmt"

	"github.com/JaejinHwang/goscrapper/dictionarys"
)

func main() {
	// Code about commit #1.1 ~ #1.3

	// account := accounts.CreateAccount("Jay")
	// account.Deposit(1000)
	// account.Deposit(3000)
	// account.Withdraw(2000)
	// err := account.Withdraw(5000)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.CheckBalance())
	// account.ChangeOwner("Hwang")
	// fmt.Println(account.CheckOwner())

	// Code about commit #1.4

	dictionary := dictionarys.Dictionary{"first": "first word"}
	definition, error := dictionary.Search("first")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}

}