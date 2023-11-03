package main

import "fmt"

func AddtwoArrs(arr1, arr2 [5]int) []int {

	resultArray := make([]int, len(arr1))

	for i := 0; i < len(arr1); i++ {

		resultArray[i] = arr1[i] + arr2[i]
	}
	return resultArray
}

func main() {

	arr1 := [5]int{5, 8, 13, 71, 6}

	arr2 := [5]int{4, 8, 83, 76, 9}

	resultArray := AddtwoArrs(arr1, arr2)

	fmt.Println("Result of adding the arrays:", resultArray)
}
