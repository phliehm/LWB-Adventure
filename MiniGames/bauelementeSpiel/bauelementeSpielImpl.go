// Paket zum Bauelementespiel

// Implementation

// Martin Seiß    21.3.2023


package bauelementeSpiel


import "fmt"
import "gfx"
import "time"
import sch "../../Klassen/schaltungen"
import level "../../Klassen/beLevel"
import "../../Klassen/buttons"

var path string = "" //"../"


func BauelementeSpiel() (float32,uint32) {

	var ilevel uint16	  			// aktuelle Levelnummer
	var ilevelGeschafft	uint16		// höchstes geschafftes Level
	var nlevel uint16				// Anzahl der Level
	var nPunkte uint16				// neue Punkte im Level
	var ePunkte []uint16			// Punkte erreicht im Level
	var gPunkte	uint16				// Gesamtpunkte erreicht
	var maxPunkte uint16			// maximale erreichbare Geamtpunktzahl
	var happy bool = true			// Winnie sieht happy aus
	var levelNeuLaden bool			// Level neu laden
	var neuZeichnen bool 			// Schaltkreis neu zeichnen
	var bestanden bool				// Prüfung bestanden?
									
	var text []string = WillkommenText()
	var font string = path + "Schriftarten/Ubuntu-B.ttf"
	var soundAn bool				// soll Sound gespielt werden?


	// ---------------- Erzeuge Schaltkreis ----------------------- //
	// ---------------- und lade Level 1 ---------------------------//
	var lev level.Level = level.New()
	var sk sch.Schaltung = lev.GibSchaltkreis(ilevel)
	var xSize uint16 = lev.GibXSize(ilevel) 		
							// Größe des Bauelements in x-Richtung
							// in Pixel zur Skalierung
	nlevel = lev.AnzahlLevel()
	if len(ePunkte) != int(nlevel) {
		ePunkte = make([]uint16,nlevel)
	}
	sk.SchaltungBerechnen()
	nPunkte = lev.GibMaxPunktzahl(ilevel) + lev.GibMinSchalter(ilevel)
	// Zähle maximale erreichbare Punktzahl
	for i:=uint16(0); i<nlevel; i++ { 
		maxPunkte = maxPunkte + lev.GibMaxPunktzahl(i)
	}


	//  --------------------   Buttons ------------------------------//
	
	// erzeuge eine Tabelle von Buttons zu den zugehörigen Schaltern //
	// id gibt die Zuordnung
	var sbutton map[uint16]buttons.Button = makeSchalterbuttonTab(sk,xSize)

	// Erzeuge Buttons zur Spielsteuerung
	var weiter,zurueck,beenden,nochmal buttons.Button
	weiter = buttons.New(1090,650,100,40,255,255,100,false,"  weiter")
	weiter.SetzeFont(font)
	zurueck = buttons.New(850,650,100,40,255,255,100,false,"  zurück")
	zurueck.SetzeFont(font)
	beenden = buttons.New(30,650,100,40,255,255,100,true,"   Ende")
	beenden.SetzeFont(font)
	nochmal = buttons.New(970,650,100,40,255,255,100,false,"nochmal")
	nochmal.SetzeFont(font)
 
 
	// ---------------- Zeichne Spielfeld -------------------------- //
	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",20)
	zeichneSpielfeld(happy,xSize,ilevel,gPunkte,maxPunkte,sk,text)
	zeichneButtons(weiter,zurueck,beenden,nochmal)

	// ---------------- starte Musik ------------------------------- //
	go hintergrundmusik(beenden)

	// ----------- Mausabfrage - Spielsteuerung ---------------------//
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
		if taste==1 && status==1 {
			for id,but:= range sbutton {		// Überprüfe Schalter
				if but.TesteXYPosInButton(mausX,mausY) {
					if sk.GibSchalterwert(id) {
						sk.SchalteSchalterAn(id,false)
					} else {
						sk.SchalteSchalterAn(id,true)
					}
					sk.SchaltungBerechnen()
					nPunkte--
					neuZeichnen = true
				}
			}
			// check Level gewonnen oder verloren?
			if alleLampenAn(sk) {			// check: Level geweonnen?
				inaktiviereSchalter(sbutton)
				// Merke die Punkte im Level
				if nPunkte > ePunkte[ilevel] {ePunkte[ilevel] = nPunkte} // Verbesserung?
				gPunkte = 0					// Berechne Gesamtpunktzahl
				for i:=uint16(0); i<nlevel; i++ {
					gPunkte = gPunkte + ePunkte[i]
				}
				// Wie gewonnen?
				if ilevel+1 == nlevel {		// Letztes Level? => Spiel zu Ende
					text = schreibeGewonnenEnde()
				// oder Prüfung bestanden?
				} else if !bestanden && berechneNote(gPunkte,maxPunkte) != "n.B."{
					bestanden = true
					text = schreibeBestanden()
					weiter.AktiviereButton()
					gfx.SpieleSound(path + "Sounds/Applaus.wav")						
				} else { // oder nur Level gewonnen 
					text = schreibeGewonnen(nPunkte,lev.GibMaxPunktzahl(ilevel))
					weiter.AktiviereButton()
				}
				nochmal.AktiviereButton()
				if ilevel == ilevelGeschafft {ilevelGeschafft++}			
				if !soundAn {
					gfx.SpieleSound(path + "Sounds/Sparkle.wav")
					soundAn = true	
				}
			} else if nPunkte == 0 {	   	// wenn zu viele Versuche!!!
				inaktiviereSchalter(sbutton)
				text = schreibeVerloren()
				happy = false
				nochmal.AktiviereButton()
				if !soundAn {			// Spiele Sound nur einmal
					gfx.SpieleSound(path + "Sounds/GameOver.wav")
					soundAn = true	
				}
			}
			if weiter.TesteXYPosInButton(mausX,mausY) { // nächstes Level
				// Lade nächtes Level
				ilevel++
				zurueck.AktiviereButton()
				// weiter-Button nur wenn Level schon gewonnen
				// und Spiel noch nicht fertig
				if ilevel == ilevelGeschafft || ilevel+1 == nlevel {
					weiter.DeaktiviereButton()
				}
				levelNeuLaden = true
				neuZeichnen = true		
			}
			if nochmal.TesteXYPosInButton(mausX,mausY) { // Level nochmal
				levelNeuLaden = true
				neuZeichnen = true
			}
			if zurueck.TesteXYPosInButton(mausX,mausY) { // Level zurück
				// Lade vorheriges Level 
				ilevel--
				if ilevel == 0 {zurueck.DeaktiviereButton()}
				if ilevel < ilevelGeschafft {weiter.AktiviereButton()}
				levelNeuLaden = true
				neuZeichnen = true
			}
			if beenden.TesteXYPosInButton(mausX,mausY) { // Ende des Spiels
				beenden.DeaktiviereButton()
				gfx.StoppeAlleSounds()
				fmt.Println("Spiel beenden")
				break
			}
			if levelNeuLaden {
				happy = true
				lev  = level.New()		// Veränderungen rückgängig machen
				sk = lev.GibSchaltkreis(ilevel)
				sbutton = makeSchalterbuttonTab(sk,xSize)
				nPunkte = lev.GibMaxPunktzahl(ilevel) + lev.GibMinSchalter(ilevel)
				xSize = lev.GibXSize(ilevel)
				text = lev.GibText(ilevel)
				nochmal.DeaktiviereButton()
				levelNeuLaden = false
				soundAn = false
			}
			if neuZeichnen {
					gfx.UpdateAus()
					zeichneSpielfeld(happy,xSize,ilevel,gPunkte,maxPunkte,sk,text)
					zeichneButtons(weiter,zurueck,beenden,nochmal)
					gfx.UpdateAn()
					neuZeichnen = false
			}
		}
	}

	endbildschirm(ilevel+1, berechneNoteFloat32(gPunkte,maxPunkte), ePunkte, gPunkte)

	return berechneNoteFloat32(gPunkte,maxPunkte),uint32(gPunkte)
	
}



