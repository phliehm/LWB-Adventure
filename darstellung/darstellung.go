// Autor: A. Cyriacus
// Datum: 05.06.2023
// Zweck: Implementierung des ADO darstellung

package darstellung

/* Der ADO vdarstellung übernimmt die Darstellung der Objekte.
 * Die Darstellung wird vom passenden Controller aufgerufen.
 */

import (
	. "gfx"
	"../MiniGames/1_FP"
	"../MiniGames/bauelementeSpiel"
	//"../MiniGames/2_ALP2/vadeROBIgame"
	//"../MiniGames/3_DDI"
	"../MiniGames/3_DBSA"
	//"../MiniGames/4_BugAttack"
	"../MiniGames/4_Moorhuhn"
	"../MiniGames/theNETgame"
)

// Globale Variablen
// ------------------



// interne Hilfsfunktionen
// ------------------------

func ladeEndeBildschirm() {										//TODO
	Vollrechteck(0,0,1200,700)
	TastaturLesen1()
}


// Methoden
// ---------
	
func MainfloorDarstellen() {
	
	Stiftfarbe(255,255,255)
	Vollrechteck(0,0,1200,700)
	LadeBild(0,50,"./Bilder/mainfloor.bmp")
	
}

func SemesterraumDarstellen(n int) {
	
	switch n {
		case 0:
		ladeEndeBildschirm()
		case 1:
		LadeBild(0,0,"./Bilder/MainGame/raum1.bmp")
		case 2:
		LadeBild(0,0,"./Bilder/MainGame/raum2.bmp")
		case 3:
		LadeBild(0,0,"./Bilder/MainGame/raum3.bmp")
		case 4:
		LadeBild(0,0,"./Bilder/MainGame/raum4.bmp")
		case 5:
		Stiftfarbe(255,255,255)
		Vollrechteck(0,0,1200,700)
		LadeBild(250,50,"./Bilder/Tür5.bmp")
	}
	
}

/*
//Hilfsfunktion BauelementeSpiel:
func bauelemente() {
	var ilevel uint16	  			// aktuelle Levelnummer
	var ePunkte []uint16			// Punkte erreicht im Level

	//ilevel = 3	  								// aktuelle Levelnummer
	//ePunkte = []uint16{3,3,3,0,0,0,0} 			// Punkte erreicht im Level

	ilevel = 0
	ePunkte = []uint16{}
	bauelementeSpiel.BauelementeSpiel(ilevel,ePunkte)
}
*/


func MinigameLaden(raum,n int) {
	
	switch raum {
		
		case 1:
		switch n {
			case 1:
			muster.Muster()
			case 2:
			bauelementeSpiel.BauelementeSpiel()
		}
		
		case 2:
		switch n {
			case 1:
			//vadeROBIgame.Vaderobi()
			case 2:
			//TODO (noch kein importierbares package vorhanden)
		}
		
		case 3:
		switch n {
			case 1:
			sqlGame.SQLgame()
			case 2:
			//TODO (noch kein importierbares package vorhanden)
		}
		
		case 4:
		switch n {
			case 1:
			moorhuhn.Moorhuhn()
			case 2:
			//TODO (noch kein importierbares package vorhanden)
			case 3:
			theNETgame.TheNETgame()
		}
					
	}
}

