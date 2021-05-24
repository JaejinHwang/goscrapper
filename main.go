package main

import (
	"fmt"

	"github.com/JaejinHwang/goscrapper/accounts"
)

func main() {
	account := accounts.CreateAccount("Jay")
	fmt.Println(account)
}
