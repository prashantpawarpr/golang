package main

import "fmt"

// define global variable
var globalVar = "i am globla variabl"

func main() {
	localVar := "i am a local variabl"

	fmt.Printf("%s \n", globalVar)
	fmt.Printf("%s \n", localVar)
	second()
}

func second() {
	fmt.Printf("run from the secon function %s \n", globalVar)
	fmt.Printf("%s \n", globalVar)
}
