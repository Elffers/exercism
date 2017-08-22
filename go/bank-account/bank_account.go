package account

import  "sync"

const testVersion = 1

type Account struct {
	open    bool
	balance int64
	mtx sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{open: true, balance: initialDeposit}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	if !a.open {
		return
	}
	a.open = false

	return a.balance, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	if !a.open {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	if !a.open {
		return 0, false
	}


	newAmount := a.balance + amount

	if newAmount < 0 {
		return 0, false
	}
	a.balance = newAmount

	return a.balance, true
}
