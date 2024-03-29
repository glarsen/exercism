package account

import (
	"sync"
)

// Account implements a bank account.
type Account struct {
	active  bool
	balance int64
	mu      sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{active: true, balance: amount}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if !a.active {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.active || a.balance+amount < 0 {
		return 0, false
	}

	a.balance += amount

	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.active {
		return 0, false
	}

	a.active = false

	return a.balance, true
}
