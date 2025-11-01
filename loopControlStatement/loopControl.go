package main

import "fmt"

func main() {
	for i := range 10 {
		fmt.Println("Value of i is:", i)

	label1:
		for j := range 5 {
			fmt.Println("Value of j is:", j)
			break label1
		}

		if i == 5 {

			fmt.Println("going to the loop label 1", i)
			goto label1

		}

	}
}

// it will be forming the infinite loop because when i reaches 5 it will goto label1 and start the inner loop again and again
// culprit is lable1 and goto statement and the statement is not breaking
