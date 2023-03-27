package raeume
// Autor: B. Schneider
// Datum: 21.03.2023

import . "gfx"

func Moorhuhn(breite uint16) {
	if breite == 800 {
		LadeBild (0,0, "../../Bilder/Seminarraum800600.bmp")
	}
	Stiftfarbe(100,100,100)
	Vollrechteck(breite/4,breite/8,breite/2,breite/4)
}
