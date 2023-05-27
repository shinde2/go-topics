package main

import "fmt"

func example2_x(ptr *string) {

	*ptr = "henlo"

}

// new to get a pointer
// objects are garbage collected auto when nothing points to them
func example2() {

	ptr := new(string)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	example2_x(ptr)
	fmt.Println(*ptr)
}

func example1_x(ptr *int) {
	*ptr += 1
}

// * and & has same meaning as C++
func example1() {

	a := 5
	example1_x(&a)
	fmt.Println(a)

}

func main() {
	example1()
	example2()
}
