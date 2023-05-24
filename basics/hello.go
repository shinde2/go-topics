package main

import "fmt"

func print3(str string) string {
	return str + "!!"
}

func print2(str string) {
	fmt.Println(str)
}

func print1() {
	fmt.Println("Hello World!!")
}

func main() {
	//print1()
	//print2("Hello World!!")
	fmt.Println(print3("Hello World"))
}
