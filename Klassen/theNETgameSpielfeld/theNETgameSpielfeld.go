

//  ADT	- Spielfeld für theNETgame

//	Martin Seiß		29.5.2023


package theNETgameSpielfeld


import "gfx"
import "fmt"
import "../textboxen"
import "../netze"
import "../buttons"
import "time"



type Spielfeld interface {

	netze.Netz

	// Vor: -
	// Erg: Die aktuelle Punktzahl ist geliefert.
	GibPunktzahl() uint16

	// Vor: -
	// Eff: Die Distanz wird um den angegebenen Wert erhöht.
	ErhoeheDistanz(punkte uint16)
	
	// Vor: -
	// Eff. Setzt den Text für die Textbox auf dem Spielfeld. \n
	//		ist als Zeilenumbruch erlaubt.
	// SetzeText(txt string)
	
	
	// Vor: -
	// Erg: Gibt die aktuelle Note aus.
	GibNote() float32
	
	// Vor: -
	// Eff. Setzt die Knotenid, wo sich das Paket befindet.
	SetzePaketID(id uint32)

	// Vor: Ein gfx-Grafikfenster ist geöffnet.
	// Eff: Das Spielfeld wird in einer Schleife wiederholt gezeichnet,
	//		solange das Spiel läuft, also weder gewonnen noch verloren
	//		ist.
	Zeichnen()

	
	// Vor: Ein gfx-Grafikfenster ist geöffnet.
	// Eff: Die Hintergrundmusik wird in einer Schleife wiederholt
	//		gestartet.
	Hintergrundmusik()
	
	
	// Vor: -
	// Eff: Das Spiel und die Uhr wird gestartet und das Paket kann bewegt werden.
	StartGame()


	// Vor: -
	// Eff: True ist geliefert, wenn aktuelles Level gestartet wurde und
	// 		das Spiel noch gewonnen oder verloren ist.
	SpielLaeuft() bool

	
	// Vor: -
	// Eff: True - das nächste Level wird geladen, sonst das
	//		alte Level wird mit neuem Graphen geladen.
	NeuesLevel(neu bool)

}



type data struct {

	weiter,starter,beenden,nochmal buttons.Button		// Buttons
	netze.Netz							// vererbe Computernetz
	time0 uint16						// Startzeit
	ilevel uint16						// aktuelles Level-1
	nlevel uint16						// maximale Anzahl der Level
	txt textboxen.Textbox				// Text für die seitlichen Textbox
	maxPunkte uint16					// maximale Punktzahl beim Start
	punkte uint16						// aktuelle Punktzahl incl. Zeit
	punkteGesamt uint16					// Gesamtpunktzahl
	distanz uint16						// zurückgelegte Distanz nach Kantenlängen
										// Punkte (ohne Zeit) - init = 0
	note float32						// Note
	paketid	uint32						// id des Pakets - init = 0
	start bool							// true - Spiel läuft, false = Stop

}



func New(weiter,starter,beenden,nochmal buttons.Button) *data {
	var sf *data = new(data)
	sf.weiter = weiter
	sf.beenden = beenden
	sf.nochmal = nochmal
	sf.starter = starter
	sf.Netz = netze.New(0.1,0.1)
	sf.nlevel = 5
	sf.txt = textboxen.New(850,400,300,300)
	sf.NeuesLevel(false)
	return sf
}



func (sf *data)	GibPunktzahl() uint16 {
	return sf.punkteGesamt //+ sf.punkte
}



func (sf *data) ErhoeheDistanz(punkte uint16) {
	sf.distanz = sf.distanz + punkte
	//fmt.Println("neue Distanz und Erhöhung: ",sf.distanz,punkte)
}
	

	
func (sf *data) GibNote() float32 {
	return sf.note
}



func (sf *data) SetzePaketID(id uint32) {
	sf.paketid = id
}



