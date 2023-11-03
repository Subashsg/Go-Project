package main

import (
	"fmt"
)

func Findsmallest(arr [5]int) int {

	min := arr[0]

	for _, value := range arr {

		if value < min {
			min = value
		}
	}
	return min
}

func main() {

	arr := [5]int{5, 8, 13, 71, 6}

	smallest := Findsmallest(arr)

	fmt.Printf("the Smallest Element In a Array is : %d\n", smallest)
}
