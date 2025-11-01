package main

import "fmt"

func main() {
	arr := [6]int{10, 2, 3, 54, 4}
	for i, j := range arr {
		fmt.Println(i, j)
	}
}
