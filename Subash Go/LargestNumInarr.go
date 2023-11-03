package main

import (
	"fmt"
)

func FindlargestNum(arr [5]int) int {

	max := arr[0]

	for _, value := range arr {

		if value > max {
			max = value
		}
	}
	return max
}

func main() {

	arr := [5]int{4, 8, 83, 76, 9}

	largest := FindlargestNum(arr)

	fmt.Printf("the largest Element in a array is : %d\n", largest)
}
