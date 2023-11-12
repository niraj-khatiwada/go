package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	time.Sleep(1001 * time.Millisecond) // Timeout
	//time.Sleep(1000 * time.Millisecond) // Fine

	if ctx.Err() == context.DeadlineExceeded {
		log.Fatal("Context timeout")
	} else {
		fmt.Println("Hello World")
	}
}
