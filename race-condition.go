package main

//import (
//	"errors"
//	"fmt"
//	"log"
//	"math/rand"
//	"time"
//)
//
//var balance = 200
//
//func main() {
//
//	go func() {
//		if err := withdraw(100); err != nil {
//			log.Fatal(err)
//		}
//
//	}()
//	go func() {
//		if err := withdraw(150); err != nil {
//			log.Fatal(err)
//		}
//
//	}()
//	time.Sleep(2 * time.Second)
//}
//
//func withdraw(amount int) error {
//	if amount > balance {
//		return errors.New("cannot withdraw amount more than the balance")
//	}
//	random := rand.Intn(10)
//	time.Sleep(time.Duration(random) * time.Millisecond)
//	balance -= amount
//	fmt.Println("Withdraw successful. Current balance is", balance)
//	return nil
//}
