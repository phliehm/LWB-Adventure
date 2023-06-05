// Autor: A. Cyriacus
// Datum: 05.06.2023
// Zweck: Implementierung des ADO darstellung

package darstellung

/* Der ADO vdarstellung übernimmt die Darstellung der Objekte.
 * Die Darstellung wird vom passenden Controller aufgerufen.
 */

import (
	. "gfx"
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
		LadeBild(0,50,"./Bilder/Seminarraum_1_neu2.bmp")
		case 2:
		LadeBild(0,50,"./Bilder/Seminarraum_2_neu.bmp")
		case 3:
		LadeBild(0,50,"./Bilder/Seminarraum_3_neu.bmp")
		case 4:
		LadeBild(0,50,"./Bilder/Seminarraum_4.bmp")
		case 5:
		Stiftfarbe(255,255,255)
		Vollrechteck(0,0,1200,700)
		LadeBild(250,50,"./Bilder/Tür5.bmp")
	}
	
}
