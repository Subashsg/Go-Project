package main

import (
	"fmt"
	"functionformula"
)

func Allpolicy(policylife int, policyname string, premium int) {

	var youget string
	var bonus string

	if policyname == "Guaranteed Income Plan" && policylife == 15 {

		functionformula.Logic(policylife, premium)
		youget = "$1.22"
		bonus = "$1.74"

	} else if policyname == "Guaranteed Lumpus Plan" && policylife == 15 {

		youget = "$30 L"

	}

	fmt.Println(youget, policyname, bonus)
}

func Suggestpolicy(age int, salary int, premium int) {

	var policylife int = 15
	var policyname string = "Guaranteed Income Plan"

	if age >= 25 && premium >= 6000 {

		Allpolicy(policylife, policyname, premium)
	}
}

func Lic() {

	var age int = 25
	var salary int = 25000
	var premium int = 6000
	var ret bool
	ret = functionformula.Checkeligibility(age, salary)

	if ret == true {

		Suggestpolicy(age, salary, premium)
	}

}

func main() {

	Lic()
	fmt.Println("lic policy lifetime")
}
