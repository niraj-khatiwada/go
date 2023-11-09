package main

import (
	"errors"
	"fmt"
)

var targetError = errors.New("error not found")

func main() {
	err := throwsError()
	fmt.Println(errors.Is(err, targetError))
}

func throwsError() error {
	return fmt.Errorf("error occurred: %w", targetError)
}
