// ADT bauelemente - Spezifikation

// Martin Seiß    21.3.2023

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

//func New(id,x,y uint16, eingang1,eingang2,ausgang bool, typ Bautyp) *data


// Font muss ggf einmal gesetzt werden
// z.B: and.SetzeFont("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf")
// gfx.SetzeFont("Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf")


package bauelemente


// ------------  importierte Pakete und Klassen ---------------- //
import l "../leitungen"
//import "gfx"


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
	//		Der Font wurde gesetzt.
	// Eff: Der Font wird gesetzt.
	// Eff: Das Leitung ist vom Eingang des Bauelements zur x,y-Position
	//		gezeichnet.
	// Erg: -
	ZeichneLeitung(xSize,x,y uint16, v l.Leitung)
	
	// Vor: Ein gfx-Grafikfenster ist geöffet.
	SetzeFont(font string)

}
