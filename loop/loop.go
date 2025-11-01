package main

import "fmt"

func main() {
	for a := 1; a < 5; a++ {
		fmt.Println("for loop", a)
	}

	b := 1
	for b < 5 {
		fmt.Println("while loop", b)
		b++

	}

	c := 1
	for {
		fmt.Println("infinite loop", c)
		if c > 5 {
			break
		}
		c++
	}

	// third important way of loop
	// for i,j := range rvariable{
	//
	// statement
	//
	// }
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	for i, j := range arr {
		fmt.Println(i, j)
	}
}
