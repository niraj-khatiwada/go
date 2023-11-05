package main

//import (
//	"fmt"
//)
//
//func main() {
//	//go printTo(1000_000)
//	//go printTo(2000_000)
//	//time.Sleep(10 * time.Second)
//
//	channel1 := make(chan int)
//	go printToChannel(channel1, 2)
//	// This is just like calling .next() in generator function in JavaScript.
//	fmt.Println(<-channel1)
//	fmt.Println(<-channel1)
//}
//
//func printTo(target int) {
//	for i := 0; i < target; i++ {
//		fmt.Println(target, i)
//	}
//}
//
//// Consider channels to be like generators in javascript
//func printToChannel(channel chan int, target int) {
//	for i := 0; i < target; i++ {
//		channel <- i
//	}
//}