func (sf *data) Zeichnen() {

		var dtime uint16				// Zeit seit Start

		for {
			
			// Läuft das Level noch?
			if sf.GibGewonnen() && sf.start {
				if sf.start {sf.schreibeGewonnen()}
				sf.start = false
			} else if sf.GibVerloren() > 0 && sf.start {
				if sf.start {sf.schreibeVerloren()}
				sf.start = false
			}

			// Level neu geladen?
			if !sf.GibGewonnen() && sf.GibVerloren() == 0 && !sf.start {
				dtime = 0
			}

			if sf.start {
				dtime = uint16(float64(time.Now().UnixNano())/1e9) - sf.time0
			}
			
			sf.punkte = sf.maxPunkte - sf.distanz - dtime
			//fmt.Println("punkte,maxPunkte,distanz,dtime: ",sf.punkte,sf.maxPunkte,sf.distanz,dtime)
			
			// Noch Punkte übrig?
			if (sf.punkte <= 0 || sf.punkte > sf.maxPunkte) {
				if sf.start {
					sf.SetzeVerloren(3)
					sf.schreibeVerloren()
					sf.start = false
					sf.nochmal.AktiviereButton()
				}
				sf.punkte = 0
			}

			gfx.UpdateAus()
			sf.zeichneSpielfeld()
			//if sf.start {sf.zeichnePaket()}
			sf.zeichnePaket()
			sf.zeichneButtons()
			gfx.UpdateAn()

			time.Sleep (time.Duration(2e8))
		}

}



func (sf *data) Hintergrundmusik() {
	var soundstr string = "Sounds/Music/30s_Surf.wav"
	for {
		gfx.SpieleSound(soundstr)
		time.Sleep (time.Duration(40e9))
	}
}



func (sf *data)	StartGame() {
	sf.time0 = uint16(float64(time.Now().UnixNano())/1e9)
	sf.start = true
	//go sf.Hindernisse()
}



func (sf *data)	SpielLaeuft() bool {
	return sf.start
}


func (sf *data)	NeuesLevel(neu bool) {
	if neu {
		sf.ilevel++
	}
	sf.ladeLevel(sf.ilevel)
	go sf.Hindernisse()
}




//func (sf *data) CheckGewonnen() bool {			// check: Level gewonnen?


/////////////////////////////////////////////////////////////////

////			Hilfsfunktionen				/////////////////////

/////////////////////////////////////////////////////////////////


func (sf *data) willkommenText() {
	sf.txt.SchreibeText("Willkommen zum NET-Game!\n\n" +
		"Bewegen Sie das Datenpaket entlang der Verbindungen " + 
		"zum Ziel, indem Sie auf die grünen Nachbarrouter " +
		"klicken.\n\n" +
		"Nutzen Sie die kostengünstigste Verbindung." +
		"und achten Sie auf die Zeit.\n\n" +
		"Viel Spaß!")
		
}


func (sf *data) schreibeGewonnen(){
	var erg string
	erg = "Glückwunsch Sie haben gewonnen!\n \n"
	var diff uint16 
	
	
	sf.note = sf.note - 1							// bessere Note
	sf.punkteGesamt = sf.punkteGesamt + sf.punkte	// neuer Punktestand
			
	// !!!! sollte immer stimmen, aber Bug im Code !!!!
	// Berechnung von MinDist oder von distanz fehlerhaft
	if sf.distanz >= uint16(sf.GibMinDist()) {
		diff = sf.distanz - uint16(sf.GibMinDist())
	} else {
		diff = 0
	}
	
	fmt.Println("diff,distanz,MinDist: ",diff,sf.distanz,sf.GibMinDist())
	
	// Wie gewonnen?
	if sf.ilevel+1 == sf.nlevel {		// Letztes Level? => Spiel zu Ende
		sf.weiter.DeaktiviereButton()
		erg = erg + "Glückwunsch, Sie haben alle Level geschafft!\n \n"
		erg = erg + "Auf Wiedersehen Meister!"
		gfx.SpieleSound("Sounds/Applaus.wav")
	} else if sf.ilevel == 1 {	// oder Prüfung bestanden?
		sf.weiter.AktiviereButton()
		//bestanden = true
		erg = erg + "Sie haben damit auch die Prüfung bestanden!\n\n"
		erg = erg + "Auf zur nächsten Aufgabe."	
		gfx.SpieleSound("Sounds/Applaus.wav")
	} else { // oder nur Level gewonnen
		erg = erg + "Sie haben die ideale Route um "+
			fmt.Sprint(diff) +
			" Punkte verfehlt.\n\n"
		
		sf.weiter.AktiviereButton()
		gfx.SpieleSound("Sounds/Sparkle.wav")
	}			
	
	sf.txt.SchreibeText(erg)
}



