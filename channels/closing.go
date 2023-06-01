package main

import "fmt"

func cor(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	// sender coroutine can close the channel
	close(ch)
}

func main() {
	ch := make(chan int)
	go cor(ch)

	// receiver coroutine can close the channel
	//close(ch)

	// expicit check for closed channel
	for {
		// when more is false, sender is done writing to the channel
		val, more := <-ch
		if more == false {
			break
		}
		fmt.Println(val)
	}

	// can not close closed channel
	//close(ch)
}
