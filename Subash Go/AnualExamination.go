package main

import (
	"Examination"
	"fmt"
)

func main() {

	var arun Examination.Tamil

	arun.Tamilmark("Anual Examination", 90)

	var raja Examination.StudyTime

	raja.ComputerScience(4, 95)

	var vikcy Examination.StudyTime

	vikcy.Chemistry(1, 50)

	fmt.Println(raja, arun, vikcy)
}
