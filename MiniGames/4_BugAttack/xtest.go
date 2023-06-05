package main

import "fmt"

// xposWrite = append(xposWrite,(((-b.speed-1)/2 +rand.Intn(1+b.speed)))*int(zB))

func makeSlice(z int) []int {
	var slice []int
	for i:=int(0);i<=z;i++ {
		slice = append(slice,i)
	}
	return slice
}

func main() {
	var a [7]int= [7]int{0,1,2,3,4,5,6}
	for _,zahl := range a {
		slice := makeSlice(zahl)
		for _,zufall:= range slice {
			fmt.Println(zahl,zufall,-zahl/2+zufall)
		}
	}
}
