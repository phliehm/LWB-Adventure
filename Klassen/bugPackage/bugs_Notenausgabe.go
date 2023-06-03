package bugPackage

import "fmt"

const maxPunkteProLevel uint32 = 5000
const anzahlLevel uint32 = 2

func BerechneNote() float32{
	n:= berechneProzent(berechneSummeVonSlice(punkteArray[:]),anzahlLevel*maxPunkteProLevel)
	fmt.Println("Prozent: ",n)
	if n>=0.90 {return 1.0}
	if n>=0.85 {return 1.3}
	if n>=0.80 {return 1.7}
	if n>=0.75 {return 2.0}
	if n>=0.70 {return 2.3}
	if n>=0.65 {return 2.7}
	if n>=0.60 {return 3.0}
	if n>=0.55 {return 3.3}
	if n>=0.50 {return 3.7}
	if n>=0.45 {return 4.0}
	if n>=0.40 {return 4.3}
	if n<0.40  {return 5.0}
	
	return 5.0
} 

// a ist x Prozent von b. Ausgabe in 0...1
func berechneProzent(a,b uint32) float32{
	return float32(a)/float32(b)
}


func berechneSummeVonSlice(s []uint32) uint32 {
	var sum uint32
	for _,z := range s {
		fmt.Println("Punkte: ",z)
		sum+=z
	}
	fmt.Println("Summe: ",sum)
	return sum
}

