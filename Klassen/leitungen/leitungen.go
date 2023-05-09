// ADT leitungen - Spezifikation und Implementation

// Martin Seiß    21.3.2023

package leitungen


type Leitung interface {
	
	// Erg: -
	// Eff: -
	// Erg: Gibt die ID vom Bauelement aus, woher die Verbindung kommt.
	GibVonID() uint16

	// Erg: -
	// Eff: -
	// Erg: Gibt die ID vom Bauelement aus, wohin die Verbindung geht.
	GibNachID() uint16

	// Erg: -
	// Eff: -
	// Erg: Gibt die Eingangsnummer vom Bauelement aus, wohin die Verbindung geht.
	GibEinNr() uint16
	
	// Erg: -
	// Eff: -
	// Erg: Gibt die x-Position vom Knick der Verbindung aus.
	GibXPos() uint16
	
}


type data struct {
	ausBauteilID 	uint16		// Nummer des Bauteils, wo Leitung startet
	einBauteilID	uint16		// Nummer des Bauteils, wo Leitung endet
	einNr			uint16		// Nummer des Eingangs (1 oder 2), wo Leitung endet
	x				uint16		// x-Positon des Leitungsknicks
}



func New(vonID, nachID, einNr, x uint16) *data {
	var l *data = new(data)
	l.ausBauteilID = vonID
	l.einBauteilID = nachID
	l.einNr = einNr
	l.x = x
	return l
}


func (l *data) GibVonID() uint16 {
	return l.ausBauteilID
}


func (l *data) GibNachID() uint16 {
	return l.einBauteilID
}
	
	
func (l *data) GibEinNr() uint16 {
	return l.einNr
}

	
func (l *data) GibXPos() uint16 {
	return l.x
}

