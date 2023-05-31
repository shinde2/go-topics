package main

import "fmt"

/*
	Allows object like access.
	Functions with same name not allowed even if the signature is different.
	Methods with same name are allowed if receivers are different types.
	Mixing value and pointer receivers are allowed, GO will take care of dereferencing
*/

type Person struct {
	name string
	age  int
}

func (p *Person) example2(str string) {
	fmt.Println(str, p.age)
}

func (p Person) example1(str string) {
	fmt.Println(str, p.age)
}

func main() {

	p1 := Person{
		"Alice",
		10,
	}
	p1.example1("Alice")

	p2 := Person{
		"Bob",
		20,
	}
	p2.example2("Bob")

}
