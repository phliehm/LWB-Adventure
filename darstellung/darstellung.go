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


// Methoden
// ---------
	
func MainfloorDarstellen() {
	
	Stiftfarbe(255,255,255)
	Vollrechteck(0,0,1200,700)
	LadeBild(0,50,"./Bilder/mainfloor.bmp")
	Stiftfarbe(0,0,0)
	SetzeFont("./Schriftarten/Starjedi.ttf",80)
	SchreibeFont(105,290,"1")
	SchreibeFont(103,442,"1")
	SetzeFont("./Schriftarten/Starjedi.ttf",50)
	SchreibeFont(968,314,"2")
	SchreibeFont(961,430,"2")
	SetzeFont("./Schriftarten/Starjedi.ttf",25)
	SchreibeFont(436,359,"3")
	SchreibeFont(438,414,"3")
	SetzeFont("./Schriftarten/Starjedi.ttf",20)
	SchreibeFont(725,368,"4")
	SchreibeFont(720,415,"4")
	LadeBildMitColorKey(1083,275, "./Bilder/MainGame/info-1.bmp", 255,255,255)
	//LadeBildMitColorKey (105,325, "./Bilder/MainGame/1.bmp", 255,255,255)
	//LadeBildMitColorKey (965,330, "./Bilder/MainGame/2.bmp", 255,255,255)
	//LadeBildMitColorKey (434,371, "./Bilder/MainGame/3-2.bmp", 255,255,255)
	//LadeBildMitColorKey (725,376, "./Bilder/MainGame/4.bmp", 255,255,255)
	
}

func StartFenster() {
	
	//var startText textboxen.Textbox = textboxen.New(60,110,630,480)
	//var startknopf buttons.Button = buttons.New(600,540,90,50,0,255,0,true," LOS!")
	var startText textboxen.Textbox = textboxen.New(410,110,630,480)
	var startknopf buttons.Button = buttons.New(950,540,90,50,0,255,0,true," LOS!")
	
	LadeBild(0,50,"./Bilder/MainGame/startbildschirm.bmp")
	Stiftfarbe(0,0,0)
	SetzeFont("./Schriftarten/Starjedi.ttf",80)
	SchreibeFont(105,290,"1")
	SchreibeFont(103,442,"1")
	SetzeFont("./Schriftarten/Starjedi.ttf",50)
	SchreibeFont(968,314,"2")
	SchreibeFont(961,430,"2")
	SetzeFont("./Schriftarten/Starjedi.ttf",25)
	SchreibeFont(436,359,"3")
	SchreibeFont(438,414,"3")
	SetzeFont("./Schriftarten/Starjedi.ttf",20)
	SchreibeFont(725,368,"4")
	SchreibeFont(720,415,"4")
	
	//LadeBildMitColorKey(740,250, "./Bilder/MainGame/Darth-1.bmp", 255,255,255)
	LadeBildMitColorKey(140,250, "./Bilder/MainGame/Darth-1.bmp", 255,255,255)
	LadeBildMitColorKey(1083,275, "./Bilder/MainGame/info-1.bmp", 255,255,255)
	
	Stiftfarbe(255,255,255)
	Transparenz(100)
	//Vollrechteck(50,100,650,500)
	Vollrechteck(400,100,650,500)
	Transparenz(0)
	
	startText.SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	//startText.SetzeSchriftgröße()
	//startText.SetzeFarbe(255,255,255)
	startText.SchreibeText("Willkommen zum LWB-Adventure-Game!\n\nBlablablablabla...\nBlablablablabla...\nBlablablablabla...\nBlablablablabla...\n")
	startText.Zeichne()
	startknopf.SetzeFont("./Schriftarten/Ubuntu-B.ttf")
	startknopf.ZeichneButton()
	
	for {
		taste, status, mausX, mausY := MausLesen1()
				
		if taste==1 && status==1 {
			if startknopf.TesteXYPosInButton(mausX,mausY) {
				return
			}
		}
	}
	
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
	
	var infotext textboxen.Textbox = textboxen.New(570,120,500,350)
	var ok buttons.Button = buttons.New(917,295,50,40,0,255,0,true,"OK")
	
	LadeBildMitColorKey(530,90, "./Bilder/MainGame/bubble2_red.bmp", 255,0,0)
	LadeBildMitColorKey(955,390, "./Bilder/MainGame/palimpalim.bmp", 255,255,255)
	
	infotext.SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	infotext.SchreibeText("Info: Blablablablabla...")
	infotext.Zeichne()
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

	var noten []float32 = spielstand.GibNoten()
	var punkte []uint32 = spielstand.GibPunkte()


	Stiftfarbe(255,255,255)
	Vollrechteck(0,0,1200,700)
	LadeBild(250,50,"./Bilder/Tür5.bmp")
		
	fmt.Println("Zeugnis:")
	fmt.Println("--------")
	fmt.Println()
		
	fmt.Println("Semster 1:")
	fmt.Println("Funktionale Programmierung: Note: ",noten[0]," Punkte: ",punkte[0])
	fmt.Println("Technische Informatik: Note: ",noten[1]," Punkte: ",punkte[1])
	fmt.Println()
	fmt.Println("Semster 2:")
	fmt.Println("Imperative und objektorientierte Programmierung: ",noten[3]," Punkte: ",punkte[3])
	fmt.Println("Theoretischen Informatik: ",noten[4]," Punkte: ",punkte[4])
	fmt.Println()
	fmt.Println("Semster 3:")
	fmt.Println("Datenbanksysteme: Note: ",noten[6]," Punkte: ",punkte[6])
	fmt.Println("Fachdidaktik: Note: ",noten[7]," Punkte: ",punkte[7])
	fmt.Println()
	fmt.Println("Semster 4:")
	fmt.Println("Nichtsequentielle und verteilte Programmierung: Note: ",noten[9]," Punkte: ",punkte[9])
	fmt.Println("Softwarepraktikum: Note: ",noten[10]," Punkte: ",punkte[10])
	fmt.Println("Rechnernetze: Note: ",noten[11]," Punkte: ",punkte[11])
	fmt.Println()
	fmt.Println("Gesamtnote: ",durchschnitt(noten))
	fmt.Println("Gesamtpunkte: ",summe(punkte))

}



