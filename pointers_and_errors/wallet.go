package main

import (
	"fmt"
	"errors"
)

//type MyName OriginalType(like a typedef) of a type=>int
//can be used as a conversion function: wallet := Balance(Bitcoin(20))
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// interface defined in the fmt package and lets you define 
// how your type is printed
type Stringer interface {
	String() string
}

//in Go errors are values. So we refactor to have 1 source
//var kw allows us to define values global to the package
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")


// we can access the internal balance using a "receiver" var
// the value receiver modifies the copy of the type Wallet
// the pointer modifies the wallet itself so same addrr
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// syntax for creating method on type decl same as struct
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		//creates new err with message of choosing
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
