package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPAgus NoKTP = "432432432432432432432"
	var marriedStatus Married = true
	fmt.Println(noKTPAgus)
	fmt.Println(marriedStatus)
}