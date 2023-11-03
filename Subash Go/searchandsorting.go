package main

import "fmt"

func sorting(arr [5]int) {

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr)-i-1; j++ {

			if arr[j] > arr[j+1] {

				arr[j], arr[j+1] = arr[j+1], arr[j]

			}

		}
	}
	fmt.Println("Rerange to sorting data ", arr)

}

func searching(arr [5]int, lucky int) {

	for _, v := range arr {

		if v == lucky {

			fmt.Println("found ", lucky)
		}
	}

}

func main() {

	fmt.Println("searching & sorting")

	var arr [5]int = [5]int{8, 2, 3, 1, 6}

	searching(arr, 8)

	sorting(arr)
}
