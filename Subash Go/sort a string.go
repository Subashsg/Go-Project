package main

import (
	"fmt"
	"sort" //import fmt and sort package
)

// create main function to execute the program
func main() {
	unsorted_str := "ghfisaw" //create unsorted string
	fmt.Println("The unsorted string is:", unsorted_str)
	chars := []rune(unsorted_str)
	sort.Slice(chars, func(i, j int) bool { //sort the string using the function
		return chars[i] < chars[j]
	})
	fmt.Println("The sorted string is:")
	fmt.Println(string(chars)) //print the string on the console
}
