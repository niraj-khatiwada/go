package main

import "fmt"

func main() {
	rest := getNames("Niraj", "Suraj")
	fmt.Println(rest)

	//	Spread
	getNames(rest...)
}
func getNames(names ...string) []string {
	fmt.Println("Rest Names", names)
	return names
}