// ---------------  Hilfsfunktionen  -------------------------------//


// Vor: Ein passendes gfx-Fenster (1200x700) ist geöffnet.
// Eff: Der Endbildschirm für das Spiel ist angezeigt und kann mit einem
//		Mausklick auf das Verlassen-Symbol verlassen werden.
func endbildschirm(level uint16, note float32, ePunkte []uint16, gPunkte uint16) {
	
	var path string = ""
	var beenden buttons.Button
	beenden = buttons.New(1100,570,99,129,255,255,100,true,"")

	// Lade Hintergrund
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	gfx.LadeBild(150,80,path + "Bilder/Zertifikat/sprechblase_flipped_400.bmp")
	gfx.LadeBild(20,350,path+"Bilder/Martin/WtheK_black.bmp")
	gfx.LadeBild(620,80,path + "Bilder/Zertifikat/paper_500.bmp")
	gfx.LadeBild(960,520,path + "Bilder/Zertifikat/certified_100.bmp")
	gfx.LadeBild(1100,570,path + "Bilder/Martin/Zurück-Symbol.bmp")
	
	// Ausgabe der Gesamtnote	
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont(path + "Schriftarten/ComputerTypewriter.ttf",80)
	gfx.SchreibeFont(200,10,"Bauelementespiel")
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	gfx.SchreibeFont(295,120,"Sie haben die")
	gfx.SchreibeFont(310,240,"erreicht!")
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	gfx.SchreibeFont(285,150,"Gesamtnote")
	gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
	fmt.Println("Final Level: ",level)
	gfx.SchreibeFont(325,175,fmt.Sprintf("%2.1f",note))

	// Schreibe die Punkte pro Level und Gesamtpunkte
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	for i:=1; i<=len(ePunkte); i++ {
		gfx.SchreibeFont(710,150+uint16((i-1)*30), "Level "+ fmt.Sprint(i) + ":   "+ fmt.Sprint(ePunkte[i-1]) + " Punkte")
	}
	gfx.SchreibeFont(700,140+uint16(len(ePunkte)*30),"----------------------")
	gfx.SchreibeFont(710,160+uint16(len(ePunkte)*30),"Gesamt:    " + fmt.Sprint(gPunkte) + " Punkte")

	// Warte auf Mausklick auf Beenden/Verlassen/Tür-Symbol 
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
		if taste==1 && status==1 {
			if beenden.TesteXYPosInButton(mausX,mausY) {break}
		}
	}

}



