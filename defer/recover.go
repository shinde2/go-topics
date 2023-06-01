package main

import "fmt"

/*
- defer can be used to recover from panic
- recovery is possible if deferred function is present
  in same stack of current goroutine i.e. if divide() spawns
  new goroutine which does not have its own defer, it will
  not be recovered by catchRecover()
*/

func catchRecover() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func divide(a int, b int) {
	defer catchRecover()
	d := a / b
	fmt.Printf("division is %d\n", d)
}

func main() {
	divide(5, 1)
	divide(5, 0)
	fmt.Println("main finish")
}
