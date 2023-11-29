package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	balance float64
	lock    sync.Mutex
}

func main() {
	account := Account{balance: 200}
	go func() {
		if err := withdraw(100, &account); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		if err := withdraw(150, &account); err != nil {
			fmt.Println(err)
		}
	}()
	time.Sleep(1 * time.Second)
}

func withdraw(amount float64, account *Account) error {
	account.lock.Lock()
	if amount > account.balance {
		return errors.New("cannot withdraw amount greater than balance")
	}
	random := rand.Intn(10)
	time.Sleep(time.Duration(random) * time.Millisecond)

	account.balance = account.balance - amount
	defer account.lock.Unlock()
	fmt.Printf("withdraw of %f successful. Remaining amount is %f\n", amount, account.balance)
	return nil
}
