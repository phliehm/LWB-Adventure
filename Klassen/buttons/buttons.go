// ADT button - Spezifikation und Implementation

// Martin Seiß    29.3.2023

package buttons

import "gfx"
// import "fmt"

// Buttons haben recheckige From.
// x,y - bezeichet die obere linke Ecke.
// b,h - die Breite und Höhe
// r,g,b - Hintergrundfarbwerte des Buttons
// aktiv - Buttion ist aktiv und bereit zur Positionsbestimmung
// beschriftung	- String mit Beschriftung des Buttons

// New(x,y,h,br,r,g,b,form,aktiv,beschriftung) 


type Button interface {

	// Vor: Button ist aktiv.
	// Eff: -
	// Erg: True ist geliefert, wenn die x-y-Position mit dem 
	//		Buttonbereich übereinstimmt, andernfalls False.
	TesteXYPosInButton(x,y uint16) bool

	// Vor: -
	// Eff: Button ist auf aktiv gestellt. x,y-Position kann gecheckt
	//		werden und Button ist schwarz-farbig gezeichnet.
	// Erg: -
	AktiviereButton()
	
	// Vor: -
	// Eff: Button ist auf inaktiv gestellt. x,y-Position kann gecheckt
	//		werden, liefert aber immer false, und Button ist 
	//		grau-blassfarbig gezeichnet.
	// Erg: -
	DeaktiviereButton()

	// Vor: -
	// Eff: -
	// Erg: Der Aktivwert des Buttons ist geliefert.
	GibAktivitaetButton() bool
	
	// Vor: -
	// Eff: Die Beschriftung des Button ist geändert.
	// Erg: -
	AendereBeschriftung(text string) 

	// Vor: -
	// Eff: Der angegebene Font wird beim Schreiben verwendet.
	// Erg: -
	SetzeFont(font string)

	// Vor: gfx Fenster ist geöffnet.
	// Eff: Der Button ist gezeichnet.
	// Erg: -
	ZeichneButton()

}



type data struct {
	x,y 		uint16		// bezeichet die obere linke Ecke.
	br,h 		uint16		// die Breite und Höhe des Button
	r,g,b 		uint8		// Hintergrundfarbwerte des Buttons
	aktiv		bool		// true = Button aktiv
	beschriftung 	string	// String mit Beschriftung des Buttons
	font			string	// Name mit Pfad von Font
}



func New(x,y,br,h uint16, r,g,b uint8, aktiv bool, beschriftung string) *data {
	var but *data = new(data)
	but.x = x
	but.y = y
	but.h = h
	but.br = br
	but.r = r
	but.g = g
	but.b = b
	but.aktiv = aktiv
	but.beschriftung = beschriftung
	but.font = "../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf"		// Standartfont
	return but
}



func (but *data) TesteXYPosInButton(x,y uint16) bool {
	if but.aktiv {
		if but.x <= x &&  x <= but.x+but.br {
			if but.y <= y &&  y <= but.y+but.h{
				return true
			}
		} 
	}
	return false
}



func (but *data) AktiviereButton() {
	but.aktiv = true
}

	
	
func (but *data) DeaktiviereButton() {
	but.aktiv = false
}
	
	
func (but *data) GibAktivitaetButton() bool {
	return but.aktiv
}

	
func (but *data) AendereBeschriftung(text string) {
	but.beschriftung = text
}	
	
	
	
func (but *data) SetzeFont(font string) {
	but.font = font
}

	
	
func (but *data) ZeichneButton() {
	var diff uint16 = 2
	var fsize int = int(but.h/2)
	
	gfx.SetzeFont(but.font,fsize)

	if but.aktiv {
		gfx.Stiftfarbe(but.r,but.g,but.b)
		gfx.Vollrechteck(but.x,but.y,but.br,but.h)
		gfx.Stiftfarbe(0,0,0)
		gfx.Rechteck(but.x,but.y,but.br,but.h)
		gfx.Rechteck(but.x+diff,but.y+diff,but.br-2*diff,but.h-2*diff)
		gfx.SchreibeFont(but.x+10,but.y+but.h/2-uint16(fsize)/2,but.beschriftung)
	} else {
		gfx.Stiftfarbe(but.r/2,but.g/2,but.b/2)
		gfx.Vollrechteck(but.x,but.y,but.br,but.h)
		gfx.Stiftfarbe(100,100,100)
		gfx.Rechteck(but.x,but.y,but.br,but.h)
		gfx.Rechteck(but.x+diff,but.y+diff,but.br-2*diff,but.h-2*diff)
		gfx.SchreibeFont(but.x+10,but.y+but.h/2-uint16(fsize)/2,but.beschriftung)		
		gfx.Stiftfarbe(0,0,0)
	}
}
