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

	// Code about commit #1.4 ~ #1.5

	dictionary := dictionarys.Dictionary{}
	exword := "hello"
	dictionary.Add(exword, "greeting")
	result := dictionary.Update("xdd", "first")
	if result != nil {
		fmt.Println(result)
	}
	word, _ := dictionary.Search(exword)
	fmt.Println(word)
	fmt.Println(dictionary)
	result2 := dictionary.Delete("second")
	if result2 != nil {
		fmt.Println(result2)
	}
	fmt.Println(dictionary)
	}