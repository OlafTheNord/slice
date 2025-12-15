package srv

import "fmt"

type Account struct {
	Id      int
	Owner   string
	Balance int
}

type Deposit struct {
	Amount int
}

type Withdrawal struct {
	Amount int
	Fee    int
}

type Transaction interface {
	Apply(account *Account)
}

type Bank struct {
	Accounts map[int]*Account
}

func (d Deposit) Apply(a *Account) {
	a.Balance += d.Amount
}

func (w Withdrawal) Apply(a *Account) {
	a.Balance = a.Balance - (w.Amount + w.Fee)
}

func (b Bank) Execute(accountId int, tx Transaction) error {
	_, ok := b.Accounts[accountId]
	if !ok {
		fmt.Errorf("account with %d id not found", accountId)
	}
	tx.Apply(b.Accounts[accountId])
	return nil
}
