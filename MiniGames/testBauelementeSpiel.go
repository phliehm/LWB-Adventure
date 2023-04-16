// ADT bauelemente - Spezifikation und Implementation

// Martin Seiß    21.3.2023


package main

import . "./bauelementeSpiel"

import "os"
import "strconv"




func main() {

	var ilevel uint16	  			// aktuelle Levelnummer

	// -------    Lade Level gegeben auf der Kommandozeile  ------- //
	if len(os.Args) > 1 {
		intVar, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("Levelargument falsch!")
		}		
		ilevel = uint16(intVar) - 1
	}

	BauelementeSpiel(ilevel) 
	
}
