//Annalena Cyriacus
//Juni 2023
//LWB-Adventure: Minigame "SQL-Quest"

package sqlgame

import (
	. "gfx"
	"fmt"
	"../../Klassen/buttons"
	"../../Klassen/textboxen"
	"../../Klassen/texteditoren"
)

var path string = ""
var path2 string = "./" //MiniGames/2_ALP2/"

var texte []string = make([]string,14)
var aktuellerText int = 1
var ausgTexte [][]string = make([][]string,14)
var ausg1 []string = make ([]string,7)
//var ausg2,ausg3,ausg4,ausg5,ausg6,ausg7 string
var eingTexte []string = make([]string,14)

//----------------------Hilfsfunktionen----------------------------

func bubbleTexte() {
	texte[0] = "\n\nUm die DBSA-Prüfung zu bestehen,\nmusst Du Dein SQL-Wissen nutzen,\num alle Geheimnisse der LWB-Welt\nzu enthüllen!"
	texte[1] = "Die LWB-Adventure-World ist zwar nicht ganz so verwirrend wie die STEPS-World, aber was für Räume gibt es denn hier eigentlich?"
	texte[2] = "Aufgabe 2"
	texte[3] = "Aufgabe 3"
	texte[4] = "Aufgabe 4"
}

func eingabeTexte() {
	eingTexte[1] = "SELECT  FROM raeume"
}

func ausgabeTexte() {
	
	ausg1[0] = "RAUM-NR.     RAUM-NAME                      ORT"
	ausg1[1] = "----------------------------------------------------------------"
	ausg1[2] = "1            1. Semester                    FU/KL24-26"
	ausg1[3] = "2            2. Semester                    online/homeoffice"
	ausg1[4] = "3            3. Semester                    FU/Arnimallee"
	ausg1[5] = "4            4. Semester                    STEPS/Georgenstraße"
	ausg1[6] = "5            Nichtzeugnis-Verleihung        N.N."
	ausgTexte[1] = ausg1
		
}

func erzeugeFehlerausgabe(ausgabe textboxen.Textbox) {
	ausgabe.RahmenAn(true)
	ausgabe.SetzeRahmenFarbe(255,0,0)
	ausgabe.SetzePosition(320,550)
	ausgabe.SetzeHöhe(20)
	ausgabe.SchreibeText("")
	ausgabe.Zeichne()
	Stiftfarbe(255,0,0)
	SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",20)
	SchreibeFont(328,550,"FALSCHE EINGABE! --> Überprüfe die Anfrage / korrigiere die Schreibweise!")
}

func erzeugeAusgabe(n int) {
	Stiftfarbe(0,255,0)
	SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",20)
	for i:=0; i<len(ausgTexte[n]); i++ {
		SchreibeFont(328,uint16(383+i*22),ausgTexte[n][i])
	}
}

