package main

import (
	srv "balance/src"
)

func main() {
	//create accounts
	acc1 := srv.Account{
		Id:      1,
		Owner:   "Phil",
		Balance: 10000,
	}
	acc2 := srv.Account{
		Id:      2,
		Owner:   "John",
		Balance: 70000,
	}
	acc3 := srv.Account{
		Id:      3,
		Owner:   "Mary",
		Balance: 8000,
	}

	//create bank
	alfabank := srv.Bank{
		Accounts: map[int]*srv.Account{acc1.Id: &acc1, acc2.Id: &acc2, acc3.Id: &acc3},
	}
	//create transactions
	dep1 := srv.Deposit{
		Amount: 1000,
	}
	dep2 := srv.Deposit{
		Amount: 2000,
	}
	dep3 := srv.Withdrawal{
		Amount: 1000,
		Fee:    500,
	}

	alfabank.Execute(1, dep1)
	alfabank.Execute(2, dep2)
	alfabank.Execute(3, dep3)
}