// Erg: Ein aktiver Button mit vorgegebenen Position und Größe
//		ist geliefert.
func erzeugeSchalterButton(x,y,xSize uint16) buttons.Button {
	return buttons.New(x-xSize/2,y-xSize/4,xSize,xSize/2,0,0,0,true, "")
}

// Erg: Eine Tabelle aktiver Buttons passend zu den in der Schaltung
//		vorhandenen Schalter ist geliefert.
func makeSchalterbuttonTab(sk sch.Schaltung,xSize uint16) map[uint16]buttons.Button {
	var buts map[uint16]buttons.Button = make(map[uint16]buttons.Button,0)		
	var xb,yb uint16		// Mittelpunkt des Schalters
	var schalterIDs []uint16 = sk.GibSchalterIDs()
	for _,id := range schalterIDs {
		xb,yb = sk.GibPosXY(id)
		buts[id] = erzeugeSchalterButton(xb,yb,xSize)
		buts[id].SetzeSound(path + "Sounds/Punkt.wav")
	}
	return buts
}


// Eff: Alle Buttons zu den Schaltern werden deaktiviert.
func inaktiviereSchalter(buts map[uint16]buttons.Button) {
	for _,button := range buts {
			button.DeaktiviereButton()
	}
}


// Erg: True ist geliefert, wenn alle Lampen der Schaltung
//		angeschaltet sind.
func alleLampenAn(sk sch.Schaltung) bool {
	var ok bool = true
	var status []bool = sk.GibLampenStatus()
	for i:=0; i<len(status); i++ {
		ok = ok && status[i]
	}
	return ok
}


