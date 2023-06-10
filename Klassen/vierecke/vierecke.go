package vierecke

// Annalena Cyriacus
// 24.04.2023

// Der ADT Viereck dient der Verwaltung von Vierecken. Jedes Viereck
// hat vier Koordinaten, nämlich (xA,yA), (xB,yB), (xC,yC) und (xD,yD)
// durch die es festgelegt wird. A ist dabei die obere linke Ecke.

// Vor.: -
// Erg.: Ein Viereck mit den Koordinaten (xA,yA), (xB,yB), (xC,yC) und
//       (xD,yD) ist geliefert. A ist dabei die obere linke Ecke.
// New (xA,yA,xB,yB,xC,yC,xD,yD uint16) *data // *data erfüllt das Interface Viereck

type Viereck interface {
	
	SetKoordinaten(xA,yA,xB,yB,xC,yC,xD,yD uint16)
	
	// Vor.: -
	// Erg.: Die Koordinaten der vier Eckpunkte des Vierecks sind
	//       geliefert.
    GetKoordinaten () (xA,yA,xB,yB,xC,yC,xD,yD uint16)
    
    GibFarbe() (r,g,b uint8)
    
    SetzeFarbe(r,g,b uint8) 
	
	// Vor.: -
	// Erg.: Eine Repräsentation des Vierecks in Form eines Strings
	//       ist geliefert.
	String () string
	
	// Vor.: -
	// Erg.: Ein stets gleichlanger Bytestrom, der die serialisierte Form
	//       des Vierecks darstellt, ist geliefert.
	Kodieren () []byte
	
	// Vor.: -
	// Eff.: Die Tür hat jetzt die Eigenschaften, die im Bytestrom
	//       enthalten waren. Seine vorher vorhandenen Eigenschaften
	//       gibt es nicht mehr.
	Dekodieren ([]byte)
	
	// Vor.: -
	// Erg.: Eine tiefe Kopie des Vierecks v in Form einer Variablen des
	//       Typs interface {} ist geliefert. Die Objekte v und 
	//       v.Kopie().(Viereck) sind unterschiedliche Objekte, haben aber
	//       komplett identische Eigenschaften.
	Kopie () interface {}
	
	// Vor.: -
	// Erg.: Das Viereck ist in der aktuellen Stiftfarbe gezeichnet.
	Zeichnen()

	Angeklickt(x,y uint16) bool
	
	AktiviereKlickbar()
	
	DeaktiviereKlickbar()

}

