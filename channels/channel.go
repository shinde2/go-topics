package main

import (
	"fmt"
	"time"
)

func co2(ch2 chan bool) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("ch2 coroutine done")
	// writing a value to channel
	// its blocking
	// i.e. this coroutine will block until another coroutine,
	// is ready to read from this channel
	ch2 <- true
}

func co1(ch1 chan bool) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("ch1 coroutine done")
	ch1 <- true
}

func main() {
	// creating bool type channel with 0 capacity i.e. blocking
	// its a bidirectional channel
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go co1(ch1)
	go co2(ch2)
	// reading from channel
	// its blocking
	// i.e. this coroutine will block until another coroutine,
	// is ready to write to this channel
	x := <-ch2
	fmt.Println(x)
	// channel value can be ignored
	<-ch1
	fmt.Println("Main coroutine done")
}
