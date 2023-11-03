package main

import "fmt"

func FibonacciSeries(num int) {
	a := 0
	b := 1
	sum := b

	fmt.Printf("Series is: %d %d ", a, b)

	for true {
		sum = b       //sum = b (1)    sum = 1      sum = 2        sum = 3
		b = a + b     //b = a+b (0+1)    b = 1 + 1     b = 1 + 2     b = 2 + 3
		if b >= num { //1 >= num (4)     2 >= 4        3 >= 4        5 >= 4

			fmt.Println()
			break
		}
		a = sum              //a = sum (1)    a = 1      a = 2
		fmt.Printf(" %d", b) //b = 1    b = 2      b = 3
	}
}

func main() {

	FibonacciSeries(16)
	FibonacciSeries(8)
	FibonacciSeries(16)
}
