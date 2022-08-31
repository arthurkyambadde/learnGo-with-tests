package wallet

import (
	"errors"
	"fmt"
)

var ErrinsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposite(amount Bitcoin) {

	fmt.Printf("Address for the deposite is %v\n", &w.balance)
	w.balance += amount

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrinsufficientFunds
	}

	w.balance = w.balance - amount

	return nil

}
