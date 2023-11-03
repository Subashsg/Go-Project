package main

import "fmt"

func reverse(s string) string {

	rns := []rune(s)

	for i, j := len(rns)-3, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func main() {

	str := "rajram"
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
