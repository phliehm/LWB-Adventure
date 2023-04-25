

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
func baueLevel8() EinzelLevel {

	var elevel EinzelLevel
	var sk sch.Schaltung 	 = sch.New()	// Schaltkreis

	// ----------------   Setze Level Paramter   ------------------//
	elevel.maxPunkte = 3					// maximal erreichbare Punktzahl
	elevel.punkte = 0						// Punkte ereicht
	elevel.minSchalter	= 4					// maximale Anzahl von Schaltungen,
											// um maximale Punkte zu erreichen
	elevel.xSize = 50						// Größe der Bauelemente in x-Richtung


	// --------------- Baue Schaltkreis ---------------------------//
	elevel.text = make([]string,0)
	elevel.text = append(elevel.text,"Zeigen Sie, was Sie können!")
		
	// --------------   Schalter einfügen   -----------------------//
	sk.BauteilEinfuegen(1,100,100,b.Schalter)
	//sk.SchalteSchalterAn(1,true)			// ggf. Schalter schon anschalten
	sk.BauteilEinfuegen(2,100,175,b.Schalter)
	sk.BauteilEinfuegen(3,100,250,b.Schalter)
	sk.BauteilEinfuegen(4,100,325,b.Schalter)
	sk.BauteilEinfuegen(5,100,400,b.Schalter)
	//sk.SchalteSchalterAn(1,true)			// ggf. Schalter schon anschalten
	sk.BauteilEinfuegen(6,100,475,b.Schalter)
	sk.BauteilEinfuegen(7,100,550,b.Schalter)
	sk.BauteilEinfuegen(8,100,625,b.Schalter)
	
	// --------------   Logische Gatter einfügen   ----------------//
	// 1. NOT-Reihe
	sk.BauteilEinfuegen(11,150,100,b.NOT)	
	sk.BauteilEinfuegen(13,150,250,b.NOT)	
	sk.BauteilEinfuegen(15,150,400,b.NOT)	
	sk.BauteilEinfuegen(17,150,550,b.NOT)

	// 1. Gatter-Reihe	
	sk.BauteilEinfuegen(21,250,125,b.OR)	
	sk.BauteilEinfuegen(22,250,275,b.AND)	
	sk.BauteilEinfuegen(23,250,425,b.XOR)	
	sk.BauteilEinfuegen(24,250,575,b.AND)	

	// 2. NOT-Reihe
	sk.BauteilEinfuegen(31,350,125,b.NOT)	
	sk.BauteilEinfuegen(33,350,425,b.NOT)	

	// 2. Gatter-Reihe	
	sk.BauteilEinfuegen(41,450,200,b.AND)	
	sk.BauteilEinfuegen(42,450,500,b.XOR)	
	
	// 3. NOT-Reihe
	sk.BauteilEinfuegen(51,550,500,b.NOT)	

	// 3. Gatter-Reihe	
	sk.BauteilEinfuegen(61,650,350,b.AND)	
	
	
	
	// --------------   Lampen einfügen   -------------------------//
//	sk.BauteilEinfuegen(71,750,150,b.Lampe)
	sk.BauteilEinfuegen(71,750,350,b.Lampe)
	sk.BauteilEinfuegen(72,750,550,b.Lampe)


	// --------------   Leitungen einfügen   ----------------------//
	// Schalter zu NOT zu Gatter
	sk.VerbindungEinfuegen(1,11,1,125)
	sk.VerbindungEinfuegen(11,21,1,200)
	sk.VerbindungEinfuegen(2,21,2,200)
	sk.VerbindungEinfuegen(3,13,1,125)
	sk.VerbindungEinfuegen(13,22,1,200)
	sk.VerbindungEinfuegen(4,22,2,200)
	sk.VerbindungEinfuegen(5,15,1,125)
	sk.VerbindungEinfuegen(15,23,1,200)
	sk.VerbindungEinfuegen(6,23,2,200)
	sk.VerbindungEinfuegen(7,17,1,125)
	sk.VerbindungEinfuegen(17,24,1,200)
	sk.VerbindungEinfuegen(8,24,2,200)
	// Gatter 1 zu NOT zu Gatter 2
	sk.VerbindungEinfuegen(21,31,1,300)
	sk.VerbindungEinfuegen(31,41,1,400)
	sk.VerbindungEinfuegen(22,41,2,400)
	sk.VerbindungEinfuegen(23,33,1,300)
	sk.VerbindungEinfuegen(33,42,1,400)
	sk.VerbindungEinfuegen(24,42,2,400)
	// Gatter 2 zu NOT zu Gatter 3
	sk.VerbindungEinfuegen(41,61,1,600)
	sk.VerbindungEinfuegen(42,51,1,500)
	sk.VerbindungEinfuegen(51,61,2,600)
	// Gatter 3 zu Lampe 2
	sk.VerbindungEinfuegen(61,71,1,700)
	
	sk.VerbindungEinfuegen(24,72,1,700)


	//  ---------   Schaltkreis wird zu Level hinzugefügt ---------//
	elevel.sk = sk 
	
	return elevel		

}

