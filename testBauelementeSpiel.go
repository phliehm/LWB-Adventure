// ADT bauelemente - Spezifikation und Implementation

// Martin Seiß    21.3.2023


package main

import . "./MiniGames/bauelementeSpiel"

//import "os"
//import "strconv"
import "fmt"
import "gfx"




func main() {

	var note float32	  			// aktuelle Levelnummer
	var punkte uint32				// Punkte erreicht im Level
/*
	ilevel = 3	  								// aktuelle Levelnummer
	ePunkte = []uint16{3,3,3,0,0,0,0} 			// Punkte erreicht im Level

	ilevel = 0
	ePunkte = []uint16{} 
*/
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
	gfx.Fenster(1200,700)
	note,punkte = BauelementeSpiel()
	fmt.Println("Note: ",note) 
	fmt.Println("Punkte: ",punkte) 
	
}
