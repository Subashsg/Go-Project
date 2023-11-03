package main

import "fmt"

func checkArmstrong(num int) bool {

	temp := 0
	reminder := 0
	ams := 0

	temp = num

	for {

		reminder = temp % 10
		ams += reminder * reminder * reminder
		temp = temp / 10
		if temp == 0 {
			break
		}
	}

	if ams == num {

		fmt.Printf(" %d is a Armstrong number \n", num)
	} else {
		fmt.Printf(" %d is not a Armstrong number \n", num)
	}

	return true

}

func main() {

	checkArmstrong(153)

	checkArmstrong(275)

	checkArmstrong(371)
}
