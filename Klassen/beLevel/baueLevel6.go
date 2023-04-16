

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
func baueLevel6() EinzelLevel {

	var elevel EinzelLevel
	var sk sch.Schaltung 	 = sch.New()	// Schaltkreis

	// ----------------   Setze Level Paramter   ------------------//
	elevel.maxPunkte = 3					// maximal erreichbare Punktzahl
	elevel.punkte = 0						// Punkte ereicht
	elevel.minSchalter	= 1					// maximale Anzahl von Schaltungen,
											// um maximale Punkte zu erreichen
	elevel.xSize = 75						// Größe der Bauelemente in x-Richtung


	// --------------- Baue Schaltkreis ---------------------------//
	elevel.text = make([]string,0)
	elevel.text = append(elevel.text,"So jetzt wird es ein wenig")
	elevel.text = append(elevel.text,"schwerer!")
		
	// --------------   Schalter einfügen   -----------------------//
	sk.BauteilEinfuegen(1,100,100,b.Schalter)
	//sk.SchalteSchalterAn(1,true)			// ggf. Schalter schon anschalten
	sk.BauteilEinfuegen(2,100,250,b.Schalter)
	sk.BauteilEinfuegen(3,100,400,b.Schalter)
	sk.BauteilEinfuegen(4,100,550,b.Schalter)
	
	// --------------   Logische Gatter einfügen   ----------------//
	sk.BauteilEinfuegen(5,200,100,b.NOT)	
	sk.BauteilEinfuegen(6,350,200,b.AND)	
	sk.BauteilEinfuegen(7,500,200,b.NOT)	
	sk.BauteilEinfuegen(8,350,450,b.OR)	
	sk.BauteilEinfuegen(9,600,350,b.AND)	
	
	// --------------   Lampen einfügen   -------------------------//
	sk.BauteilEinfuegen(10,700,350,b.Lampe)

	// --------------   Leitungen einfügen   ----------------------//
	sk.VerbindungEinfuegen(1,5,1,150)
	sk.VerbindungEinfuegen(5,6,1,270)
	sk.VerbindungEinfuegen(2,6,2,200)
	sk.VerbindungEinfuegen(3,8,1,200)
	sk.VerbindungEinfuegen(4,8,2,200)	
	sk.VerbindungEinfuegen(6,7,1,420)	
	sk.VerbindungEinfuegen(7,9,1,550)	
	sk.VerbindungEinfuegen(8,9,2,450)	
	sk.VerbindungEinfuegen(9,10,1,650)	


	//  ---------   Schaltkreis wird zu Level hinzugefügt ---------//
	elevel.sk = sk 
	
	return elevel		

}

