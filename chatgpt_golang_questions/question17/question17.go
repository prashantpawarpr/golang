package main

import "fmt"

func main() {
	var val int
	count := 0

	_, err := fmt.Scan(&val)
	if err != nil {
		fmt.Println("incorrect input")
	} else {
		for i := 1; i <= val; i++ {
			// fmt.Printf("%d \n", i)
			if val%i == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Println("Prime number")
		} else {
			fmt.Println("Not Prime number")
		}
	}
}
