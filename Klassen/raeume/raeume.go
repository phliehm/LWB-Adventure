package raeume
// Autor: B. Schneider
// Datum: 21.03.2023

import . "gfx"

func Moorhuhn() {
	LadeBild (0,0, "../../Bilder/Seminarraum-3.bmp")
}

func Hauptflur(breite uint16) {
	Stiftfarbe(200,200,200)
	Vollrechteck(0,0,breite,breite*3/4)
	Stiftfarbe(100,100,100)
	Vollrechteck(breite/4,0,breite/2,breite*3/4)
	
}
