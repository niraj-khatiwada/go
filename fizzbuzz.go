package main

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//
//	fmt.Println("Enter a number")
//	reader := bufio.NewReader(os.Stdin)
//	value, err := reader.ReadString('\n')
//	if err != nil {
//		panic("Error reading input!")
//	}
//
//	var output = ""
//	val, err := strconv.Atoi(strings.TrimSpace(value))
//
//	if err != nil {
//		fmt.Println(err)
//		panic("Not a valid integer!")
//	}
//	if val%5 == 0 {
//		output += "fizz"
//	}
//	if val%3 == 0 {
//		output += "buzz"
//	}
//
//	if val%5 != 0 && val%3 != 0 {
//		output = strconv.Itoa(int(val))
//	}
//	fmt.Println(output)
//
//}
