package main

import (
	"fmt"
	"time"
)

func co1() {
	for i := 0; i < 5; i++ {
		fmt.Println("In co1")
		time.Sleep(100 * time.Millisecond)
	}
}

func co2() {
	for i := 0; i < 5; i++ {
		fmt.Println("In co2")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {

	go co1()
	go co2()
	// without sleep here, main coroutine will exit immediately
	// and that will kill other coroutines started by this one.
	time.Sleep(500 * time.Millisecond)

}
