package main

import "fmt"

func main() {
	var a int32 = 56
	var b float32 = 1321.13

	var str rune = 97

	fmt.Printf("%s \n", string(str))

	fmt.Printf("%d", a)
	fmt.Printf("%f", b)
	c := int32(a) + int32(b)
	fmt.Println("the value of c is:", c)

	comp := 3 + 22i

	reaNum := real(comp)
	imaNum := imag(comp)
	fmt.Println(reaNum, imaNum)
}
