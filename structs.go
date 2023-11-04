package main

//import (
//	"fmt"
//	"reflect"
//)
//
//type Type1 interface {
//	int | string | float64
//}
//
//type Person struct {
//	name string
//	age  int
//}

//type PersonExtended struct {
//	id     string
//	person Person
//}

//
//func (p Person) PrintPerson() {
//	fmt.Printf("Hello %s, you are %d years old.", p.name, p.age)
//}
//

//type Int int
//
//func (i Int) toString() string {
//	return fmt.Sprintf("%d", i)
//}
//func main() {
//	fmt.Println(sum(1, 2))
//	fmt.Println(sum("1", "2"))
//	fmt.Println(sum(1.2, 2.2))
//
//	var person Person = Person{name: "Niraj", age: 100}
//	fmt.Println(person, reflect.TypeOf(person))

//person2 := Person{
//name: "Niraj",
//age:  26,
//}
//
//extended := PersonExtended{
//id:     "100",
//person: person2,
//}
//
//	changePerson(&person)
//	fmt.Println(person)
//	person.PrintPerson()
//
//var i Int = 100
//fmt.Println(i.toString())
//}
//
//func sum[T Type1](a T, b T) T {
//	return a + b
//}
//
//func changePerson(person *Person) {
//	(*person).name = "Suraj"
//	fmt.Println(*person)
//}
