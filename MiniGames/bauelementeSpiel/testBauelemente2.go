// ADT bauelemente - Spezifikation und Implementation

// Martin Seiß    21.3.2023


package main

import b "./bauelemente"
import "fmt"
//import "gfx"

// Vor: -
// Eff: Die Ein- und Ausgänge aller Bauelemente werden im Schaltkreis
//		berechnet.
// Erg: -
func berechneSchaltkreis(be []b.Bauelement,v []b.Leitung) {

/*
	var Nbe	= len(be)			// Anzahl der Bauelemente
	var Nv	= len(v)			// Länge der Verbindungen
	var berechnet []bool = make([]bool,Nbe)	// Liste der berechneten Ausgänge
	var anr, enr uint16			// Nummer der Bauteile mit Ein-/Ausgang
	var lwert bool				// bool-Wert, der übertragen wird

	// Schritt 1:
	// Berechne Ausgang von Schaltern und 
	// setze die Eingänge der anderen Bauelemente auf false
	for i:=0; i<Nbe; i++ {
		if be[i].typ == Schalter {
			be[i].BerechneAusgang()
			berechnet[i] = true
		} else {
			be[i].SetzeEingang(1,false)
			be[i].SetzeEingang(2,false)
		}
	}
	// Setze Eingänge, wenn alle Ausgänge der Verbindungen bereits berechnet.
	// Suche vom Eingang aus die zuegehörigen Ausgänge.
	// Berechne den Ausgang, wenn alle Eingänge gesetzt werden können.
	for i:=0; i<Nbe; i++ {
		if !berechnet[i] { // Ausgang von Bauteil noch nicht berechnet
			enr = be[i].nummer
			for j:=0; i<Nv; j++ {	// Suche Verbindungen zum Bauteil
				if v[j].einBauteilNr == be[i].nummer && v[j].ausBauteilNr {
						
				}
			} 
		}
	}
*/	
/*
	for i:=0; i<Nv; i++ {
		anr = v[i].ausBauteilNr
		enr = v[i].einBauteilNr
		for i:=0; i<Nbe; i++ {
			if be[]
		}
	}
*/

}



func main() {

// var be []Bauelement = make()
	
// Schalter 1
//	var ein1 b.Anschluss = b.Anschluss{0,false}
//	var ein2 b.Anschluss = b.Anschluss{0,false}
//	var aus b.Anschluss = b.Anschluss{1,false}

	
	var be []b.Bauelement = make([]b.Bauelement,0)
	var v []b.Leitung = make([]b.Leitung,0)
	
	// Schalter 1
	be = append(be,b.New(0,100,100,false,false,false,b.Schalter))
	// Lampe 1
	be = append(be,b.New(1,300,200,false,false,false,b.Lampe))

	// Verbindungen
	v = append(v,b.Leitung{200,0,1,1})


	fmt.Println("Eingangswerte von Schalter 1", be[0].GibEingang(1),be[0].GibEingang(2))
	fmt.Println("Eingangswerte von Lampe 1", be[1].GibEingang(1),be[1].GibEingang(2))
	fmt.Println("Ausgangswert von Lampe 1", be[1].BerechneAusgang())
	fmt.Println()

	fmt.Println("Schalte 1 an:")
	be[0].SetzeEingang(2, true)	// Schalter an

	berechneSchaltkreis(be,v)
	
	fmt.Println("Eingangswerte von Schalter 1", be[0].GibEingang(1),be[0].GibEingang(2))
	fmt.Println("Eingangswerte von Lampe 1", be[1].GibEingang(1),be[1].GibEingang(2))
	fmt.Println("Ausgangswert von Lampe 1", be[1].BerechneAusgang())
	fmt.Println()

//	ZeichneBauelement()	
	
}
