package main

import "fmt"

func main() {
	var a int32 = 56
	var b float32 = 1321.13

	fmt.Println("hello brother")
	fmt.Printf("%d", a)
	fmt.Printf("%f", b)
	c := int32(a) + int32(b)
	fmt.Println("the value of c is:", c)
}
