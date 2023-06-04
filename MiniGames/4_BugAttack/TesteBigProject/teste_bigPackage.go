package main

import (
		"./bigPackage"
		"fmt"
		)


var zweiZ bigPackage.ZweiZahlen


func main() {
	
	zweiZ.A = 5
	zweiZ.B = 10
	fmt.Println(bigPackage.Addiere2Zahlen(zweiZ.A,zweiZ.B))
	fmt.Println(bigPackage.Multipliziere2Zahlen(2,5))
	fmt.Println(bigPackage.GibSummeMitStruct(zweiZ))
}
