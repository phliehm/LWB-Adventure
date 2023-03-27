// ADT bauelemente - Spezifikation und Implementation

// Martin Seiß    21.3.2023


package main

import sch "./schaltungen"
import b "./bauelemente"
import "fmt"
//import "gfx"



func main() {

	var sk sch.Schaltung = sch.New()
	
	sk.BauteilEinfuegen(1,100,100,b.Schalter)
	sk.BauteilEinfuegen(2,300,200,b.Lampe)

	sk.VerbindungEinfuegen(0,1,2,200)

		
	//sk.GibSchalterwert(id uint16) bool
	//sk.GibLampenStatus(id uint16) []bool


	fmt.Println("Zustand von Schalter 1", sk.GibSchalterwert(1))
	fmt.Println("Ausgangswert von Lampe 1", sk.GibLampenwert(2))
	fmt.Println("Ausgangswert aller Lampen", sk.GibLampenStatus())
	fmt.Println()

	fmt.Println("Schalte 1 an:")
	sk.SchalteLampeAn(1,true)

	// sberechneSchaltkreis(be,v)
	
	fmt.Println("Zustand von Schalter 1", sk.GibSchalterwert(1))
	fmt.Println("Ausgangswert von Lampe 1", sk.GibLampenwert(2))
	fmt.Println("Ausgangswert aller Lampen", sk.GibLampenStatus())
	fmt.Println()

//	ZeichneBauelement()	
	
}
