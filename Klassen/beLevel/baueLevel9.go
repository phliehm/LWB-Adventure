

// Martin Seiß	1.4.2023


// Funktion zum Bau von Level 9


package beLevel


import sch "../schaltungen"
import b "../bauelemente"


// Funktion zu Bau der Schaltungslevel.
// Jedes Bauelement muss seine eigene ID haben, auch wenn es von einem
// anderen Typ ist.
// Vors: -
// Eff: -
// Erg: Ein Level vom Type Einzellevel ist geliefert.
func baueLevel9() EinzelLevel {

	var elevel EinzelLevel
	var sk sch.Schaltung 	 = sch.New()	// Schaltkreis

	// ----------------   Setze Level Paramter   ------------------//
	elevel.maxPunkte = 3					// maximal erreichbare Punktzahl
	elevel.punkte = 0						// Punkte ereicht
	elevel.minSchalter	= 1					// maximale Anzahl von Schaltungen,
											// um maximale Punkte zu erreichen
	elevel.xSize = 100						// Größe der Bauelemente in x-Richtung


	// --------------- Baue Schaltkreis ---------------------------//
	elevel.text = make([]string,0)
	elevel.text = append(elevel.text,"Kennen Sie sich auch mit dem")
	elevel.text = append(elevel.text,"XOR-Gatter aus?")
		
	// --------------   Schalter einfügen   -----------------------//
	sk.BauteilEinfuegen(1,100,100,b.Schalter)
	//sk.SchalteSchalterAn(1,true)			// ggf. Schalter schon anschalten
	sk.BauteilEinfuegen(2,100,400,b.Schalter)
	
	// --------------   Logische Gatter einfügen   ----------------//
	sk.BauteilEinfuegen(3,400,250,b.XOR)	
	
	// --------------   Lampen einfügen   -------------------------//
	sk.BauteilEinfuegen(4,700,100,b.Lampe)
	sk.BauteilEinfuegen(5,700,400,b.Lampe)

	// --------------   Leitungen einfügen   ----------------------//
	sk.VerbindungEinfuegen(1,3,1,250)
	sk.VerbindungEinfuegen(3,4,1,550)
	sk.VerbindungEinfuegen(2,3,2,250)
	sk.VerbindungEinfuegen(2,5,1,550)


	//  ---------   Schaltkreis wird zu Level hinzugefügt ---------//
	elevel.sk = sk 
	
	return elevel		

}

