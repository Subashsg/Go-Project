package main

import "fmt"

func factorial(num int) int {

	if num == 0 || num == 1 { // 1)num = 5  2)num = 4     3)num = 3     4)num = 2    5)num = 1

		return num
	} else {

		return num * factorial(num-1) // 1)num(5) * f(5-1) = 5*4   2) num(4) * f(4-1) = 4*3   3) num(3) * f(3-1) = 3*2  4) num(2) * f(2-1) = 2*1  5) num(1)
	}

}
func main() {

	fmt.Println(factorial(5))

	fmt.Println(factorial(3))

	fmt.Println(factorial(2))
}