//func SQLgame() (float32,uint32) {
func SQLgame() {
	
//-----------------initialisiere gfx-Fenster-----------------------	
	Fenster(1200,700)
	SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	bubbleTexte()
	eingabeTexte()
	ausgabeTexte()
	
//------------------Variablen--------------------------------------
	var next buttons.Button = buttons.New(445,240,80,35,0,255,0,true,"   next")
	var firsted textboxen.Textbox = textboxen.New(320,600,820,48)
	var ausgabe textboxen.Textbox = textboxen.New(320,380,820,190)
	var datainfo textboxen.Textbox = textboxen.New(610,265,530,85)
	//var fehlerausgabe textboxen.Textbox = textboxen.New(320,550,820,25)
	SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf",50)
	var bubblehead textboxen.Textbox = textboxen.New(220,115,305,155)
	SetzeFont(path2 + "Schriftarten/Ubuntu-R.ttf",50)
	var bubbletext textboxen.Textbox = textboxen.New(220,115,305,155)
	var infotext textboxen.Textbox = textboxen.New(600,90,540,120)
	var ted texteditoren.Texteditor
	
//------------------Grafik-Elemente--------------------------------
	
	LadeBild(150,90,path2 + "Bilder/bubble2_flipped_400.bmp")
	LadeBildMitColorKey(50,400,path2 + "Bilder/Herk_200.bmp",255,255,255)
		
	SetzeFont(path2 + "Schriftarten/brlnsdb.ttf",60)
	Stiftfarbe(0,255,0)
	SchreibeFont(70,10,"SQL-Quest:")
	SetzeFont(path2 + "Schriftarten/brlnsr.ttf",50)
	SchreibeFont(390,20,"Explore the LWB-Adventure-World!")
	
	//infotext.SetzeFarbe(0,0,0)
	infotext.SetzeZeilenAbstand(5)
	infotext.SetzeSchriftgröße(18)
	infotext.SchreibeText("Herk hat eine Datenbank erstellt, die alle Geheimnisse der LWB-Welt enthält. Würdig darauf zuzugreifen sind nur jene, die SQL beherrschen.\nAlso schärfe Deinen Verstand und gib die richtigen Anfragen ein, um die Geheimnisse zu ergründen!\n\nFolgende Tabellen sind in seiner Datenbank enthalten:")
	infotext.Zeichne()
	
	datainfo.RahmenAn(true)
	datainfo.HintergrundAn(true)
	datainfo.SetzeRahmenFarbe(0,0,0)
	datainfo.SetzeHintergrundFarbe(0,0,0)
	datainfo.SetzeFarbe(255,255,255)
	datainfo.SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf")
	datainfo.SetzeSchriftgröße(18)
	datainfo.SetzeZeilenAbstand(3)
	datainfo.SchreibeText("> raeume (raumnr, raumname, ort)\n> dozenten (doznr, dozname, lieblingsgetränk)\n> veranstaltungen (vnr, vname, thema, semester, sws)\n> mini-games (gamenr, gamename, maxpunktzahl)")
	datainfo.Zeichne()
	
	ausgabe.RahmenAn(true)
	ausgabe.SetzeRahmenFarbe(0,0,0)
	ausgabe.HintergrundAn(true)
	ausgabe.SetzeHintergrundFarbe(255,255,255)
	ausgabe.SetzeFarbe(0,255,0)
	ausgabe.SetzeZeilenAbstand(5)
	ausgabe.SchreibeText("Hier erscheint das Ergebnis Deiner SQL-Anfrage, sobald Du sie eingegeben und mit ENTER ausgeführt hast! (Hinweis: Kommandos GROSS schreiben, alles andere klein!)")
	ausgabe.Zeichne()
	ausgabe.RahmenAn(false)
	
	/*
	fehlerausgabe.RahmenAn(true)
	fehlerausgabe.SetzeRahmenFarbe(255,0,0)
	fehlerausgabe.HintergrundAn(true)
	fehlerausgabe.SetzeHintergrundFarbe(255,255,255)
	fehlerausgabe.SetzeFarbe(255,0,0)
	fehlerausgabe.SchreibeText("Falsche Eingabe! Überprüfe die Anfrage und kontrolliere die Schreibweise!")
	*/
	
	firsted.HintergrundAn(true)
	firsted.SetzeHintergrundFarbe(0,0,0)
	firsted.SetzeFarbe(255,255,255)
	firsted.SchreibeText("SELECT ...  <-- Gib Deine SQL-Anfrage hier ein und führe Sie mit ENTER aus!")
	firsted.Zeichne()
	
	//bubbletext.RahmenAn(true)
	//bubbletext.SetzeRahmenFarbe(0,0,0)
	bubbletext.HintergrundAn(true)
	bubbletext.SetzeHintergrundFarbe(255,255,255)
	bubbletext.SetzeZeilenAbstand(3)
	bubbletext.SetzeSchriftgröße(18)
	bubbletext.SchreibeText(texte[0])
	bubbletext.Zeichne()
	bubblehead.SchreibeText("Willkommen zum SQL-Quest!")
	bubblehead.Zeichne()
	
	next.SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf")
	next.ZeichneButton()
	
//----------------Eingabe-Verarbeitung----------------------------

	//ted = texteditoren.New(315,440,830,218,20,true)
	
//----------------Spiel-Steuerung-----------------------------------

	for {

		taste, status, mausX, mausY := MausLesen1()
		if taste==1 && status==1 {
			if next.TesteXYPosInButton(mausX,mausY) {
				fmt.Println("Weiter angeklickt!")
				bubbletext.SchreibeText(texte[aktuellerText])
				bubbletext.Zeichne()
				next.ZeichneButton()
				if aktuellerText < len(texte)-1 {
					aktuellerText++
				}
				
				for i:=1; i<len(eingTexte); i++ {
					
					fmt.Println("aktuelles i:",i)
					//----------------Eingabe-Verarbeitung--------------
					ted = texteditoren.New(315,595,830,63,20,true)
								
						for {
							switch ted.GibString() {
								case eingTexte[i]:
								ausgabe.RahmenAn(true)
								ausgabe.SetzeRahmenFarbe(0,0,0)
								ausgabe.HintergrundAn(true)
								ausgabe.SetzeHintergrundFarbe(255,255,255)
								ausgabe.SchreibeText("")
								ausgabe.Zeichne()
								Stiftfarbe(255,255,255)
								Vollrechteck(311,375,838,200)
								erzeugeAusgabe(i)	
								break					
								default:
								Stiftfarbe(255,255,255)
								Vollrechteck(320,375,820,200)
								erzeugeFehlerausgabe(ausgabe)
								break
							}
							ted = texteditoren.New(315,595,830,63,20,true)					
						}
				}
			}
		}
	}
		

	TastaturLesen1()
	
}
