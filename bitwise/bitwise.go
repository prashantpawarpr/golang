package main

import "fmt"

func main() {
	a := 60 // 60 = 0011 1100
	b := 13 // 13 = 0000 1101
	c := a &^ b
	fmt.Printf("a = %d \n b= %d \n", a, b)
	fmt.Printf("c = %d \n", c)
}
