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


// interne Hilfsfunktionen
// ------------------------


func ladeEndeBildschirm() {										//TODO
	Vollrechteck(0,0,1200,700)
	TastaturLesen1()
}


func schreibeZertifikat(spielstand spielstaende.Spielstand) {

	//var gamestxt,notentxt string
	var zertifikatgames textboxen.Textbox = textboxen.New(630,200,450,550)
	var zertifikatnoten textboxen.Textbox = textboxen.New(1000,200,100,550)
	var gametitel []string = gametitelSchreiben()
	var gamenoten []float32 = ordneNotenGamesZu(spielstand)
	var gamestxt,notentxt string				// Texte für das Zertifikat

	for i:=0; i<len(gamenoten); i++ {
		gamestxt = gamestxt + gametitel[i] + ":\n"
		notentxt= notentxt + "" +fmt.Sprint(gamenoten[i]) + "\n"
	}

	// Tabellenkopf Zertifikat
	SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	SchreibeFont(630,150,"Veranstaltung                   Noten")

	// Spieletitel schreiben
	zertifikatgames.SetzeFarbe(0,0,0)
	zertifikatgames.SetzeZeilenAbstand(5)
	zertifikatgames.SetzeSchriftgröße(22)
	zertifikatgames.SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	zertifikatgames.SchreibeText(gamestxt)
	zertifikatgames.Zeichne()

	// Noten schreiben
	zertifikatnoten.SetzeFarbe(0,0,0)
	zertifikatnoten.SetzeZeilenAbstand(5)
	zertifikatnoten.SetzeSchriftgröße(22)
	zertifikatnoten.SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	zertifikatnoten.SchreibeText(notentxt)
	zertifikatnoten.Zeichne()

}


func gametitelSchreiben() []string {
	
	var gametitel []string = make([]string,9)

	// Reihenfolge wie im Aufruf im MinigameLaden
	gametitel[0] = "Bauelemente-Spiel (RS)"
	gametitel[1] = "Mustererkennung (FP)"
	gametitel[2] = "Super-ALP2-Escape (ALP2)"
	gametitel[3] = "Getränkeautomaten-Spiel (EthI)"
	gametitel[4] = "SQL-Quest (DBSA)"
	gametitel[5] = "Didaktik-Game (DDI)"
	gametitel[6] = "Food-Moorhuhn (NSP)"
	gametitel[7] = "Bug Attack (SWP)"
	gametitel[8] = "TheNETgame (NET)"

/*	gametitel[1] = "Bauelemente-Spiel (RS)"
	gametitel[2] = "Mustererkennung (FP)"
	gametitel[3] = "Super-ALP2-Escape (ALP2)"
	gametitel[4] = "Getränkeautomaten-Spiel (EthI)"
	gametitel[5] = "Didaktik-Game (DDI)"
	gametitel[6] = "SQL-Quest (DBSA)"
	gametitel[7] = "Bug Attack (SWP)"
	gametitel[8] = "Food-Moorhuhn (NSP)"
	gametitel[9] = "TheNETgame (NET)"		*/

	return gametitel
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


// Eff: Der Endbildschirm mit Zertifikat und Durchschnitt ist angezeigt. 
func EndbildschirmDarstellen(spielstand spielstaende.Spielstand) {

	var exit vierecke.Viereck = vierecke.New(1100,565,1190,565,1190,685,1100,685)   	// Position wie in den anderen Räumen
	// var exit vierecke.Viereck = vierecke.New(1080,30,1080,145,1170,145,1170,30)
	var noten []float32 = spielstand.GibNoten()

	Stiftfarbe(255,255,255)
	Cls()
	
	// Hintergrund gestalten
	LadeBild(150,100,"./Bilder/sprechblase_flipped_400.bmp")
	LadeBildMitColorKey(100,350,"./Bilder/Darth_200.bmp",255,255,255)
	LadeBild(600,80,"./Bilder/MainGame/zertifikat.bmp")
	LadeBild(940,510,"./Bilder/certified_100.bmp")

	// exit-Schalter einfügen
	LadeBild(1100,565,"./Bilder/Zurück-Symbol.bmp")
	exit.SetzeFarbe(0,0,0)
	exit.Zeichnen()
	exit.AktiviereKlickbar()
	
	
	Stiftfarbe(0,0,0)
	//  Titel schreiben
	SetzeFont("./Schriftarten/brlnsr.ttf",35)
	SchreibeFont(50,20,"Herzlichen Glückwunsch zum erfolgreich absolvierten LWB-Adventure!!!")
	
	// Notendurchschnitt schreiben
	SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	SchreibeFont(295,145,"Du hast den")
	SchreibeFont(310,265,"erreicht!")
	SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	SchreibeFont(230,175,"Notendurchschnitt")
	SetzeFont("./Schriftarten/Starjedi.ttf",42)
	SchreibeFont(325,200,fmt.Sprintf("%2.1f",durchschnitt(noten)))
		
	// Inhalt des Zertifikates vorbereiten	
	schreibeZertifikat(spielstand)

/*
	// Spiel beenden?
	Stiftfarbe(255,100,0)
	Vollrechteck(350,620,200,50)
	Stiftfarbe(0,0,0)
	Rechteck(350,620,200,50)
	Stiftfarbe(0,0,0)
	SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	SchreibeFont(380,632,"Spiel beenden")
*/	

	
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


// Erg: Eine Liste von Noten passend zu den 9 Spielen ist geliefert.
// Anpassung nötig, da 4*3=12 Spielstände möglich!
func ordneNotenGamesZu(spielstand spielstaende.Spielstand) []float32 {
	var gamenoten []float32 = make([]float32,9)
	var noten = spielstand.GibNoten()
	
	gamenoten[0] = noten[0]
	gamenoten[1] = noten[1]
	gamenoten[2] = noten[3]
	gamenoten[3] = noten[4]
	gamenoten[4] = noten[6]
	gamenoten[5] = noten[7]
	gamenoten[6] = noten[9]
	gamenoten[7] = noten[10]
	gamenoten[8] = noten[11]
	
	return gamenoten

}
