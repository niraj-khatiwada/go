package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsan := map[string]string{"name": "Niraj", "age": "26"}

	stringified, _ := json.Marshal(jsan)

	fmt.Println("Stringified", stringified)

	var parsed map[string]string
	if err := json.Unmarshal(stringified, &parsed); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Parsed", parsed)

}
