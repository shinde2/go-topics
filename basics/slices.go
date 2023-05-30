package main

import "fmt"

// https://go.dev/blog/slices-intro

func example1() {

	// len = 3, cap = 3
	x := []int{0, 0, 0}
	fmt.Println(x)

	// len = 3, cap = 3
	y := make([]int, 3)
	fmt.Println(y)
	fmt.Println(y[0] == 0)

	// len = 3, cap = 5
	z := make([]int, 3, 5)
	fmt.Println(z)

	// not allowed
	//w := make([]int, 5, 3)
	//fmt.Println(w)

	// zero value is nil
	var u []int
	fmt.Println(u == nil)
	// but this is not nil ???
	v := []int{}
	fmt.Println(v == nil)

}

func main() {
	example1()
}
