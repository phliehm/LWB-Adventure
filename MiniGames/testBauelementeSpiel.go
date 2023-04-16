// ADT bauelemente - Spezifikation und Implementation

// Martin Seiß    21.3.2023


package main

import . "./bauelementeSpiel"

//import "os"
//import "strconv"
import "fmt"




func main() {

	var ilevel uint16	  			// aktuelle Levelnummer
	var ePunkte []uint16			// Punkte erreicht im Level

	ilevel = 3	  								// aktuelle Levelnummer
	ePunkte = []uint16{3,3,3,0,0,0,0} 			// Punkte erreicht im Level

	ilevel = 0
	ePunkte = []uint16{} 

	// -------    Lade Level gegeben auf der Kommandozeile  ------- //
/*
	fmt.Println(len(os.Args))
	if len(os.Args) > 1 {
		intVar, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("Levelargument falsch!")
		}		
		ilevel = uint16(intVar) - 1
	}
*/
	ilevel,note,ePunkte := BauelementeSpiel(ilevel,ePunkte)
	fmt.Println("Note: ",note) 
	fmt.Println("Starte alten Spielstand mit: ",ilevel,ePunkte )
	
}
