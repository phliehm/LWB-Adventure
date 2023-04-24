// ADT bauelemente - Spezifikation und Implementation

// Martin Seiß    21.3.2023


// package beinhaltet zwei ADTs: Bauelement und Verbindung

// Folgende Bauelemente sind enthalten:
// Schalter, Gatter (AND, OR, NOT), Lampe
// Bauelemente haben zwei Eingänge und einen Ausgang.
// Bei Bauelementen mit normalerweise nur einem Eingang,
// können beide Eingänge belegt werden.
// Es reicht aus, wenn ein Eingang true ist.
// z.B. NOT-Gatter:
//		Eingang 1	|	Eingang 2	|	Ausgang
//		----------------------------------------
//		false		|	false		|   true
//		false		|	true		|	false
//		true		|	false		|	false
//		true		|	true		|	false
//
// Bei Schaltern bedeuten die Eingänge Schalter an und Schalter aus.
// Hier müssen beide Eingänge true oder false sein. Ein Eingang ist 
// true/an, daraus folgt, beide Eingänge an/true, Schalter an.
//
// Bauelelemente können mit Leitungen verbunden werden.


package bauelemente

import l "../leitungen"
import "gfx"
// import "fmt"


var size uint16 = 4				// Linienbreite in Pixel


type Bautyp uint16


const (
	Schalter 	Bautyp = 0
	AND			Bautyp = 1
	OR			Bautyp = 2
	NOT			Bautyp = 3
	XOR			Bautyp = 4
	Lampe		Bautyp = 5
)




type Bauelement interface {
	
	// Vor: n nur 1 oder 2 für Eingang
	// Eff: Der Eingang 1 oder 2 des Bauelements wird auf den gegebenen
	//		Wert true oder false gesetzt.
	// Erg: -	
	SetzeEingang(n uint8, wert bool)

	// Vor: n nur 1 oder 2 für Eingang
	// Eff: -
	// Erg: Fügt die Verbindung in die Liste des Baulementes mit
	// dem betreffenenden Eingang ein.
	VerbindungZumEingang(v l.Leitung)

	// Vor: n nur 1 oder 2 für Eingang
	// Eff: -
	// Erg: Der Wert des Eingangs n des Bauelements ist geliefert.	
	GibEingang(n uint8) bool
	
	// Vor: -
	// Eff: -
	// Erg: Die ID des Bauelements ist geliefert.	
	GibID() uint16
	
	// Vor: -
	// Eff: -
	// Erg: Die ID des Bauelements ist geliefert.	
	GibBauelementtyp () Bautyp 
		
		
	// Vor: -
	// Eff: -
	// Erg: Die Liste der Verbindungen zu den Eingängen des
	// 		Bauelements ist geliefert.	
	GibVerbindungen() []l.Leitung

	// Vor: -
	// Eff: Der Status "Ausgang des Bauteils berechnet" ist auf true
	// 		oder false gesetzt.
	// Erg: -	
	SetzeBerechnet(wert bool)

	// Vor: -
	// Eff: Der aktuelle Ausgangswert des Bauelementes ist auf 
	// 		Grundlage der Eingangswerte neu berechnet.
	// Erg: Der aktuelle Ausgangswert des Bauelementes ist geliefert.
	BerechneAusgang() bool
	
	// Vor: -
	// Eff: -
	// Erg: Der Status "Ausgang des Bauteils berechnet" ist 
	//		geliefert.	
	GibBerechnet() bool
	
	// Vor: -
	// Eff: -
	// Erg: Die x und y Position (Mittelpunkt) des Bauelements ist
	//		geliefert.		
	GibPosXY() (uint16,uint16)


	// Vor: Ein gfx-Grafikfenster ist geöffnet. xSize gibt die Größe
	//		in Pixel des zu zeichnenden Bauelementes in x Richtung an.
	// Eff: Das Bauelement ist an angegebener Stelle gezeichnet.
	// Erg: -
	ZeichneBauelement(xSize uint16)

	
	// Vor: Ein gfx-Grafikfenster ist geöffnet. xSize gibt die Größe
	//		in Pixel des zu zeichnenden Bauelementes in x Richtung an.
	// Eff: Das Leitung ist vom Eingang des Bauelements zur x,y-Position
	//		gezeichnet.
	// Erg: -
	ZeichneLeitung(xSize,x,y uint16, v l.Leitung)

}



