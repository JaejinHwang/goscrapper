package accounts

//define struct Account
type Account struct {
	owner string
	balance int
}

//function CreateAccount makes new struct Account
func CreateAccount (owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}