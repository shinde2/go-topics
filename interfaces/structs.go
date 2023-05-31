package main

import "fmt"

// struct embedding

type Address struct {
	city string
}

type Person struct {
	name string
	age  int
	Address
}

func main() {

	p1 := Person{
		name: "Alice",
		age:  20,
		Address: Address{
			city: "LA",
		},
	}
	// city from embedded struct is promoted
	fmt.Println(p1.name, p1.Address.city, p1.city)

	// writing to promoted field does not work
	p2 := Person{
		name: "Bob",
		age:  30,
		city: "NY",
	}
	fmt.Println(p2.name, p2.city)

}
