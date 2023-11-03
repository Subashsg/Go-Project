package main

import "fmt"

func removeDuplicate(arr [8]int) []int {

	map_var := map[int]bool{}

	result := []int{}

	for i := range arr {

		if map_var[arr[i]] != true {

			map_var[arr[i]] = true

			result = append(result, arr[i])
		}
	}
	return result
}

func main() {

	arr := [8]int{1, 2, 2, 4, 4, 5, 7, 5}

	fmt.Println("The unsorted array entered is:", arr)

	result := removeDuplicate(arr)

	fmt.Println("The array obtained after removing the duplicate values is:", result)
}
