package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "Prashant Pawar"
	str4 := "hello"
	str3 := str1 == str2
	str5 := str1 == str4
	fmt.Println(str3)
	fmt.Printf("%t", str5)
	// so we have to define type we have use %t sign

	concat := str1 + str2
	fmt.Printf("\n%s", concat)
}
