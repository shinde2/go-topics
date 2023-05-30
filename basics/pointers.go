package main

import "fmt"

// pointer arithmetic not supported
func example6() {

	arr := [2]int{0, 1}
	ptr := &arr
	fmt.Println((*ptr)[0])
	ptr++
	fmt.Println((*ptr)[0])

}

// Can pass pointer to array or pass array as a slice
func example5_x(arr []int) {
	arr[1] = 2
}

func example5() {

	arr := [2]int{0, 1}
	example5_x(arr[:])
	fmt.Println(arr)

}

// Everyting is pass by values inclding arrays and slices
func example4_x(arr [2]int) {
	arr[1] = 2
}

func example4() {

	arr := [2]int{0, 1}
	example4_x(arr)
	fmt.Println(arr)

}

// Go compiler will perform escape analysis to allocate
// local variable ptr to heap when its returned from func
func example3() {
	var ptr *int
	fmt.Println(ptr == nil)
	// can not dereference null pointer
	//fmp.Println(*ptr)

	// newPtr := new(int)
	var newPtr *int = new(int)
	fmt.Println(newPtr == nil)
	// new will create zero initialized pointer
	fmt.Println(*newPtr)

}

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
	//example1()
	//example2()
	//example3()
	//example4()
	//example5()
	example6()
}
