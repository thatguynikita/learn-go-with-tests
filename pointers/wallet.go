package main

import "fmt"

import "errors"

// ErrInsufficientFunds returns when there are not enough money in the wallet
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin is what you wants
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is our deposit acount
type Wallet struct {
	balance Bitcoin
}

// Deposit method allows to add money to the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance show actual balance in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw some money from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}