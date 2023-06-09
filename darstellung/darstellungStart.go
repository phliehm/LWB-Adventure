// Autor: A. Cyriacus und M. Seiß
// Datum: 08.06.2023
// Zweck: Implementierung des ADO darstellung
// Hier nur Startbildschirm

package darstellung

/* Der ADO vdarstellung übernimmt die Darstellung der Objekte.
 * Die Darstellung wird vom passenden Controller aufgerufen.
 */

import (
	. "gfx"
//	"../MiniGames/2_ALP2/vadeROBIgame"
	//"../MiniGames/3_DDI"
	//"../MiniGames/4_BugAttack"
	"../Klassen/spielstaende"
	"../Klassen/textboxen"
	"../Klassen/buttons"
	"fmt"
	"../Klassen/felder"
)


func StartFenster() spielstaende.Spielstand {

	//var startText textboxen.Textbox = textboxen.New(60,110,630,480)
	//var startknopf buttons.Button = buttons.New(600,540,90,50,0,255,0,true," LOS!")

	// für Steuerung
	var startText textboxen.Textbox = textboxen.New(410,110,630,480)
	var startknopf buttons.Button = buttons.New(950,540,90,50,0,255,0,false," LOS!")
	var alt buttons.Button = buttons.New(480,540,210,50,0,255,0,false,"  Weiterspielen")
	var neu buttons.Button = buttons.New(780,540,210,50,0,255,0,false," Überschreiben")
	var eingabefeld felder.Feld = felder.New (410,  190, 50, 'l', "")

	// für Spielstand
	var spielstand spielstaende.Spielstand
	var spielername string
	var path string = "./SAVE"	// Pfad in dem sich die Speicher-
								// standsdateien befinden 
	var punkte []uint32 = make([]uint32,12)		// 12 = maximale Anzahl der Minispiele: 3 Spiele je Semester
	var noten []float32 = make([]float32,12)
	

	
	zeichneStartHintergrund()
	startText.SetzeFont("./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	startText.SchreibeText("Willkommen zum LWB-Adventure-Game!\n\nSchreiben Sie sich ein und das Abenteuer kann beginnen: \n")
	startText.Zeichne()

	// Eingabe des Namens
	spielername = eingabefeld.Edit ()
	fmt.Println("spielername: ",spielername)
	
	spielstand = spielstaende.New(spielername,path)
	
	if spielstand.GibVorhanden() {
		zeichneStartHintergrund() 
		startText.SchreibeText("Es existiert ein alter Spielstand.\n\n"+
			"Weiterspielen oder Überschreiben?")
		startText.Zeichne()
		alt.AktiviereButton()
		alt.SetzeFont("./Schriftarten/Ubuntu-B.ttf")
		alt.ZeichneButton()
		neu.AktiviereButton()
		neu.SetzeFont("./Schriftarten/Ubuntu-B.ttf")
		neu.ZeichneButton()
		// Warte auf Tastendruck
		for {
			taste, status, mausX, mausY := MausLesen1()	
			if taste==1 && status==1 {
				if alt.TesteXYPosInButton(mausX,mausY) { // behalten
					break
				}
				if neu.TesteXYPosInButton(mausX,mausY) { // verwerfen
					spielstand.Speichern(noten,punkte)
					break
				}
			}
		}
	} else {
		// neuen Spielstand initialisieren
		spielstand.Speichern(noten,punkte)	
	}
	
	
	zeichneStartHintergrund()
	startText.SchreibeText("Klicken Sie sich einfach durch die Welt und" +
		" bestehen Sie die Prüfungen der LWB Informatik. \n \n" +
		" Durch die Türen können Sie neue Räume betreten," +
		" aber erst wenn Sie eine Freigabe durch bestandene " +
		" Prüfungen erhalten haben. Wenn Sie eine Prüfung ablegen" + 
		" wollen, klicken Sie auf einen der Dozenten." + 
		" Beginnen Sie im offenen Raum 1. \n \n" + 
		" Viel Spaß!")
	startText.Zeichne()

	// Startknopf aktivieren
	startknopf.AktiviereButton()
	startknopf.SetzeFont("./Schriftarten/Ubuntu-B.ttf")
	startknopf.ZeichneButton()
	// Startknopf abfragen
	for {
		taste, status, mausX, mausY := MausLesen1()	
		if taste==1 && status==1 {
			if startknopf.TesteXYPosInButton(mausX,mausY) {
				break
			}
		}
	}

	return spielstand
	
}



/*
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
			if laden == "W" {
				break
			} else if laden == "U" {
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
*/

// -------------   Hilfsfunktionen  -------------------------//

func zeichneStartHintergrund() {
	
	LadeBild(0,50,"./Bilder/MainGame/startbildschirm.bmp")
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
	
	//LadeBildMitColorKey(740,250, "./Bilder/MainGame/Darth-1.bmp", 255,255,255)
	LadeBildMitColorKey(140,250, "./Bilder/MainGame/Darth-1.bmp", 255,255,255)
	LadeBildMitColorKey(1083,275, "./Bilder/MainGame/info-1.bmp", 255,255,255)
	
	Stiftfarbe(255,255,255)
	Transparenz(100)
	//Vollrechteck(50,100,650,500)
	Vollrechteck(400,100,650,500)
	Transparenz(0)

}