type data struct {
	id				uint16		// Nummer des Bauteils
	x,y				uint16		// Position des Mittelpunktes des Bauelementes
	eingang1		bool
	eingang2		bool		// nur wenn zwei Anschlüsse nötig
								// bei Schalter wert = true geschlossen
								// wert = false offen
	ausgang			bool
	typ				Bautyp
	berechnet		bool		// Wurde der Ausgang berechnet?
	verbindungen	[]l.Leitung	// Liste aller Verbindungen, die an den
								// Eingängen enden	
}



func New(id,x,y uint16, eingang1,eingang2,ausgang bool, typ Bautyp) *data {
	var bt *data = new(data)
	bt.id = id
	bt.x = x
	bt.y = y
	bt.eingang1 = eingang1
	bt.eingang2 = eingang2
	if typ == Schalter {	// bei Schalter Eingänge gleicher Wert 
		bt.eingang1 = bt.eingang1 || bt.eingang2
		bt.eingang2 = bt.eingang1 || bt.eingang2
	}
	bt.ausgang = ausgang
	bt.typ = typ
	return bt
}

	
	
func (bt *data) SetzeEingang(n uint8, wert bool) {
	if bt.typ == Schalter {
		if n == 1 || n==2 {
			bt.eingang1 = wert
			bt.eingang2 = wert
		}
	} else {
		if n == 1 {
			bt.eingang1 = wert
		} else if n == 2 {
			bt.eingang2 = wert		
		} else {
			panic("Fehler: Nummer vom Eingang falsch gesetzt. Nur 1 oder 2 erlaubt.")
		}
	}	
}



func (bt *data) GibEingang(n uint8) bool {
	if n == 1 {
		return bt.eingang1
	} else if n == 2 {
		return bt.eingang2	
	} else {
		panic("Fehler: Nummer vom Eingang falsch gesetzt. Nur 1 oder 2 erlaubt.")
	}

}



func (bt *data) GibID() uint16 {
	return bt.id
}



func (bt *data) GibBauelementtyp() Bautyp {
	return bt.typ
}



func (bt *data) VerbindungZumEingang(v l.Leitung) {
	bt.verbindungen = append(bt.verbindungen,v)
}



func (bt *data) GibVerbindungen() []l.Leitung {
	return bt.verbindungen
}



func (bt *data) SetzeBerechnet(wert bool) {
	bt.berechnet = wert
}


func (bt *data) BerechneAusgang() bool {
	if bt.typ == Schalter {
		bt.ausgang = bt.eingang2
	} else if bt.typ == Lampe {
		bt.ausgang = bt.eingang1 || bt.eingang2
	} else if bt.typ == AND {
		bt.ausgang = bt.eingang1 && bt.eingang2
	} else if bt.typ == OR {
		bt.ausgang = bt.eingang1 || bt.eingang2
	} else if bt.typ == NOT { // egal an welchem Eingang Spannung anliegt
		bt.ausgang = !(bt.eingang1 || bt.eingang2) 
	} else if bt.typ == XOR {
		bt.ausgang = (bt.eingang1 && !bt.eingang2) || (!bt.eingang1 && bt.eingang2)
	} else {
		panic("Bauteiltyp nicht bekannt!")		// noch zu spezifizieren!!??
	}
	return bt.ausgang
}


func (bt *data) GibBerechnet() bool {
	return bt.berechnet
}



func (bt *data) GibPosXY() (uint16,uint16) {
	return bt.x,bt.y
}


func dickeLinie(x0,y0,x1,y1,size uint16) {
	size = size/2 + 1
	for i:=uint16(0);i<size;i++ {
		if x0 == x1 {
			gfx.Linie(x0+i,y0,x1+i,y1)
			gfx.Linie(x0-i,y0,x1-i,y1)
		} else {
			gfx.Linie(x0,y0+i,x1,y1+i)
			gfx.Linie(x0,y0-i,x1,y1-i)
		} 

	}
}


func dickesRechteck(x0,y0,b,h,size uint16) {
	for i:=uint16(0);i<size;i++ {
			gfx.Rechteck(x0+i,y0+i,b-2*i,h-2*i)
	}
}


