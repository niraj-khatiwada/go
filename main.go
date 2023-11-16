package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	balance int
	lock    sync.Mutex
}

var account = Account{balance: 200}

func main() {
	go (func() {
		if err := withdraw(&account, 100); err != nil {
			fmt.Println(err)
		}
	})()
	go (func() {
		if err := withdraw(&account, 150); err != nil {
			fmt.Println(err)
		}
	})()
	time.Sleep(1 * time.Second)
}

func withdraw(account *Account, amount int) error {
	account.lock.Lock()
	if amount > account.balance {
		return fmt.Errorf("[error] withdraw amount greater than the balance")
	}
	random := rand.Intn(10)
	time.Sleep(time.Duration(random) * time.Millisecond)
	account.balance -= amount
	account.lock.Unlock()
	fmt.Printf("\n%d withdrawn. Remaining amount = %d\n", amount, account.balance)
	return nil
}
