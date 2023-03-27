// ADT bauelemente - Spezifikation und Implementation

// Martin Seiß    21.3.2023


package main

import b "./bauelemente"
import "fmt"
//import "gfx"




func main() {

// var be []Bauelement = make()
	
// Schalter 1
//	var ein1 b.Anschluss = b.Anschluss{0,false}
//	var ein2 b.Anschluss = b.Anschluss{0,false}
//	var aus b.Anschluss = b.Anschluss{1,false}
	var s1 b.Bauelement = b.New(0,100,100,false,false,false,b.Schalter)

	// Lampe 1
	var l1 b.Bauelement = b.New(1,0,0,false,false,false,b.Lampe)


	s1.SetzeEingang(2, true)	// Schalter an
	fmt.Println("Eingangswerte von Schalter 1", s1.GibEingang(1),s1.GibEingang(2))
	fmt.Println("Ausgangswert von Schalter 1", s1.BerechneAusgang())

	
	l1.SetzeEingang(1, true)
	fmt.Println("Eingangswerte von Leitung 1", l1.GibEingang(1),l1.GibEingang(2))
	fmt.Println("Ausgangswert von Leitung 1", l1.BerechneAusgang())



//	ZeichneBauelement()	
	
}