func (sf *data) schreibeVerloren() {	
	var erg string 
	var soundstr string = "Sounds/GameOver.wav"
	// 1 = Kante gesperrt, 2 Knoten gesperrt,
	// 3 = Bugget zu Ende

	if sf.GibVerloren() == 1 {
		//soundstr = "Sounds/sfx_sounds_negative1.wav"
		erg = erg + "Es tut mit Leid, aber DarthSchmidtar hat das " +
			"Paket abgefangen!\n\n"
		erg = erg + "Aber versuchen Sie es noch einmal."
	} else if sf.GibVerloren()  == 2 {
		erg = erg + "Es tut mit Leid, aber das Datenpaket ging durch " +	
			"einen Routerdefekt leider verloren!\n\n" +
			"Aber versuchen Sie es noch einmal."
	} else if sf.GibVerloren()  == 3 {
		erg = erg + "Es tut mit Leid, aber die Kosten waren zu groß. " +
			"Sie haben leider verloren!\n\n" +			
			"Aber versuchen Sie es noch einmal und sparen Sie Zeit " +
			"und Kosten."
	}

	soundstr = "Sounds/sfx_sounds_negative1.wav"

	gfx.SpieleSound(soundstr)

	sf.txt.SchreibeText(erg)

}



func berechneNote(ilevel uint16) float32 {
	var note float32
	if ilevel == 0  {
		note = 6
	} else if ilevel == 1 {
		note = 5
	} else if ilevel == 2 {
		note = 4
	} else if ilevel == 3 {
		note = 3
	} else if ilevel == 4 {
		note = 2
	}
	// else if ilevel == 5 {
	//		note = 1
	// }
	return note
}


func (sf *data) schreibeSpielstand() {
	gfx.SchreibeFont(20,15,"Level: " + fmt.Sprint(sf.ilevel+1))
	gfx.SchreibeFont(150,15,"Kredit: " + fmt.Sprint(sf.punkte))
	gfx.SchreibeFont(320,15,"Punkte: " + fmt.Sprint(sf.punkteGesamt))
	gfx.SchreibeFont(520,15,"Note: " + fmt.Sprint(sf.note))
}


func (sf *data) zeichneButtons() {
	if sf.weiter.GibAktivitaetButton() {
			sf.weiter.ZeichneButton()
	}
	if sf.starter.GibAktivitaetButton() {
			sf.starter.ZeichneButton()
	}
	if sf.beenden.GibAktivitaetButton() {
			sf.beenden.ZeichneButton()
	}
	if sf.nochmal.GibAktivitaetButton() {
			sf.nochmal.ZeichneButton()
	}
}



func (sf *data) zeichneStart() {
	var x,y uint16 = sf.KnotenKoordinaten(0)
	//var ids []uint32 = netz.KnotenID_Liste()
	r,g,_ := sf.Knotenfarbe(0)
	if r == 255 && g == 255 {
		gfx.Stiftfarbe(r,g,0)
	} else {
		gfx.Stiftfarbe(255,255,255)
	}
	gfx.Vollrechteck(x-30,y-15,60,30)
	gfx.Stiftfarbe(0,0,0)
	gfx.Rechteck(x-30,y-15,60,30)
	gfx.SetzeFont("Schriftarten/Ubuntu-B.ttf",20)
	gfx.SchreibeFont(x-24,y-12,"Start")
}


