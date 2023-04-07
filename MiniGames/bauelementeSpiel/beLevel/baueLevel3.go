

// Martin Seiß	1.4.2023


// Funktion zum Bau von Level 1 


package beLevel


import sch "../schaltungen"
import b "../bauelemente"


// Funktion zu Bau der Schaltungslevel.
// Jedes Bauelement muss seine eigene ID haben, auch wenn es von einem
// anderen Typ ist.
// Vors: -
// Eff: -
// Erg: Ein Level vom Type Einzellevel ist geliefert.
func baueLevel3() EinzelLevel {

	var elevel EinzelLevel
	var sk sch.Schaltung 	 = sch.New()	// Schaltkreis

	// ----------------   Setze Level Paramter   ------------------//
	elevel.maxPunkte = 3					// maximal erreichbare Punktzahl
	elevel.punkte = 0						// Punkte ereicht
	elevel.minSchalter	= 2					// maximale Anzahl von Schaltungen,
											// um maximale Punkte zu erreichen
	elevel.xSize = 100						// Größe der Bauelemente in x-Richtung


	// --------------- Baue Schaltkreis ---------------------------//
	elevel.text = make([]string,0)
	elevel.text = append(elevel.text,"Schaffen Sie auch diese")
	elevel.text = append(elevel.text,"Aufgabe mit dem NOT-Gatter?")
		
	// --------------   Schalter einfügen   -----------------------//
	sk.BauteilEinfuegen(1,100,100,b.Schalter)
	sk.SchalteSchalterAn(1,true)			// ggf. Schalter schon anschalten
	sk.BauteilEinfuegen(2,100,400,b.Schalter)
	
	// --------------   Logische Gatter einfügen   ----------------//
	sk.BauteilEinfuegen(3,300,100,b.NOT)	
	sk.BauteilEinfuegen(4,550,250,b.AND)	
	
	// --------------   Lampen einfügen   -------------------------//
	sk.BauteilEinfuegen(5,700,250,b.Lampe)

	// --------------   Leitungen einfügen   ----------------------//
	sk.VerbindungEinfuegen(1,3,1,200)
	sk.VerbindungEinfuegen(2,4,2,250)
	sk.VerbindungEinfuegen(3,4,1,400)	
	sk.VerbindungEinfuegen(4,5,1,650)


	//  ---------   Schaltkreis wird zu Level hinzugefügt ---------//
	elevel.sk = sk 
	
	return elevel		

}

