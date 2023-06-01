package main

import (
	"fmt"
	"time"
)

func cor(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("send ", i)
	}
	// when buffered channel is closed, receiver
	// can continue reading from channel until len(ch) == 0
	// i.e there is still some data in channel
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go cor(ch)
	// this allows cor to run
	// cor will write two values and block
	time.Sleep(500 * time.Millisecond)
	for v := range ch {
		// when one value is read from channel cor will unblock
		// and write one value
		fmt.Println("receive", v)
		time.Sleep(200 * time.Millisecond)
	}
}