func (sf *data) zeichneZiel() {
	var max uint32 = sf.GibZielID()
	var x,y uint16 = sf.KnotenKoordinaten(max)
	//var ids []uint32 = netz.KnotenID_Liste()
	r,g,_ := sf.Knotenfarbe(max)
	if r == 255 && g == 255 {
		gfx.Stiftfarbe(r,g,0)
	} else {
		gfx.Stiftfarbe(255,255,255)
	}
	gfx.Vollrechteck(x-30,y-15,60,30)
	gfx.Stiftfarbe(0,0,0)
	gfx.Rechteck(x-30,y-15,60,30)
	gfx.SetzeFont ("Schriftarten/Ubuntu-B.ttf",20)
	gfx.SchreibeFont(x-17,y-12,"Ziel")
}



func (sf *data) zeichneComputer() {
	var max uint32 = sf.GibZielID()
	var x,y uint16 			// netz.KnotenKoordinaten(id)
	var ids []uint32 = sf.KnotenID_Liste()
	for i:=0; i<len(ids); i++ {
		x,y = sf.KnotenKoordinaten(ids[i])
		r,g,_ := sf.Knotenfarbe(ids[i])
		if ids[i] != 0 && ids[i] != max && ids[i] != sf.paketid {
			if r == 255 {
				gfx.LadeBild(x-20,y-25,"Bilder/Computer_klein_rot.bmp")
				gfx.LadeBild(x-20,y-40,"Bilder/Feuer.bmp")
			} else if g == 255 {
				gfx.LadeBild(x-20,y-25,"Bilder/Computer_klein_gruen.bmp")
			} else {
				gfx.LadeBild(x-20,y-25,"Bilder/Computer_klein.bmp")
			}
		}
	}

}



func (sf *data) zeichneDarthSchmidter() {
	//var max uint32 = maxID(netz)	
	var x0,y0,x,y uint16 			// netz.KnotenKoordinaten(id)
	var ids []uint32 = sf.KnotenID_Liste()
	//fmt.Println("Neue Liste")
	for i:=0; i<len(ids); i++ {
		x0,y0 = sf.KnotenKoordinaten(ids[i])
		for j:=0; j<len(ids); j++ {
			//fmt.Println(i,j)		
//			if i == 0 && j == 1 { 
//				gfx.LadeBild(100,100,"../Bilder/DarthSchmidtarExtraTiny.bmp")
//			}
			if sf.Benachbart(ids[i],ids[j]) { //&&  ids[i]<ids[j] {
				r,_,_ := sf.Kantenfarbe(ids[i],ids[j])
				if r == 255 {
					x,y = sf.KnotenKoordinaten(ids[j])
					//fmt.Println("Knotenkoord: ",ids[i],x0,y0)
//					fmt.Println("Kantenkoord: ",ids[i],ids[j],x,y)
					if x > x0 {
						x = x0 + (x-x0)/2
					} else {
						x = x + (x0-x)/2
					}
					if y < 0 {
						y = y0 + (y-y0)/2
					} else {
						y = y + (y0-y)/2
					}
					//fmt.Println("Darth Pos: ",x,y)
					gfx.LadeBild(x-16,y-25,"Bilder/DarthSchmidtarExtraTiny.bmp")
				}
			}
		}
	}

}



func (sf *data) zeichnePaket() {
	var x,y uint16 = sf.KnotenKoordinaten(sf.paketid)
	gfx.LadeBild(x-25,y-25,"Bilder/paket_klein.bmp")
	if sf.GibVerloren() > 0 {
		gfx.LadeBild(x-20,y-40,"Bilder/Feuer.bmp")
	}
}



