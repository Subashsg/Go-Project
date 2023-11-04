package main

import "fmt"

func checkPalindrome(num int) string {

	var input_num = num //1234321

	var rem int

	s := 0

	for num > 0 { //1234321 > 0

		rem = num % 10 //1234321 % 10
		s = (s * 10) + rem
		num = num / 10
	}

	if input_num == s {

		fmt.Println(s)
		return "the given no. is Palindrome"

	}
	fmt.Println(s)
	return "the given no. is not a Palindrome"

}

func main() {

	fmt.Println(checkPalindrome(1234321))

	fmt.Println(checkPalindrome(345123))

	fmt.Println(checkPalindrome(234371))
}
