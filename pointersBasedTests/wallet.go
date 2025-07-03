package pointersbasedtests

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	if amount < 0 {
		return
	}
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrBalanceStatement = errors.New("insufficient Balance")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount < 0 {
		return ErrBalanceStatement
	}
	if w.balance < amount {
		return ErrBalanceStatement
	}
	w.balance -= amount
	return nil
}
