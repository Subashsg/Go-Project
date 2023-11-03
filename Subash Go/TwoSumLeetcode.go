package main

import "fmt"

func Twosum(num [5]int, target int) []int {

	res := []int{}

	for i := 0; i < len(num); i++ {

		for j := 1; j < len(num); j++ {

			if num[i]+num[j] == target && i != j {
				res = []int{i, j}
				return res
			}
		}
	}
	return res
}

func main() {

	fmt.Println("Twosum Leetcode Problem")

	num := [5]int{2, 7, 8, 11, 15}

	target := 10

	res := Twosum(num, target)

	fmt.Println(res)
}
