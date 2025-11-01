package main

// this is regarding understanding of the basic data types of the go lang

import "fmt"

var Exported = "This is an exported variable"

func main() {
	var a byte = 255
	var b byte = 244
	fmt.Printf("addition of two number %d + %d= %d \n", a, b, int32(a)+int32(b))
	fmt.Printf("addition of two number %d - %d= %d \n", a, b, a-b)
	fmt.Printf("addition of two number %d / %d= %d \n", a, b, a/b)
	fmt.Printf("addition of two number %d %% %d= %d \n", a, b, a%b)
}
