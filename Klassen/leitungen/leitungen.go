// ADT leitungen - Spezifikation

// Martin Seiß    21.3.2023

// Erg: Eine Instanz des ADT leitungen ist geliefert. Die Leitung
// 		verläuft vom Ausgang eines Bauelementes (vonID) zu
//		einem von zwei Eingängen (einNr = 1 oder 2) eines anderen
//		Bauelementes (nachID). x Bezeichnet die x-Position des 
//		Knicks in der Leitung.
// func New(vonID, nachID, einNr, x uint16) *data 


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
