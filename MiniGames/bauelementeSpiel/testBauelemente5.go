// ADT bauelemente - Spezifikation und Implementation

// Martin Seiß    21.3.2023


package main

import sch "./schaltungen"
import b "./bauelemente"
import "fmt"
import "gfx"



func main() {

	// ---------------- Erzeuge Schaltkreis -------------- //

	var sk sch.Schaltung = sch.New()
	
	sk.BauteilEinfuegen(1,100,100,b.Schalter)
	sk.BauteilEinfuegen(2,100,400,b.Schalter)
	
	sk.BauteilEinfuegen(3,400,250,b.AND)	
	
	sk.BauteilEinfuegen(4,700,100,b.Lampe)
	sk.BauteilEinfuegen(5,700,400,b.Lampe)

	sk.VerbindungEinfuegen(1,3,1,250)
	sk.VerbindungEinfuegen(1,4,1,550)
	sk.VerbindungEinfuegen(2,3,2,250)
	sk.VerbindungEinfuegen(3,5,1,550)

	sk.SchaltungBerechnen()


	// ---------------- Zeichne Schaltkreis -------------- //

//	gfx.Fenster(1200,700)
	gfx.Fenster(800,500)
		
	// ---------------- Teste Schaltkreis -------------- //

	fmt.Println("Zustand von Schalter 1", sk.GibSchalterwert(1))
	fmt.Println("Zustand von Schalter 2", sk.GibSchalterwert(2))
	fmt.Println("Ausgangswert aller Lampen", sk.GibLampenStatus())
	fmt.Println()
	sk.Zeichnen(100)
	fmt.Scanln()


	fmt.Println("Schalte 1 an:")
	sk.SchalteLampeAn(1,true)
	sk.	SchaltungBerechnen()

	fmt.Println("Zustand von Schalter 1", sk.GibSchalterwert(1))
	fmt.Println("Zustand von Schalter 2", sk.GibSchalterwert(2))
	fmt.Println("Ausgangswert aller Lampen", sk.GibLampenStatus())
	fmt.Println()
	sk.Zeichnen(100)
	fmt.Scanln()


	fmt.Println("Schalte 2 an:")
	sk.SchalteLampeAn(2,true)
	sk.	SchaltungBerechnen()

	fmt.Println("Zustand von Schalter 1", sk.GibSchalterwert(1))
	fmt.Println("Zustand von Schalter 2", sk.GibSchalterwert(2))
	fmt.Println("Ausgangswert aller Lampen", sk.GibLampenStatus())
	fmt.Println()
	sk.Zeichnen(100)
	fmt.Scanln()

	// sberechneSchaltkreis(be,v)
	


//	ZeichneBauelement()	
	
}
