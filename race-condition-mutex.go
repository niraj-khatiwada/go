package main

//import (
//	"errors"
//	"fmt"
//	"log"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//type Account struct {
//	balance int
//	lock    sync.Mutex
//}
//
//func main() {
//
//	var account Account = Account{balance: 200}
//
//	go func() {
//		if err := withdraw(&account, 100); err != nil {
//			log.Fatal(err)
//		}
//
//	}()
//	go func() {
//		if err := withdraw(&account, 150); err != nil {
//			log.Fatal(err)
//		}
//
//	}()
//	time.Sleep(2 * time.Second)
//}
//
//func withdraw(account *Account, amount int) error {
//	account.lock.Lock()
//	if amount > account.balance {
//		return errors.New("cannot withdraw amount more than the balance")
//	}
//	random := rand.Intn(10)
//	time.Sleep(time.Duration(random) * time.Millisecond)
//	account.balance -= amount
//	fmt.Println("Withdraw successful. Current balance is", account.balance)
//	defer account.lock.Unlock()
//	return nil
//}
