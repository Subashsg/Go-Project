package main

import "fmt"

func sorting(arr [6]int) {

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr)-1; j++ {

			if arr[j]%2 != 0 {

				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)

}

func main() {

	fmt.Println("sorting")

	fmt.Println("array to sort divided by 2 in first Indexs")

	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}

	sorting(arr)
}
