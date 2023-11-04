package main

import "fmt"

func isLeapYear(year int) bool {

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	} else {
		return false
	}
}

func main() {

	year := 2024

	if isLeapYear(year) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
