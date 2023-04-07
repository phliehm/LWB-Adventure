// ADT bauelemente - Spezifikation und Implementation

// Martin Seiß    21.3.2023


package main

import sch "./schaltungen"
//import b "./bauelemente"
import level "./beLevel"
import "fmt"
import "gfx"
import "../../Klassen/buttons"
import "os"
import "strconv"


func erzeugeSchalterButton(x,y,xSize uint16) buttons.Button {
	return buttons.New(x-xSize/2,y-xSize/4,xSize,xSize/2,0,0,0,true, "")
}


func makeSchalterbuttonTab(sk sch.Schaltung,xSize uint16) map[uint16]buttons.Button {
	var buts map[uint16]buttons.Button = make(map[uint16]buttons.Button,0)		
	var xb,yb uint16		// Mittelpunkt des Schalters
	var schalterIDs []uint16 = sk.GibSchalterIDs()
	for _,id := range schalterIDs {
//		if sk.GibBauelementtyp(id) == b.Schalter {
			xb,yb = sk.GibPosXY(id)
			// unsichbaren Button zuordnen
			buts[id] = erzeugeSchalterButton(xb,yb,xSize)	
//		}
	}
	return buts
}


func inaktiviereSchalter(buts map[uint16]buttons.Button) {
	for _,button := range buts {
			button.DeaktiviereButton()
	}
}


func alleLampenAn(sk sch.Schaltung) bool {
	var ok bool = true
	var status []bool = sk.GibLampenStatus()
	for i:=0; i<len(status); i++ {
		ok = ok && status[i]
	}
	return ok
}


func WillkommenText() []string {
	var erg []string = make([]string,0)
	erg = append(erg,"Willkommen zur Prüfung der")
	erg = append(erg,"Technischen Informatik!")	
	erg = append(erg,"")
	erg = append(erg,"Ihre Aufgabe ist es die Schalter")
	erg = append(erg,"so zu schalten, dass alle Lampen")
	erg = append(erg,"leuchten. Versuchen Sie die Schalter")
	erg = append(erg,"so wenig wie möglich zu betätigen.")
	erg = append(erg,"")
	erg = append(erg,"Viel Spaß!")
	return erg
}


func schreibeGewonnen() []string {
	var erg []string = make([]string,0)
	erg = append(erg,"Glückwunsch Sie haben alle")
	erg = append(erg,"Aufgaben geschafft!")
	erg = append(erg,"")
	erg = append(erg,"Auf zur nächsten Aufgabe oder")
	erg = append(erg,"versuchen Sie es noch einmal.")
	return erg
}


func schreibeGewonnenEnde() []string {
	var erg []string = make([]string,0)
	erg = append(erg,"Glückwunsch Sie haben alle")
	erg = append(erg,"Aufgaben geschafft!")
	erg = append(erg,"")
	erg = append(erg,"Aber können Sie sich noch")
	erg = append(erg,"verbessern?")
	return erg
}


func schreibeVerloren() []string {
	var erg []string = make([]string,0)
	erg = append(erg,"Verloren! Versuchen Sie")
	erg = append(erg,"weniger Schalter zu benutzen.")
	erg = append(erg,"")
	return erg

}


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


func schreibeSpielstand(level,punkte, maxPunkte uint16) {
	var note string =  berechneNote(punkte, maxPunkte)
	gfx.SchreibeFont(20,15,"Level: " + fmt.Sprint(level))
	gfx.SchreibeFont(150,15,"Punkte: " + fmt.Sprint(punkte))
	gfx.SchreibeFont(320,15,"Note: " + fmt.Sprint(note))
}


func zeichneButtons(weiter,zurueck,beenden,nochmal buttons.Button) {
//	gfx.UpdateAus()
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
//	gfx.UpdateAn()
}


func zeichneSpielfeld(happy bool, xSize, ilevel, punkte, maxPunkte uint16, sk sch.Schaltung, text []string) {

	var fontsize int = 20

	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	gfx.Stiftfarbe(0,0,0)		

	if happy {
		gfx.LadeBild(840,10,"../../Bilder/WtheK_black.bmp")
	} else {
		gfx.LadeBild(840,10,"../../Bilder/WtheK_black_sad.bmp")
	}
	gfx.Linie(830,0,830,700-1)
	gfx.Linie(830,380,1200-1,380)

	gfx.SetzeFont ("../../Schriftarten/Ubuntu-B.ttf",fontsize)
	schreibeSpielstand(ilevel+1,punkte,maxPunkte)
	for i:=0; i<len(text); i++ {
		gfx.SchreibeFont(850,400+20*uint16(i),text[i])
	}
	
	
	sk.Zeichnen(xSize)

}



