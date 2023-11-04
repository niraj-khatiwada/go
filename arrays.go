package main

//
//import "fmt"
//
//func main() {
//	var array [2]string
//
//	array[0] = "cde"
//	array[1] = "cde"
//
//	for index, word := range array {
//		fmt.Println(index, word)
//	}
//
//	matrix := [2][2]int32{
//		{1, 2},
//		{2, 3},
//	}
//	fmt.Println(matrix[0:1])
//
//	var name string = "Niraj"
//	var character []rune = []rune(name)
//	for _, char := range character {
//		fmt.Println(char - 'a')
//		fmt.Println(string(char))
//	}
//	fmt.Println(string(character))
//
//	slice := make([]string, 6)
//	fmt.Println(slice, len(slice), slice[0] == "", slice == nil)
//
//	// Slice changes the original array as well
//	// In JavaScript, array.slice() will return the copy.
//	slice2 := array[:]
//	slice2[0] = "gif"
//	fmt.Println(array, slice2)
//
//	// Append wil create a new copy
//	appended := append(slice2, "lol")
//	fmt.Println(appended)
//	fmt.Println(array, slice2)
//
//}
