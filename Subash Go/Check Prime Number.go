package main

import (
	"fmt"
	"math"
)

func CheckPrimeNo(num int) {

	if num < 2 {
		fmt.Println("the prime no. must greater than 2")
		return
	}

	squ_root := int(math.Sqrt(float64(num)))

	for i := 2; i <= squ_root; i++ {

		if num%i == 0 {

			fmt.Println(num, "is a non prime number")
			return
		}
	}
	fmt.Println(num, "is a prime number")
	return
}

func main() {

	CheckPrimeNo(3)

	CheckPrimeNo(6)

	CheckPrimeNo(11)
}