func main() {

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
									
	var text []string = WillkommenText()
	var font string = "../../Schriftarten/Ubuntu-B.ttf"


	// -------    Lade Level gegeben auf der Kommandozeile  ------- //
	if len(os.Args) > 1 {
		intVar, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("Levelargument falsch!")
		}		
		ilevel = uint16(intVar) - 1
	}


	// ---------------- Erzeuge Schaltkreis ----------------------- //
	// ---------------- und lade Level 1 ---------------------------//
	var lev level.Level = level.New()
	var sk sch.Schaltung = lev.GibSchaltkreis(ilevel)
	var xSize uint16 = lev.GibXSize(ilevel) 		
							// Größe des Bauelements in x-Richtung
							// in Pixel zur Skalierung
	nlevel = lev.AnzahlLevel()
	ePunkte = make([]uint16,nlevel)
	sk.SchaltungBerechnen()
	nPunkte = lev.GibMaxPunktzahl(ilevel) + lev.GibMinSchalter(ilevel)
	// Zähle maximale erreichbare Punktzahl
	for i:=uint16(0); i<nlevel; i++ { 
		maxPunkte = maxPunkte + lev.GibMaxPunktzahl(i)
	}
	// fmt.Println("maxPunkte: ",maxPunkte)


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
 
 
	// ---------------- Zeichne Spielfeld -------------- //

	gfx.Fenster(1200,700)
	gfx.SetzeFont ("../../Schriftarten/Ubuntu-B.ttf",20)

	zeichneSpielfeld(happy,xSize,ilevel,gPunkte,maxPunkte,sk,text)
	zeichneButtons(weiter,zurueck,beenden,nochmal)


	// Mausabfrage - Spielsteuerung
	for {
		// neuZeichnen = false
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
			if alleLampenAn(sk) {
				inaktiviereSchalter(sbutton)
				if ilevel+1 == nlevel {		// Spiel zu Ende?
					text = schreibeGewonnenEnde()
				} else {
					text = schreibeGewonnen()
					weiter.AktiviereButton()
				}
				nochmal.AktiviereButton()
				neuZeichnen = true
				if ilevel == ilevelGeschafft {ilevelGeschafft++}
				// Merke die Punkte im Level und berechne die Gesamtpunktzahl
				if nPunkte > ePunkte[ilevel] {ePunkte[ilevel] = nPunkte}
				gPunkte = 0
				for i:=uint16(0); i<nlevel; i++ {
					gPunkte = gPunkte + ePunkte[i]
					//fmt.Println(gPunkte)
				}
			} else if nPunkte == 0 {	   // wenn zu viele Versuche!!!
				inaktiviereSchalter(sbutton)
				text = schreibeVerloren()
				happy = false
				nochmal.AktiviereButton()
				neuZeichnen = true
			}
			if weiter.TesteXYPosInButton(mausX,mausY) {
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
			if nochmal.TesteXYPosInButton(mausX,mausY) {
				levelNeuLaden = true
				neuZeichnen = true
			}
			if zurueck.TesteXYPosInButton(mausX,mausY) {
				// Lade vorheriges Level 
				ilevel--
				if ilevel == 0 {zurueck.DeaktiviereButton()}
				if ilevel < ilevelGeschafft {weiter.AktiviereButton()}
				levelNeuLaden = true
				neuZeichnen = true
			}
			if levelNeuLaden {
				happy = true
				lev  = level.New()	// Veränderungen rückgängig machen
				sk = lev.GibSchaltkreis(ilevel)
				sbutton = makeSchalterbuttonTab(sk,xSize)
				nPunkte = lev.GibMaxPunktzahl(ilevel) + lev.GibMinSchalter(ilevel)
				xSize = lev.GibXSize(ilevel)
				text = lev.GibText(ilevel)
				nochmal.DeaktiviereButton()
				levelNeuLaden = false
			}
			if neuZeichnen {
//					fmt.Println("Schalter 1: ",sk.GibSchalterwert(1))
//					fmt.Println("Schalter 2: ",sk.GibSchalterwert(2))
					gfx.UpdateAus()
					zeichneSpielfeld(happy,xSize,ilevel,gPunkte,maxPunkte,sk,text)
					zeichneButtons(weiter,zurueck,beenden,nochmal)
					gfx.UpdateAn()
					neuZeichnen = false
			}
			if beenden.TesteXYPosInButton(mausX,mausY) { // Ende?
				break
			}
		}
	}

	
}