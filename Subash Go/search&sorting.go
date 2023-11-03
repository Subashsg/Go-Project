package main

import (
	"fmt"
)

func main() {

	fmt.Println("searching & sorting")

	var arr [5]int = [5]int{8, 2, 3, 1, 6}

	var ram searchandsort.searching
	ram.Searching(arr, 8)

	var raja searchandsort.sorting
	raja.Sorting(arr)

	fmt.Println(arr, arr)
}
