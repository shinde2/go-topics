package main

import "fmt"

/*
- used to execute a piece of code just before
  current func or method returns
- executed in LIFO manner
- varables have values when func is pushed on
  stack i.e. when defer is defined
*/

func example(a int) {
	fmt.Println(a)
}

func main() {
	a := 0
	for i := 0; i < 5; i++ {
		a++
		defer example(a)
	}

	// all deferred funcs are executed now

}
