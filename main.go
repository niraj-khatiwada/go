package main

import (
	"errors"
	"fmt"
)

var targetError = errors.New("error not found")

func main() {
	err := throwsError()
	fmt.Println(errors.Is(err, targetError))

	hashMap := map[string]int{"niraj": 2}

	fmt.Println(len(hashMap))
}

func throwsError() error {
	return fmt.Errorf("error occurred: %w", targetError)
}
