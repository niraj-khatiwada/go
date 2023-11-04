package main

//
//import (
//	"fmt"
//)
//
//func main() {
//	//fmt.Println(factorial(5))
//	//q, err := getQuotient(1, 0)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//fmt.Println(q)
//
//	//spread(1, 2, 3)
//	//array := []int{1, 2, 3}
//	//spread(array...)
//	//receiveArray(array)
//
//	//var value = 1
//	//fmt.Println(changeValue(&value))
//	//fmt.Println(value)
//
//	var array = [2]int{1, 2}
//	fmt.Println(changeArray(&array))
//	fmt.Println(array)
//
//	f4 := 100
//	var f4Pointer *int = &f4
//	*f4Pointer = 101
//	fmt.Println(f4)
//
//}
//
//func factorial(num int) int {
//	if num == 0 {
//		return 1
//	}
//	return num * factorial(num-1)
//}
//
//func getQuotient(a float64, b float64) (float64, error) {
//	if b == 0 {
//		return 0, fmt.Errorf("you can't divide by zero")
//	}
//	return b / a, nil
//}
//
//func spread(num ...int) {
//	fmt.Println(num)
//}
//
//func receiveArray(array []int) {
//	fmt.Println(array)
//
//}
//
//// Pass the pointer reference if you want to change the value
//func changeValue(val *int) int {
//	*val++
//	return *val
//}
//
//// Even passing the actual array, the original array is not changed unlike in JavaScript. You need to use pointer if you want to change the original array
////func changeArray(array [2]int) [2]int {
////	array[0] = 100
////	return array
////}
//
//func changeArray(array *[2]int) [2]int {
//	(*array)[0] = 100
//	return *array
//}