func (sf *data) zeichneSpielfeld() {

//	var fontsize int = 20

	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	gfx.Stiftfarbe(0,0,0)		

	if sf.GibVerloren() == 0 {
		gfx.LadeBild(840,10,"Bilder/WtheK_black.bmp")
	} else {
		gfx.LadeBild(840,10,"Bilder/WtheK_black_sad.bmp")
	}
	gfx.Linie(830,0,830,700-1)
	gfx.Linie(830,380,1200-1,380)

	sf.schreibeSpielstand()
	sf.txt.Zeichne()		// Schreibe in die Textbox
	
	//if sf.start {
		sf.Darstellen()
		sf.zeichneStart()
		sf.zeichneZiel()
		sf.zeichneComputer()
		sf.zeichneDarthSchmidter()
	//}

}



func (sf *data)	ladeLevel(ilevel uint16) {

	if ilevel == 0 {
		
		sf.start = false 	
		sf.SetzeVerloren(0)
		sf.SetzeGewonnen(false)
		sf.Netz = netze.New(0.,0.)
		sf.txt.SchreibeText("Willkommen zum NET-Game!\n\n" +
			"Bewegen Sie das Datenpaket entlang der Verbindungen " + 
			"zum Ziel, indem Sie auf die grünen Nachbarrouter " +
			"klicken.\n\n" +
			"Nutzen Sie die günstigste Verbindung " +
			"und achten Sie auf die Zeit.\n\n" +
			"Viel Spaß!")	
		sf.maxPunkte = 100
		sf.punkte = sf.maxPunkte
		sf.note = berechneNote(sf.ilevel)
		sf.distanz = 0					
		sf.paketid = 0
		
	} else if ilevel == 1 {
		
		//sf.Netz = netze.New(0.,0.0)
		sf.Netz = netze.New(0.1,0.0)
		sf.txt.SchreibeText("Achtung! Weichen Sie den defekten " + 
			"Routern aus? \n\n")	
		sf.SetzeVerloren(0)
		sf.maxPunkte = 100
		sf.punkte = sf.maxPunkte
		sf.note = berechneNote(sf.ilevel)
		sf.distanz = 0					
		sf.paketid = 0
		sf.start = false 	
		sf.SetzeGewonnen(false)

	} else if ilevel == 2 {
		
		//sf.Netz = netze.New(0.,0.0)
		sf.Netz = netze.New(0.1,0.1)
		sf.txt.SchreibeText("Darth Schmidter versucht das Paket " + 
			"abzufangen. Weichen Sie ihm aus! \n\n")	
		sf.SetzeVerloren(0)
		sf.maxPunkte = 100
		sf.punkte = sf.maxPunkte
		sf.note = berechneNote(sf.ilevel)
		sf.distanz = 0					
		sf.paketid = 0
		sf.start = false 	
		sf.SetzeGewonnen(false)

	} else if ilevel == 3 {
		
		//sf.Netz = netze.New(0.,0.0)
		sf.Netz = netze.New(0.1,0.1)
		sf.txt.SchreibeText("Und jetzt noch schneller! Bei idealer " +
				 "Strecke haben Sie 20 Sekunden Zeit.\n\n")	
		sf.SetzeVerloren(0)
		sf.maxPunkte = uint16(sf.GibMinDist()) + 20
		sf.punkte = sf.maxPunkte
		sf.note = berechneNote(sf.ilevel)
		sf.distanz = 0					
		sf.paketid = 0
		sf.start = false 	
		sf.SetzeGewonnen(false)

	} else if ilevel == 4 {
		
		//sf.Netz = netze.New(0.,0.0)
		sf.Netz = netze.New(0.2,0.2)
		sf.txt.SchreibeText("Oh je, noch mehr Routerausfälle und " +
			"Dark Schmidther will es jetzt wissen! \n\n")	
		sf.SetzeVerloren(0)
		sf.maxPunkte = uint16(sf.GibMinDist()) + 20 
		sf.punkte = sf.maxPunkte
		sf.note = berechneNote(sf.ilevel)
		sf.distanz = 0					
		sf.paketid = 0
		sf.start = false 	
		sf.SetzeGewonnen(false)

	} else {
		panic("Level existiert nicht!")
	}
	
}



