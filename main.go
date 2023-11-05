package main

import (
	"fmt"
	"os"
)

func main() {
	// FPrint asks where to write this string
	fmt.Fprint(os.Stdout, "Hello World\n")
	// Println writes to Stdout
	fmt.Println("Hello World")
}
