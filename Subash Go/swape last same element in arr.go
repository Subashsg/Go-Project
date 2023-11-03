package main

import "fmt"

func Zero(arr []int, n int) {

	count := 0

	for i := 0; i < n; i++ {

		if arr[i] != 3 {

			arr[count] = arr[i]
			count++
		}

		for count < n {

			arr[count] = 3
			count++
		}
	}
}

func main() {

	arr := []int{3, 5, 3, 7, 2, 8, 1}
	n := len(arr)
	Zero(arr, n)
	fmt.Println(arr)
}
