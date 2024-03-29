// ADT schaltungen - Implementation

// Martin Seiß    22.3.2023 (Start)

// VORSICHT: ID kann nur einmal vergeben werden, auch nicht bei
// verschiedenen Bauelementen

package schaltungen


import b "../bauelemente"
import l "../leitungen"
import "gfx"


type data struct {
	bauelementeTab 	map[uint16]b.Bauelement
	verbindungen 	[]l.Leitung
}


func New() *data {
	var sch *data = new(data)
	sch.bauelementeTab = make( map[uint16]b.Bauelement,0)
	sch.verbindungen = make([]l.Leitung,0)
	return sch
}

	
func (sch *data) BauteilEinfuegen(id,x,y uint16,typ b.Bautyp) {
	if _,ok:=sch.bauelementeTab[id];ok { panic ("FEHLER: Bauelement ID existiert schon!!") }
	sch.bauelementeTab[id]=b.New(id,x,y,false,false,false,typ)
}

	
func (sch *data) VerbindungEinfuegen(vonID,nachID,eingangNr,x uint16) {
		sch.verbindungen = append(sch.verbindungen,l.New(vonID,nachID,eingangNr,x))
		sch.bauelementeTab[nachID].VerbindungZumEingang(l.New(vonID,nachID,eingangNr,x))
}

	
func (sch *data) PruefeSchaltung() bool  {
	return false
}

	
func (sch *data) SchaltungBerechnen() {

	var verb []l.Leitung				// Liste der Verbindungen am Eingang			

	// Schritt 1:
	// Berechne Ausgang von Schaltern und 
	// setze die Eingänge der anderen Bauelemente auf false
	// setze berechnet[index] = true für alle Schalter,
	for _,be:= range sch.bauelementeTab {
		if be.GibBauelementtyp() == b.Schalter {
			be.BerechneAusgang()
			be.SetzeBerechnet(true)
		} else {
			be.SetzeEingang(1, false)
			be.SetzeEingang(2, false)
			be.SetzeBerechnet(false)
		}
	}

	// Setze Eingänge, wenn alle Ausgänge der Verbindungen bereits berechnet.
	// Suche vom Eingang aus die zuegehörigen Ausgänge.
	// Berechne den Ausgang, wenn alle Eingänge gesetzt werden konnten.
	// Wiederhole bis alle Bauteile berechnet
	var e1,e2 bool			// neue Werte der Eingänge?
	var	ok	bool			// alle Bauteile berechnet?
	for {					// wiederhole bis alle berechnet: ok == true
		ok = true
A:		for _,be:= range sch.bauelementeTab {
			e1 = false
			e2 = false
			if !be.GibBerechnet() {
				ok = false
				verb = be.GibVerbindungen()
				for _,v:= range verb {		// alle verbundenen Ausgänge berechnet?
					if sch.bauelementeTab[v.GibVonID()].GibBerechnet() {
						if v.GibEinNr() == 1 {
							// Nur eine Eingangsleitung muss jeweils true sein (Spannung haben).
							e1 = e1 || sch.bauelementeTab[v.GibVonID()].BerechneAusgang()
						} else if v.GibEinNr() == 2 {
							e2 = e2 || sch.bauelementeTab[v.GibVonID()].BerechneAusgang()
						}
					} else {
						continue A			// noch nicht alle zugehörigen Ausgänge
					}						// berechnet. Abbruch!
	
				}
				// Eingänge können gesetzt werden und der Ausgang wird berechnet.
				be.SetzeEingang(1, e1)
				be.SetzeEingang(2, e2)
				be.BerechneAusgang()
				be.SetzeBerechnet(true)
			}
		}
		if ok {break}			// ja alle Bauteile berechnet
	}

}

		
func (sch *data) SchalteSchalterAn(id uint16, wert bool) {

	if _,ok:=sch.bauelementeTab[id];!ok { panic ("FEHLER: Bauelement ID existiert nicht!") }
	if sch.bauelementeTab[id].GibBauelementtyp() == b.Schalter {
		sch.bauelementeTab[id].SetzeEingang(1, wert)
		sch.bauelementeTab[id].SetzeEingang(2, wert)
		sch.bauelementeTab[id].BerechneAusgang()
	} else {
		panic("FEHLER: Baulement ist keine Lampe")
	}
	
}


func (sch *data) GibSchalterwert(id uint16) bool {
	if _,ok:=sch.bauelementeTab[id];!ok { panic ("FEHLER: Bauelement ID existiert nicht!") }
	if sch.bauelementeTab[id].GibBauelementtyp() == b.Schalter {
		return sch.bauelementeTab[id].BerechneAusgang()
	} else {
		panic("FEHLER: Baulement ist kein Schalter")
	}
	return false	
}



func (sch *data) GibLampenwert(id uint16) bool {
	if _,ok:=sch.bauelementeTab[id];!ok { panic ("FEHLER: Bauelement ID existiert nicht!") }
	if sch.bauelementeTab[id].GibBauelementtyp() == b.Lampe {
		return sch.bauelementeTab[id].BerechneAusgang()
	} else {
		panic("FEHLER: Baulement ist keine Lampe")
	}
	return false
}


func (sch *data) GibLampenStatus() []bool {
	var erg []bool = make([]bool,0)
	for _,be:= range sch.bauelementeTab {
		if be.GibBauelementtyp() == b.Lampe {
			erg = append(erg,be.BerechneAusgang())
		}
	}
	return erg
}


func (sch *data) GibBauelementtyp(id uint16) b.Bautyp {
	return sch.bauelementeTab[id].GibBauelementtyp()
}


func (sch *data) GibSchalterIDs() []uint16 {
	var erg []uint16 = make([]uint16,0)
	for index,be:= range sch.bauelementeTab {
		if be.GibBauelementtyp() == b.Schalter {
			erg = append(erg,index)
		}
	}
	return erg
}


func (sch *data) GibPosXY(id uint16) (uint16,uint16) {
	return sch.bauelementeTab[id].GibPosXY()
}



func (sch *data) Zeichnen(xSize uint16) {
	var verb []l.Leitung	
	gfx.Stiftfarbe(0,0,0)	
	for _,be:= range sch.bauelementeTab {
		be.ZeichneBauelement(xSize)
		verb = be.GibVerbindungen()
		for _,v:= range verb {
				x,y := sch.bauelementeTab[v.GibVonID()].GibPosXY()
				be.ZeichneLeitung(xSize,x+xSize/2,y,v)
		}		
	}
}



