package main

//import (
//	"encoding/json"
//	"fmt"
//	"log"
//)
//
//type Person struct {
//	name string
//	age  int
//}
//
//func main() {
//	person := map[string]string{"name": "Niraj"}
//
//	// JSON.stringify()
//	bt, _ := json.Marshal(person)
//	fmt.Println("stringified", string(bt[:]))
//
//	// JSON.parse()
//	var value map[string]string
//	if err := json.Unmarshal(bt[:], &value); err != nil {
//		log.Fatal("[error] error parsing json")
//	}
//	fmt.Println("parsed", value)
//
//}

//package main
//
//import (
//"encoding/json"
//"fmt"
//)
//
//func main() {
//	jsan := map[string]string{"name": "Niraj", "age": "26"}
//
//	stringified, _ := json.Marshal(jsan)
//
//	fmt.Println("Stringified", stringified)
//
//	var parsed map[string]string
//	if err := json.Unmarshal(stringified, &parsed); err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Parsed", parsed)
//
//}
