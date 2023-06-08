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
)


func StartFenster() spielstaende.Spielstand {
		
		
	var spielstand spielstaende.Spielstand = Startbildschirm()	
	
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
				return spielstand
			}
		}
	}

	return spielstand
	
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
