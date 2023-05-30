package main

import "fmt"

// https://go.dev/blog/slices-intro
// https://trstringer.com/golang-append/

// slice passing by value and reference
// https://trstringer.com/golang-slice-references/
/*
   slices will be passed by values with same array.
   new copy of slice will be created in the called func
   so len and cap of slice in calling func is not changed
*/

// when using append, if new values can be apended without
// increasing capacity, then same underlying array is used
// else new underlying array is created and new slice points to it
func example2() {
	x := make([]int, 3, 4)
	y := append(x, 10)
	fmt.Println(x)
	y[0] = 5
	fmt.Println(x, len(x), cap(x))
	fmt.Println(y, len(y), cap(y))

	u := make([]int, 3, 3)
	v := append(u, 10)
	fmt.Println(u)
	v[0] = 5
	fmt.Println(u, len(u), cap(u))
	fmt.Println(v, len(v), cap(v))

	fmt.Println("---------")

	// uses same underlying array
	a := make([]int, 3, 5)
	b := append(a[2:3], 10)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println("---------")

	// its possible that large underlying array has only one slice
	// referencing small part of it. Large array wont be GCed and
	// this is memory inefficient.
	// use copy to create new slice with new array
	c := make([]int, 5)
	d := make([]int, 2)
	copy(d, c[1:3])
	fmt.Println(c)
	d[0] = 5
	fmt.Println(c)
	fmt.Println(d)

}

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
	//example1()
	example2()
}
