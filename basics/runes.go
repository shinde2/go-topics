package main

import "fmt"

// IMP - https://go.dev/blog/strings
/*
	String is a Slice of bytes
	Each character is UTF-8 encoded
	Each character can have more than 1 byte
	0 - 127 is same as ASCII set
	Code point U+1236 is a hex value 12 36
	Rune is a int32 equivalent
*/

// chineese "hello"
const str = "你好"

func main() {

	fmt.Println("Plain string")
	fmt.Printf("%s", str)
	fmt.Println()

	fmt.Println("Hex representation")
	fmt.Printf("% x", str)
	fmt.Println()

	fmt.Println("Indexed for loop produces byte-by-byte values")
	for i := 0; i < len(str); i++ {
		fmt.Printf("% x", str[i])
	}
	fmt.Println()

	fmt.Println("Range based for loop produces rune-by-rune values")
	for i, r := range str {
		fmt.Printf("%#U at %d\n", r, i)
	}
	fmt.Println()

}
