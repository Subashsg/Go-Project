package main

import "fmt"

func main() {

	var a int = 17

	fmt.Println("enter a number....")

	if a%2 == 0 {

		fmt.Printf("%d The Given Num is Even ", a)

	} else {

		fmt.Printf("%d The Given Num is Odd ", a)
	}
}