// Erg:_Ein string-Slice für den Willkommenstext und Kurzanleitung
//		ist geliefert.
func WillkommenText() []string {
	var erg []string = make([]string,0)
	erg = append(erg,"Willkommen zur Prüfung der")
	erg = append(erg,"Technischen Informatik!")	
	erg = append(erg,"")
	erg = append(erg,"Ihre Aufgabe ist es die Schalter")
	erg = append(erg,"so zu schalten, dass alle Lampen")
	erg = append(erg,"leuchten. Versuchen Sie die")
	erg = append(erg,"Schalter so wenig wie möglich zu")
	erg = append(erg,"betätigen.")
	erg = append(erg,"")
	erg = append(erg,"Sie benötigen 45% der Punkte, um")
	erg = append(erg,"die Prüfung zu bestehen.")
	erg = append(erg,"")
	erg = append(erg,"Viel Spaß!")
	return erg
}


// Erg:_Ein string-Slice für den Gewonnentext ist geliefert.
func schreibeGewonnen(punkte, punktemax uint16) []string {
	var erg []string = make([]string,0)
	var punktestr string = fmt.Sprint(punkte)
	erg = append(erg,"Glückwunsch Sie haben die")
	erg = append(erg,"Aufgabe geschafft!")
	erg = append(erg,"")
	if punkte == punktemax {
		erg = append(erg,"Sie haben die optimale")
		erg = append(erg,"Lösung gefunden.")
		erg = append(erg,"Es gibt volle "+punktestr+" Punkte.")
	} else {
		erg = append(erg,"Sie haben aber nicht die")
		erg = append(erg,"optimale Lösung gefunden.")
		if punkte > 1 {
			erg = append(erg,"Es gibt nur "+punktestr+" Punkte.")				
		} else {
			erg = append(erg,"Es gibt nur einen Punkt.")
		}
	}
	erg = append(erg,"")	
	erg = append(erg,"Auf zur nächsten Aufgabe oder")
	erg = append(erg,"versuchen Sie es noch einmal.")
	return erg
}


// Erg:_Ein string-Slice für den Bestandentext ist geliefert.
func schreibeBestanden() []string {
	var erg []string = make([]string,0)
	erg = append(erg,"Glückwunsch Sie die Prüfung")
	erg = append(erg,"bestanden!")
	erg = append(erg,"")
	erg = append(erg,"Aber können Sie noch die Note")
	erg = append(erg,"verbessern?")
	return erg
}


// Erg:_Ein string-Slice für den Gewonnentext - alle ist Level geschafft -
//		ist geliefert.
func schreibeGewonnenEnde() []string {
	var erg []string = make([]string,0)
	erg = append(erg,"Glückwunsch Sie haben alle")
	erg = append(erg,"Aufgaben geschafft!")
	erg = append(erg,"")
	erg = append(erg,"Aber können Sie sich noch")
	erg = append(erg,"verbessern?")
	return erg
}


// Erg:_Ein string-Slice für den Verlorentext ist geliefert.
func schreibeVerloren() []string {
	var erg []string = make([]string,0)
	erg = append(erg,"Verloren! Versuchen Sie")
	erg = append(erg,"weniger Schalter zu benutzen.")
	erg = append(erg,"")
	return erg

}


