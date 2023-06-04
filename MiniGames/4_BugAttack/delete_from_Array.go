package main

import "fmt"

var bugArray [5]*int

const zB, zH uint16 = 9,14			// Maße für die Zahlen 0 und 1 (Zellengröße)
var bug1Shape [21][2]uint16 = [21][2]uint16{{0,0},{6*zB,0},
												{1*zB,1*zH},{5*zB,1*zH},
												{2*zB,2*zH},{3*zB,2*zH},{4*zB,2*zH},
												{0,3*zH},{1*zB,3*zH},{2*zB,3*zH},{3*zB,3*zH},{4*zB,3*zH},{5*zB,3*zH},{6*zB,3*zH},
												{2*zB,4*zH},{3*zB,4*zH},{4*zB,4*zH},
												{1*zB,5*zH},{5*zB,5*zH},
												{0,6*zH},{6*zB,6*zH}}
func main() {
	var a int
	for i:=0;i<5;i++{
		bugArray[i] = &a
		fmt.Println(i,bugArray[i])
	}
	for i:=0;i<5;i++{
		bugArray[i] = nil
		fmt.Println(i,bugArray[i])
	}
	for i:=0;i<5;i++{
		bugArray[i] = &a
		fmt.Println(i,bugArray[i])
	}
	for index,b := range bugArray {
		fmt.Println(index,b)
		bugArray[index]=nil
	}
	for index,b := range bugArray {
		fmt.Println(index,b)
	}
	
	for k:= range bug1Shape {
		fmt.Println(bug1Shape[k][1])
	}
}
