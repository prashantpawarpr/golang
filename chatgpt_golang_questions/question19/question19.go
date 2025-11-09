package main

import "fmt"

func main() {
	var val1, val2 int

	_, err := fmt.Scan(&val1, &val2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1 . Addition")
		fmt.Println("2 . Subtraction")
		fmt.Println("3 . Division")
		fmt.Println("4 . Multiplication")
		fmt.Println("5 . Remainder")
		var choice int
		_, err = fmt.Scan(&choice)
		if err != nil {
			fmt.Println(err)
		} else {
			switch choice {
			case 1:
				fmt.Printf("%d + %d = %d \n", val1, val2, val1+val2)
			case 2:
				fmt.Printf("%d - %d = %d \n", val1, val2, val1-val2)
			case 3:
				fmt.Printf("%d / %d = %d \n", val1, val2, val1/val2)
			case 4:
				fmt.Printf("%d * %d = %d \n", val1, val2, val1*val2)
			case 5:
				fmt.Printf("%d %% %d = %d \n", val1, val2, val1%val2)

			}
		}

	}
}
