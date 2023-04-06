

// Martin Seiß	31.3.2023


// generiert die Levelparameter und gibt sie zurück 


package level


import sch "./schaltungen"
//import b "./bauelemente"


type EinzelLevel struct {
	sk 			sch.Schaltung 		// Schaltkreis
	maxPunkte	uint16				// maximal erreichbare Punktzahl
	punkte		uint16				// Punkte ereicht
	minSchalter uint16				// maximale Anzahl von Schaltungen,
									// um maximale Punkte zu erreichen
	xSize		uint16				// Größe der Bauelemente in x-Richtung
}


type Level interface {
	
//	LadeLevel(nummer uint) uint16

	GibLevel(nummer uint16) EinzelLevel
	
	GibPunktzahl(nummer uint16) uint16
	
	SetzePunktzahl(nummer, punkte uint16)
	
}


type data struct {
	elevel []EinzelLevel	// Liste von Eizelleveln
}


func New() *data {
	var lev *data = new(data)
	lev = make(EinzelLevel,0)
	lev = append(lev,baueLevel1())		// Füge Level hinzu
	return data
}


// Funktion zu Bau der Schaltungslevel.
// Jedes Bauelement muss seine eigene ID haben, auch wenn es von einem
// anderen Typ ist.
// Vors: -
// Eff: -
// Erg: Ein Level vom Type Einzellevel ist geliefert.
func baueLevel1() EinzelLevel {

	var elevel EinzelLevel
	var sk sch.Schaltung 	 = sch.New()	// Schaltkreis

	elevel.maxPunkte = 3					// maximal erreichbare Punktzahl
	elevel.punkte = 0						// Punkte ereicht
	elevel.minSchalter	= 2					// maximale Anzahl von Schaltungen,
											// um maximale Punkte zu erreichen
	elevel.xSize = 100						// Größe der Bauelemente in x-Richtung

	sk =  = sch.New()

	
	// --------------   Schalter einfügen   -----------------------//
	sk.BauteilEinfuegen(1,100,100,b.Schalter)
	//sk.SchalteSchalterAn(1,true)			// ggf. Schalter schon anschalten
	sk.BauteilEinfuegen(7,100,400,b.Schalter)
	
	// --------------   Logische Gatter einfügen   ----------------//
	sk.BauteilEinfuegen(3,400,250,b.AND)	
	
	// --------------   Logische Gatter einfügen   ----------------//
	sk.BauteilEinfuegen(4,700,100,b.Lampe)
	sk.BauteilEinfuegen(5,700,400,b.Lampe)

	// --------------   Leitungen einfügen   ----------------------//
	sk.VerbindungEinfuegen(1,3,1,250)
	sk.VerbindungEinfuegen(1,4,1,550)
	sk.VerbindungEinfuegen(7,3,2,250)
	sk.VerbindungEinfuegen(3,5,1,550)


	//  ---------   Schaltkreis wird zu Level hinzugefügt ---------//
	elevel.sk = sk 			

}



func (lev *data) GibLevel(nummer uint16) EinzelLevel {
	return 
}




func GibSchaltkreis(nummer uint16) sch.Schaltung {
	var sk sch.Schaltung
	maxPunkte	uint16				// maximal erreichbar Punktzahl
	minSchalter uint16				// maximale Anzahl von Schaltungen,
									// um maximale Punkte zu erreichen
	punkte		uint16				// ereichte Punkte im Level
	Xsize		uint16				// Größe der Bauelemente in x-Richtung
	

	return
}

	GibMaxPunkte(nummer uint16) uint16
	
	GibMinSchalter(nummer uint16) uint16

	GibXsize(nummer uint16) uint16
	
	GibPunkte(nummer uint16) uint16
	
	SetzePunkte(nummer uint16)

