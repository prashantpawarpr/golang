package main

import "fmt"

var Exported = "This is an exported variable"

func main() {
	var b byte = 255
	var be int16 = 32000
	fmt.Println(b)
	fmt.Printf("/n   %d", be)
	a := 1
	if a != 10 {
		fmt.Printf("the number is %d\n", a)
	}
}
