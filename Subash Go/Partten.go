package main

import "fmt"

func hallow(side int) {

	for i := 0; i < side; i++ {

		for j := 0; j < side; j++ {

			if i == 0 || i == side-1 || j == 0 || j == side-1 {

				fmt.Print("*")
			} else {
				fmt.Print()
			}
		}
		fmt.Println()
	}
}

func main() {

	hallow(5)
}
