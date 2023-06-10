// Autor: A. Cyriacus und M. Seiß
// Datum: 07.06.2023
// Zweck: Implementierung des ADO darstellung

package darstellung

/* Der ADO vdarstellung übernimmt die Darstellung der Objekte.
 * Die Darstellung wird vom passenden Controller aufgerufen.
 */

import (
	. "gfx"
	"../MiniGames/1_FP"
	"../MiniGames/bauelementeSpiel"
	"../MiniGames/2_ALP2/vaderobigame"
//	"../MiniGames/2_ALP2/vadeROBIgame"
	//"../MiniGames/3_DDI"
	"../MiniGames/3_DBSA"
	//"../MiniGames/4_BugAttack"
	"../MiniGames/4_Moorhuhn"
	"../MiniGames/theNETgame"
	"../Klassen/spielstaende"
	"../Klassen/textboxen"
	"../Klassen/buttons"
	"../Klassen/vierecke"
	"fmt"
)

// Globale Variablen
// ------------------

var gametitel []string = make([]string,9)


// interne Hilfsfunktionen
// ------------------------

func ladeEndeBildschirm() {										//TODO
	Vollrechteck(0,0,1200,700)
	TastaturLesen1()
}

func gametitelSchreiben() {										//Bitte überprüfen, ob die Titel zur Noten-Speicherzelle passen!
	gametitel[1] = "Bauelemente-Spiel (RS)"
	gametitel[2] = "Mustererkennung (FP)"
	gametitel[3] = "Super-ALP2-Escape (ALP2)"
	gametitel[4] = "Getränkeautomaten-Spiel (EthI)"
	gametitel[5] = "Didaktik-Game (DDI)"
	gametitel[6] = "SQL-Quest (DBSA)"
	gametitel[7] = "Bug Attack (SWP)"
	gametitel[8] = "Food-Moorhuhn (NSP)"
	gametitel[9] = "TheNETgame (NET)"
}


// Methoden
// ---------
	
func MainfloorDarstellen() {
	
	Stiftfarbe(255,255,255)
	Vollrechteck(0,0,1200,700)
	LadeBild(0,50,"./Bilder/mainfloor.bmp")
	Stiftfarbe(0,0,0)
	SetzeFont("./Schriftarten/Starjedi.ttf",80)
	SchreibeFont(105,290,"1")
	//SchreibeFont(103,442,"1")
	SetzeFont("./Schriftarten/Starjedi.ttf",50)
	SchreibeFont(968,314,"2")
	//SchreibeFont(961,430,"2")
	SetzeFont("./Schriftarten/Starjedi.ttf",25)
	SchreibeFont(436,359,"3")
	//SchreibeFont(438,414,"3")
	SetzeFont("./Schriftarten/Starjedi.ttf",20)
	SchreibeFont(725,368,"4")
	//SchreibeFont(720,415,"4")
	LadeBildMitColorKey(1083,275, "./Bilder/MainGame/info-1.bmp", 255,255,255)
	//LadeBildMitColorKey (105,325, "./Bilder/MainGame/1.bmp", 255,255,255)
	//LadeBildMitColorKey (965,330, "./Bilder/MainGame/2.bmp", 255,255,255)
	//LadeBildMitColorKey (434,371, "./Bilder/MainGame/3-2.bmp", 255,255,255)
	//LadeBildMitColorKey (725,376, "./Bilder/MainGame/4.bmp", 255,255,255)
	
}

func SemesterraumDarstellen(n int) {
	
	switch n {
		case 0:
		ladeEndeBildschirm()
		case 1:
		LadeBild(0,0,"./Bilder/MainGame/raum1.bmp")
		case 2:
		LadeBild(0,0,"./Bilder/MainGame/raum2.bmp")
		case 3:
		LadeBild(0,0,"./Bilder/MainGame/raum3.bmp")
		case 4:
		LadeBild(0,0,"./Bilder/MainGame/raum4.bmp")
		case 5:
		panic("Sollte nicht passieren!")
		
	}
	
}

func InfoDarstellen() {
	
	var infotexthead textboxen.Textbox = textboxen.New(570,120,500,350)
	var infotext textboxen.Textbox = textboxen.New(570,160,500,310)
	var ok buttons.Button = buttons.New(917,295,50,40,0,255,0,true,"OK")
	
	LadeBildMitColorKey(530,90, "./Bilder/MainGame/bubble2_red.bmp", 255,0,0)
	LadeBildMitColorKey(955,390, "./Bilder/MainGame/palimpalim.bmp", 255,255,255)
	
	infotext.SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	infotext.SchreibeText("Info: Blablablablabla...")									//TODO
	infotext.Zeichne()
	infotexthead.SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	infotexthead.SetzeSchriftgröße(35)
	infotexthead.SchreibeText("PALIMPALIM!!!")
	infotexthead.Zeichne()
	ok.SetzeFont("./Schriftarten/Ubuntu-B.ttf")
	ok.ZeichneButton()
	
	for {
		taste, status, mausX, mausY := MausLesen1()
				
		if taste==1 && status==1 {
			if ok.TesteXYPosInButton(mausX,mausY) {
				return
			}
		}
	}
	
}