func (bt *data) ZeichneBauelement(xSize uint16) {

	var ySize uint16 = xSize
	var fSize uint16 = xSize/2		// Größe des Fonts

	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont ("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",int(fSize))
	
	if bt.typ == Schalter {
		dickeLinie(bt.x-xSize/2,bt.y,bt.x-xSize/6,bt.y,size)
		dickeLinie(bt.x+xSize/6,bt.y,bt.x+xSize/2,bt.y,size)
		dickeLinie(bt.x+xSize/6,bt.y,bt.x+xSize/6,bt.y-xSize/10,size)
		if bt.ausgang == true {
			dickeLinie(bt.x-xSize/6,bt.y,bt.x+xSize/5,bt.y-xSize/20,size)
		} else {
			dickeLinie(bt.x-xSize/6,bt.y,bt.x+xSize/6,bt.y-xSize/5,size)		
		}
	} else if bt.typ == AND {
		dickesRechteck(bt.x-xSize/2,bt.y-ySize/2,xSize,ySize,size)
		gfx.SchreibeFont (bt.x-fSize/5,bt.y-fSize/2,"&")
	} else if bt.typ == OR {
		dickesRechteck(bt.x-xSize/2,bt.y-ySize/2,xSize,ySize,size)
		gfx.SchreibeFont (bt.x-fSize/2,bt.y-fSize/2,string(rune(0x2265))+"1")
	} else if bt.typ == NOT {
		dickesRechteck(bt.x-xSize/2,bt.y-ySize/2,xSize-xSize/10,ySize,size)
		gfx.Vollkreis(bt.x+xSize/2-xSize/20,bt.y,xSize/20+size)
		gfx.Stiftfarbe(255,255,255)
		gfx.Vollkreis(bt.x+xSize/2-xSize/20,bt.y,xSize/20)
		gfx.Stiftfarbe(0,0,0)
		gfx.SchreibeFont (bt.x-fSize/3,bt.y-fSize/2,"1")
	} else if bt.typ == XOR {
		dickesRechteck(bt.x-xSize/2,bt.y-ySize/2,xSize,ySize,size)
		gfx.SchreibeFont (bt.x-fSize/2,bt.y-fSize/2,"=1")
	} else if bt.typ == Lampe {
		gfx.Vollkreis(bt.x,bt.y,xSize/3)
		if bt.ausgang == true {
			gfx.Stiftfarbe(255,255,0)
		} else {
			gfx.Stiftfarbe(255,255,255)
		}
		gfx.Vollkreis(bt.x,bt.y,xSize/3-size)
		gfx.Stiftfarbe(0,0,0)				
		dickeLinie(bt.x-xSize/3*71/100.,bt.y-xSize/3*71/100.,
			bt.x+xSize/3*71/100.,bt.y+xSize/3*71/100.,size)
		dickeLinie(bt.x-xSize/3*71/100.,bt.y+xSize/3*71/100.,
			bt.x+xSize/3*71/100.,bt.y-xSize/3*71/100.,size)
	}

}


	
func (bt *data) ZeichneLeitung(xSize,x,y uint16, v l.Leitung) {

	if  bt.typ == NOT {
		dickeLinie(bt.x-xSize/2,bt.y,v.GibXPos(),bt.y,size)		
		dickeLinie(v.GibXPos(),bt.y,v.GibXPos(),y,size)
		dickeLinie(v.GibXPos(),y,x,y,size)
	} else if bt.typ == AND || bt.typ == OR || bt.typ == XOR {
		if v.GibEinNr() == 1 {
			dickeLinie(bt.x-xSize/2,bt.y-xSize/4,v.GibXPos(),bt.y-xSize/4,size)
			dickeLinie(v.GibXPos(),bt.y-xSize/4,v.GibXPos(),y,size)
		} else {
			dickeLinie(bt.x-xSize/2,bt.y+xSize/4,v.GibXPos(),bt.y+xSize/4,size)		
			dickeLinie(v.GibXPos(),bt.y+xSize/4,v.GibXPos(),y,size)
		}
		dickeLinie(v.GibXPos(),y,x,y,size)
	} else if bt.typ == Lampe {
		dickeLinie(bt.x-xSize/3,bt.y,v.GibXPos(),bt.y,size)
		dickeLinie(v.GibXPos(),bt.y,v.GibXPos(),y,size)
		dickeLinie(v.GibXPos(),y,x,y,size)
	}

}