func Startbildschirm() spielstaende.Spielstand {
	
	//  ------------------   Initialisierung   --------------------  //
	var spielstand spielstaende.Spielstand
	var spielername string
	var laden string
	var path string = "./SAVE"	// Pfad in dem sich die Speicher-
								// standsdateien befinden 
	var punkte []uint32 = make([]uint32,12)		// 12 = maximale Anzahl der Minispiele: 3 Spiele je Semester
	var noten []float32 = make([]float32,12)
	
	
	//  ----------  Beispiel 1: Lade Spielstand  -------------------//	
	fmt.Println("Geben Sie Ihren Spielernamen ein:")
	fmt.Scanln(&spielername)
	
	spielstand = spielstaende.New(spielername,path)
	if spielstand.GibVorhanden() {
		fmt.Println("Es existiert ein alter Spielstand.")
		for laden != "W" {
			fmt.Println("(W)eiterspielen oder (U)eberschreiben?")
			fmt.Scanln(&laden)
			if laden == "U" {
				spielstand.Speichern(noten,punkte)
				break
			} else {
				fmt.Println("Eingabe ungültig!")
				fmt.Println()				
			}
		}
	} else {
		spielstand.Speichern(noten,punkte)		// Grundzustand
	}
	
	return spielstand
	
}


/*
//Hilfsfunktion BauelementeSpiel:
func bauelemente() {
	var ilevel uint16	  			// aktuelle Levelnummer
	var ePunkte []uint16			// Punkte erreicht im Level

	//ilevel = 3	  								// aktuelle Levelnummer
	//ePunkte = []uint16{3,3,3,0,0,0,0} 			// Punkte erreicht im Level

	ilevel = 0
	ePunkte = []uint16{}
	bauelementeSpiel.BauelementeSpiel(ilevel,ePunkte)
}
*/

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
			note, punkte = vadeROBIgame.Vaderobi()
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
