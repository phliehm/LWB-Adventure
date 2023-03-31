// ADT bauelemente - Spezifikation und Implementation

// Martin Seiß    21.3.2023


package main

import sch "./schaltungen"
import b "./bauelemente"
import "fmt"
import "gfx"
import "../../Klassen/buttons"


func erzeugeSchalterButton(x,y,xSize uint16) buttons.Button {
	return buttons.New(x-xSize/2,y-xSize/4,xSize,xSize/2,0,0,0,true, "")
}


func makeSchalterbuttonTab(sk sch.Schaltung,xSize uint16) map[uint16]buttons.Button {
	var buts map[uint16]buttons.Button = make(map[uint16]buttons.Button,0)		
	var xb,yb uint16		// Mittelpunkt des Schalters
	var schalterIDs []uint16 = sk.GibSchalterIDs()
	for _,id := range schalterIDs {
		if sk.GibBauelementtyp(id) == b.Schalter {
			xb,yb = sk.GibPosXY(id)
			// unsichbaren Button zuordnen
			buts[id] = erzeugeSchalterButton(xb,yb,xSize)	
		}
	}
	return buts
}


func WillkommenText() []string {
	var erg []string = make([]string,0)
	erg = append(erg,"Willkommen beim Schaltungsspiel!")
	erg = append(erg,"")
	erg = append(erg,"Deine Aufgabe ist es die Schalter")
	erg = append(erg,"so zu schalten, dass alle Lampen")
	erg = append(erg,"leuchten. Versuche die Schalter")
	erg = append(erg,"so wenig wie möglich zu betätigen.")
	erg = append(erg,"")
	erg = append(erg,"Viel Spaß!")
	return erg
}


func schreibeSpielstand(level,punkte uint16, note string) {
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


func zeichneSpielfeld(happy bool, xSize uint16, sk sch.Schaltung, text []string) {

	var fontsize int = 20

//	gfx.UpdateAus()

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
	schreibeSpielstand(3,128,"n.B.")
	for i:=0; i<len(text); i++ {
		gfx.SchreibeFont(850,400+20*uint16(i),text[i])
	}
	
	
	sk.Zeichnen(xSize)

//	gfx.UpdateAn()
}



func main() {

	var xSize uint16 = 100  		// Größe des Bauelements in x-Richtung
									// in Pixel zur Skalierung
	var happy bool = true
									
	var text []string = WillkommenText()
	var font string = "../../Schriftarten/Ubuntu-B.ttf"

	// ---------------- Erzeuge Schaltkreis -------------- //

	var sk sch.Schaltung = sch.New()
	
	sk.BauteilEinfuegen(1,100,100,b.Schalter)				// Schalter einbauen
	sk.BauteilEinfuegen(7,100,400,b.Schalter)
	
	
	sk.BauteilEinfuegen(3,400,250,b.NOT)	
	
	sk.BauteilEinfuegen(4,700,100,b.Lampe)
	sk.BauteilEinfuegen(5,700,400,b.Lampe)

	sk.VerbindungEinfuegen(1,3,1,250)
	sk.VerbindungEinfuegen(1,4,1,550)
	sk.VerbindungEinfuegen(7,3,2,250)
	sk.VerbindungEinfuegen(3,5,1,550)

	sk.SchaltungBerechnen()


	//  --------------------   Buttons ------------------------------//

	// erzeuge eine Tabelle von Buttons zu den zugehörigen Schaltern //
	// id gibt die Zuordnung
	var sbutton map[uint16]buttons.Button = makeSchalterbuttonTab(sk,xSize)

	var weiter,zurueck,beenden,nochmal buttons.Button
	weiter = buttons.New(1090,650,100,40,255,255,100,true,"  weiter")
	weiter.SetzeFont(font)
	zurueck = buttons.New(850,650,100,40,255,255,100,false,"  zurück")
	zurueck.SetzeFont(font)
	beenden = buttons.New(30,650,100,40,255,255,100,true,"   Ende")
	beenden.SetzeFont(font)
	nochmal = buttons.New(970,650,100,40,255,255,100,false,"nochmal")
	nochmal.SetzeFont(font)
 
	// ---------------- Zeichne Spielfeld -------------- //

	gfx.Fenster(1200,700)
	//gfx.SetzeFont ("../../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",int(xSize/10))
	gfx.SetzeFont ("../../Schriftarten/Ubuntu-B.ttf",20)
	//gfx.SetzeFont ("../../Schriftarten/CollegiateBlackFLF.ttf",int(xSize/10))

	zeichneSpielfeld(happy,xSize,sk,text)
	zeichneButtons(weiter,zurueck,beenden,nochmal)
//	weiter.ZeichneButton()
//	zurueck.ZeichneButton()
//	beenden.ZeichneButton()
//	nochmal.ZeichneButton()

	
	// ---------------- Teste Schaltkreis -------------- //


	// Mausabfrage
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
		if taste==1 && status==1 {
			for id,but:= range sbutton {		// Überprüfe Schalter
				if but.TesteXYPosInButton(mausX,mausY) {
					if sk.GibSchalterwert(id) {
						sk.SchalteLampeAn(id,false)
					} else {
						sk.SchalteLampeAn(id,true)
					}
					sk.SchaltungBerechnen()
					gfx.UpdateAus()
					zeichneSpielfeld(happy,xSize,sk,text)
					zeichneButtons(weiter,zurueck,beenden,nochmal)
					gfx.UpdateAn()
				}
			}
			if beenden.TesteXYPosInButton(mausX,mausY) { // Ende?
				break
			}
		}
	}

	
}
