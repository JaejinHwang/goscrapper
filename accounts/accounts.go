package accounts

import (
	"errors"
	"fmt"
)

//define struct Account
type Account struct {
	owner string
	balance int
}

var errNomoney = errors.New("amount of withdraw is bigger than your balance")

//function CreateAccount makes new struct Account
func CreateAccount (owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//method Deposit makes input amount of deposit into account
func (a *Account) Deposit (amount int) {
	a.balance += amount
}

//method CheckBalance checks balance of account
func (a Account) CheckBalance () int {
	return a.balance
}

//method Withdraw makes input amount of withdraw into account
func (a *Account) Withdraw (amount int) error {
	if a.balance < amount {
		return errNomoney
	}
	a.balance -= amount
	return nil
}

//method ChangeOwner changes owner of account
func (a *Account) ChangeOwner (newOwner string) {
	a.owner = newOwner
}

//method CheckOwner Checks owner of account
func (a Account) CheckOwner () string {
	return a.owner
}

func (a Account) String () string {
	return fmt.Sprint(a.CheckOwner(), "'s account has", a.CheckBalance())
}