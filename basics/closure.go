package main

import "fmt"

// Func can access vars in enclosing scope
// Func is bound to a variable ouside its local scope
func example2() {

	count := 0

	closed := func(str string) int {
		count += 1
		fmt.Println(str + "called closed")
		return count
	}

	fmt.Println(closed("Example2: "))
	fmt.Println(closed("Example2: "))

}

// Func and Anon functions can be assigned to variables
func example1() {

	var VarFunc1 func() = func() {
		fmt.Println("Example1: First call")
	}

	VarFunc1()

	VarFunc2 := func() {
		fmt.Println("Example1: Second call")
	}

	VarFunc1 = VarFunc2
	VarFunc1()

}

func main() {

	//example1()
	example2()

}