// Erg:_Die Note ist als string geliefert.
func berechneNote(punkte, maxPunkte uint16) string {
	var note string
	var prozente float32 = float32(punkte)/float32(maxPunkte)*100.
	if prozente > 90 {
		note = "1.0"
	} else if prozente > 85 {
		note = "1.3"
	} else if prozente > 80 {
		note = "1.7"
	} else if prozente > 75 {
		note = "2.0"
	} else if prozente > 70 {
		note = "2.3"
	} else if prozente > 65 {
		note = "2.7"
	} else if prozente > 60 {
		note = "3.0"
	} else if prozente > 55 {
		note = "3.3"
	} else if prozente > 50 {
		note = "3.7"
	} else if prozente > 45 {
		note = "4.0"
	} else {
		note = "n.B."
	}
	return note
}


// Erg:_Die Note ist als float32 geliefert.
func berechneNoteFloat32(punkte, maxPunkte uint16) float32 {
	var note float32
	var prozente float32 = float32(punkte)/float32(maxPunkte)*100.
	if prozente > 90 {
		note = 1.0
	} else if prozente > 85 {
		note = 1.3
	} else if prozente > 80 {
		note = 1.7
	} else if prozente > 75 {
		note = 2.0
	} else if prozente > 70 {
		note = 2.3
	} else if prozente > 65 {
		note = 2.7
	} else if prozente > 60 {
		note = 3.0
	} else if prozente > 55 {
		note = 3.3
	} else if prozente > 50 {
		note = 3.7
	} else if prozente > 45 {
		note = 4.0
	} else {
		note = 6.0
	}
	return note
}


// Vor: Ein passendes gfx-Fenster ist geöffnet.
// Eff: Schriebt Note, Punkte und Level ins Spielfeld.
func schreibeSpielstand(level,punkte, maxPunkte uint16) {
	var note string =  berechneNote(punkte, maxPunkte)
	gfx.SchreibeFont(20,15,"Level: " + fmt.Sprint(level))
	gfx.SchreibeFont(150,15,"Punkte: " + fmt.Sprint(punkte))
	gfx.SchreibeFont(320,15,"Note: " + fmt.Sprint(note))
}


// Vor: Ein passendes gfx-Fenster ist geöffnet.
// Eff: Die Buttons sind gezeichnet.
func zeichneButtons(weiter,zurueck,beenden,nochmal buttons.Button) {
	if weiter.GibAktivitaetButton() {
			weiter.ZeichneButton()
	}
	if zurueck.GibAktivitaetButton() {
			zurueck.ZeichneButton()
	}
	if beenden.GibAktivitaetButton() {
			beenden.ZeichneButton()
	}
	if nochmal.GibAktivitaetButton() {
			nochmal.ZeichneButton()
	}
}


// Vor: Ein passendes gfx-Fenster ist geöffnet.
// Eff: Das Spielfeld mit Scjaltplan, Text, WtheK, Buttons und
//	 	Spielstand ist gezeichnet.
func zeichneSpielfeld(happy bool, xSize, ilevel, punkte, maxPunkte uint16, sk sch.Schaltung, text []string) {

	var fontsize int = 20

	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	gfx.Stiftfarbe(0,0,0)		

	if happy {
		gfx.LadeBild(840,10,path + "Bilder/Martin/WtheK_black.bmp")
	} else {
		gfx.LadeBild(840,10,path + "Bilder/Martin/WtheK_black_sad.bmp")
	}
	gfx.Linie(830,0,830,700-1)
	gfx.Linie(830,380,1200-1,380)

	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",fontsize)
	schreibeSpielstand(ilevel+1,punkte,maxPunkte)
	for i:=0; i<len(text); i++ {
		gfx.SchreibeFont(850,400+20*uint16(i),text[i])
	}
	
	sk.Zeichnen(xSize)

}

// Vor: Ein gfx-Grafikfenster ist geöffnet.
// Eff: Hintergrundmusik ist gestartet. (Als go-Routine ausführen
//		damit das Spiel weitergeht.)
func hintergrundmusik(beenden buttons.Button) {
	var soundstr string = "Sounds/Music/bauelemente.wav"
	for beenden.GibAktivitaetButton() {
		gfx.SpieleSound(soundstr)
		time.Sleep (time.Duration(19206e6))
	}
}




