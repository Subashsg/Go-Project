package main

import "fmt"

func findFactors(n int) []int {

	factors := []int{}

	for i := 1; i <= n; i++ {

		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func main() {

	number := 12

	factors := findFactors(number)

	fmt.Printf("The factors of %d are: %v\n", number, factors)

}
