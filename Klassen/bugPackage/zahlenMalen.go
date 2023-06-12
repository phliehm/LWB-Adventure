/* Philipp Liehm
 * Juni 2023
 * LWB-Adventure: BugAttack
 * 
 * Malt Pixel für Pixel einzelne Zahlen
 */
package bugPackage

import (
		"gfx"
		)

// Box hat Höhe von 14, Breite 9 (mit schwarzer Umrandung von überall 1px)

// Eine 1 wird gezeichnet
func male_1(x,y uint16){
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(x,y,zB,zH)
	gfx.Stiftfarbe(sr,sg,sb)
	gfx.Punkt(x+2,y+4)
	gfx.Punkt(x+3,y+3)
	gfx.Linie(x+4,y+2,x+4,y+11)
	gfx.Linie(x+2,y+11,x+6,y+11)
	
}

// Eine 0 wird gezeichnet
func male_0(x,y uint16) {
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(x,y,zB,zH)
	gfx.Stiftfarbe(sr,sg,sb)
	gfx.Linie(x+3,y+2,x+5,y+2)
	gfx.Linie(x+3,y+11,x+5,y+11)
	gfx.Linie(x+2,y+3,x+2,y+10)
	gfx.Linie(x+6,y+3,x+6,y+10)	
}

// Ein Feld wird mit schwarz übermalt
func male_schwarz(x,y uint16) {
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(x,y,zB,zH)
}

// Zeichne eine 0,1 oder schwarz
func male_Zahl(x,y uint16,z uint8) {
	switch z {
		case 0: male_0(x,y)
		case 1: male_1(x,y)
		case 2: male_schwarz(x,y)
	}
}
