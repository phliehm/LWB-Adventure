// ADT schaltungen - Spezifikation

// Martin Seiß    22.3.2023 (Start)

// VORSICHT: ID kann nur einmal vergeben werden, auch nicht bei
// verschiedenen Bauelementen

// Erg: Eine noch leere Schaltung ist geliefert. Die Bauelemente und
// 		Leitungen müssen noch eingefügt werden.
// func New() *data {

package schaltungen

//  ----------    importierte Pakete und Klassen    ---------------//
import b "../bauelemente"
//import l "../leitungen"
//import "gfx"


type Schaltung interface {

	// Vor: Bauteil ID ist noch nicht vergeben.
	// Eff: Ein Bauteil ist dem Schaltkreis hinzugefügt.
	// Erg: -		
	BauteilEinfuegen(id,x,y uint16,typ b.Bautyp)

	// Vor: Bauteile, die verbunden werden sollen. existieren.  
	// Eff: Eine Verbindung ist dem Schaltkreis hinzugefügt.
	// Erg: -			
	VerbindungEinfuegen(vonID,nachID,eingangNr,x uint16)
	
//	PruefeSchaltung() bool

	// Vor: - 
	// Eff: Alle Ausgänge der Bauteile werden systematisch berechnet,
	//		so dass die aktuellen Lampenwerte ermittelt werden könenn.
	// Erg: -				
	SchaltungBerechnen()
	
	// Vor: Bauelement mit id existiert. 
	// Eff: Setzt den Schalter auf an/true oder auf aus/false.
	// Erg: -				
	SchalteSchalterAn(id uint16, wert bool)

	// Vor: Bauelement mit id existiert. 
	// Eff: -
	// Erg: Gibt den aktuellen Schalterwert aus (an/true oder aus/false).				
	GibSchalterwert(id uint16) bool

	// Vor: - 
	// Eff: -
	// Erg: Gibt den aktuellen Lampenwert aus (an/true oder aus/false).				
	GibLampenwert(id uint16) bool

	// Vor: - 
	// Eff: -
	// Erg: Gibt die aktuellen Werte aller Lampen als Liste aus.				
	GibLampenStatus() []bool

	// Vor: Bauelement mit id existiert. 
	// Eff: -
	// Erg: Gibt die Typ des  Bauelementes aus.					
	GibBauelementtyp(id uint16) b.Bautyp

	// Vor: - 
	// Eff: -
	// Erg: Gibt die ID des Schalters aus.						
	GibSchalterIDs() []uint16 
	
	// Vor: Bauelement mit id existiert.
	// Eff: -
	// Erg: Die x und y Position (Mittelpunkt) des Bauelements mit
	//		passender ID ist geliefert.		
	GibPosXY(id uint16) (uint16,uint16)

	// Vor: -
	// Eff: -
	// Erg: Der Schaltkreis ist gezeichnet.		
	Zeichnen(xSize uint16)
		
}