func EndbildschirmDarstellen(spielstand spielstaende.Spielstand) {

	var exit vierecke.Viereck = vierecke.New(1080,30,1080,145,1170,145,1170,30)
	var zertifikatinhalt textboxen.Textbox = textboxen.New(650,150,480,550)
	var noten []float32 = spielstand.GibNoten()
	//var punkte []uint32 = spielstand.GibPunkte()

	Stiftfarbe(255,255,255)
	Cls()
	
	//LadeBild(250,50,"./Bilder/Tür5.bmp")
	
	LadeBild(150,100,"./Bilder/sprechblase_flipped_400.bmp")
	LadeBildMitColorKey(100,350,"./Bilder/Darth_200.bmp",255,255,255)
	LadeBild(600,80,"./Bilder/MainGame/zertifikat.bmp")
	LadeBild(940,510,"./Bilder/certified_100.bmp")
	LadeBild(1080,30,"./Bilder/Zurück-Symbol.bmp")
	exit.SetzeFarbe(0,0,0)
	exit.Zeichnen()
	exit.AktiviereKlickbar()
	
	Stiftfarbe(0,0,0)
	SetzeFont("./Schriftarten/brlnsr.ttf",42)
	SchreibeFont(50,10,"Herzlichen Glückwunsch zum erfolgreich absolvierten LWB-Adventure!!!")
	Stiftfarbe(0,0,0)
	SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	SchreibeFont(295,145,"Du hast den")
	SchreibeFont(310,265,"erreicht!")
	SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	SchreibeFont(230,175,"Notendurchschnitt")
	SetzeFont("./Schriftarten/Starjedi.ttf",42)
	
	/*
	var notensumme,abschlussnote float32
	for i:=0; i<len(noten); i++ {
		notensumme = notensumme + noten[i]
	}
	abschlussnote = notensumme/float32(len(noten))
	*/
	
	SchreibeFont(325,200,fmt.Sprintf("%2.1f",durchschnitt(noten)))
	
	//SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	for i:=0; i<len(noten); i++ {
		
		zertifikatinhalt.SetzeFarbe(0,0,0)
		zertifikatinhalt.SetzeZeilenAbstand(5)
		zertifikatinhalt.SetzeSchriftgröße(22)
		zertifikatinhalt.SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
		zertifikatinhalt.SchreibeText(gametitel[i]+":   "+fmt.Sprintf("%2.1f",noten[i]))
		zertifikatinhalt.Zeichne()
		
		//SchreibeFont(650,150+uint16((i-1)*68),veranstaltungstitel[i] + ":   "fmt.Sprintf("%2.1f",noten[i]))
	}
	
	for {
		taste, status, mausX, mausY := MausLesen1()
		if taste==1 && status==1 {
			if exit.Angeklickt(mausX,mausY) { 							// Ende des Spiels
				fmt.Println("exit geklickt")
				break
			}
		}
	}
}


// Erg: Die erspielte Note und die Punkte sind geliefert.
func MinigameLaden(raum,n int) (note float32, punkte uint32){
	
	switch raum {
		
		case 1:
		switch n {
			case 1:
			// note, punkte = muster.Muster()
			// noch falsche Rückgabe
			note = 6 
			punkte = uint32(muster.Muster())
			case 2:
			note, punkte = bauelementeSpiel.BauelementeSpiel()
		}
		
		case 2:
		switch n {
			case 1:
			note, punkte = vaderobigame.Vaderobi()
			case 2:
			//TODO (noch kein importierbares package vorhanden)
		}
		
		case 3:
		switch n {
			case 1:
			sqlGame.SQLgame()		// Ausgabe fehlt noch
			note, punkte = 6,0
			//note, punkte = sqlGame.SQLgame()
			case 2:
			//TODO (noch kein importierbares package vorhanden)
		}
		
		case 4:
		switch n {
			case 1:
			note, punkte = moorhuhn.Moorhuhn()
			case 2:
			//TODO (noch kein importierbares package vorhanden)
			case 3:
			note, punkte = theNETgame.TheNETgame()
		}
					
	}
	
	return
	
}



// --------------    Hilfsfunktionen ------------------------//


// Erg: Die Durchschnittsnote ist geliefert. Alle Prüfungen, die
//		noch nicht abgelegt wurden, sind nicht berücksichtigt.
func durchschnitt(xs []float32) float32 {

	var erg float32
	var anzahl float32

	for i:=0; i<len(xs); i++ {
		if xs[i] != 0 {
			erg = erg + xs[i]
			anzahl++
		}
	}
	
	if anzahl > 0 {
		erg = erg/anzahl
	} else {
		erg = 6
	}
	
	return erg

}


// Erg: Die Summe ist geliefert.
func summe(ns []uint32) uint32 {

	var erg uint32
		
	for i:=0; i<len(ns); i++ {
		erg = erg + ns[i]  
	}
	
	return erg

}
