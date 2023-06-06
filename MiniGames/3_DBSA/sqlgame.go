//Annalena Cyriacus
//Juni 2023
//LWB-Adventure: Minigame "SQL-Quest"

package sqlGame

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
var ausg2 []string = make ([]string,7)
var ausg3 []string = make ([]string,8)
var ausg4 []string = make ([]string,5)
var ausg5 []string = make ([]string,3)
var ausg6 []string = make ([]string,3)
var ausg7 []string = make ([]string,3)
var ausg8 []string = make ([]string,3)
var ausg9 []string = make ([]string,5)
var ausg10 []string = make ([]string,7)
var eingTexte []string = make([]string,14)
var eingTexte2 []string = make([]string,14)

//----------------------Hilfsfunktionen----------------------------

func bubbleTexte() {
	texte[0] = "Um die DBSA-Prüfung zu bestehen,\nmusst Du Dein SQL-Wissen nutzen,\num alle Geheimnisse der LWB-Welt\nzu enthüllen!"
	texte[1] = "Die LWB-Adventure-World ist zwar nicht ganz so verwirrend wie die STEPS-World, aber was für Räume gibt es denn hier eigentlich?"
	texte[2] = "Welche Lehrveranstaltungen finden im 4. Semester statt?"
	texte[3] = "Wie heißen die Dozenten im LWB-Adventure?"
	texte[4] = "Welche Lehrveranstaltungen haben etwas mit 'Programmierung' zu tun?"
	texte[5] = "Ob Du es auch schaffst herauszubekommen, was mein ganz persönliches Lieblingsgetränk ist?"
	texte[6] = "Wieviele Mini-Games gibt es in der LWB-Adventure-World? (Überschrift der Ausgabe: AnzahlMiniGames)"
	texte[7] = "Wie heißt die Veranstaltung mit den meisten SWS?"
	texte[8] = "Wieviele SWS müssen in der gesamten LWB absolviert werden? (Überschrift der Ausgabe: GesamtanzahlSWS)"
	texte[9] = "Lasse die Anzahl der Veranstaltungen pro Standort anzeigen! (Ausgabe aufsteigend, Überschrift der Ausgabe: AnzahlVeranstaltungen)"
	texte[10] = "Mal schauen, ob Du die sechs Kommandos richtig eingibst, mit denen Dir die Namen, Semester und SWS aller Veranstaltungen von Winnie the K nach SWS-Anzahl absteigend sortiert angezeigt werden!"
}

func eingabeTexte() {
	eingTexte[1] = "SELECT*FROMraeume;"
	eingTexte[2] = "SELECT*FROMveranstaltungenWHEREsemester=4;"
	eingTexte[3] = "SELECTdoznameFROMdozenten;"
	eingTexte[4] = "SELECT*FROMveranstaltungenWHEREvnameLIKE'%Programmierung';"
	eingTexte[5] = "SELECTlieblingsgetraenkFROMdozentenWHEREdozname='Herk';"
	eingTexte[6] = "SELECTCOUNT(*)ASAnzahlMiniGamesFROMmini-games;"
	eingTexte[7] = "SELECTvname,MAX(sws)FROMveranstaltungen;"
					//"SELECTvname,MAX(sws)FROMveranstaltungenGROUPBYvname;"			//TODO!!!
	eingTexte[8] = "SELECTSUM(sws)ASGesamtanzahlSWSFROMveranstaltungen;"
	eingTexte[9] = "SELECTort,COUNT(*)ASAnzahlVeranstaltungenFROMraeume,veranstaltungenWHEREraumnr=semesterGROUPBYortORDERBYCOUNT(*);"					//TODO
	eingTexte[10] = "SELECTvname,semester,swsFROMveranstaltungenNATURALJOINdozentenWHEREdozname='WinnietheK'ORDERBYswsDESC;"
}

func eingabeTexte2() {
	//eingTexte2[1] = "SELECT*FROMraeume;"
	//eingTexte2[2] = "SELECT * FROM veranstaltungen WHERE semester = 4;"
	//eingTexte2[3] = "SELECT dozname FROM dozenten;"
	//eingTexte2[4] = "SELECT * FROM veranstaltungen WHERE vname LIKE '%Programmierung';"
	//eingTexte2[5] = "SELECT lieblingsgetraenk FROM dozenten WHERE dozname = 'Herk';"
	//eingTexte2[6] = "SELECT COUNT(*) FROM mini-games AS Anzahl_MiniGames;"
	//eingTexte2[7] = "SELECT vname, MAX(sws) FROM veranstaltungen;"
	//eingTexte2[8] = "SELECT SUM(sws) FROM veranstaltungen AS Gesamtanzahl_SWS;"
	//eingTexte2[9] = "SELECT ort, COUNT(*) FROM raeume, veranstaltungen AS Anzahl_Veranstaltungen WHERE raumnr = semester GROUP BY ort ORDER BY COUNT(*);"					//TODO
	//eingTexte2[10] = "SELECT vname, semester, sws FROM veranstaltungen NATURAL JOIN dozenten WHERE dozname = 'Winnie the K' ORDER BY sws DESC;"
}


func ausgabeTexte() {
	
	ausg1[0] = " RAUM-NR.     RAUM-NAME                      ORT"
	ausg1[1] = "------------------------------------------------------------------"
	ausg1[2] = " 1            1. Semester                    FU (Dahlem)"
	ausg1[3] = " 2            2. Semester                    digital (BBB)"
	ausg1[4] = " 3            3. Semester                    FU (Dahlem)"
	ausg1[5] = " 4            4. Semester                    STEPS (Mitte)"
	ausg1[6] = " 5            Nichtzeugnis-Verleihung        N.N."
	ausgTexte[1] = ausg1
	
	ausg2[0] = " V-NR.   VERANSTALTUNGS-NAME                             DOZ-NR.   SEM.    SWS"
	ausg2[1] = "--------------------------------------------------------------------------------"
	ausg2[2] = " 4.1     Nichtsequentielle & verteilte Programmierung    2         4       6"
	ausg2[3] = " 4.2     Rechnernetze                                    6         4       2"
	ausg2[4] = " 4.3     Unterrichtsbezogenes Softwarepraktikum          1         4       3"
	ausg2[5] = " 4.4     Unterrichtsbezogenes Datenbankpraktikum         4         4       3"
	ausg2[6] = " 4.5     Analyse fachlichen Lernens                      3         4       2"
	ausgTexte[2] = ausg2
	
	ausg3[0] = " DOZENTEN-NAME"
	ausg3[1] = "---------------"
	ausg3[2] = " Amoebi"
	ausg3[3] = " Darth Schmidter"
	ausg3[4] = " Fab Web"
	ausg3[5] = " Herk"
	ausg3[6] = " J.EthI"
	ausg3[7] = " Winnie the K"
	ausgTexte[3] = ausg3
		
	ausg4[0] = " V-NR.  VERANSTALTUNGS-NAME                              DOZ-NR.  SEM.    SWS"
	ausg4[1] = "--------------------------------------------------------------------------------"
	ausg4[2] = " 1.2    Funktionale Programmierung                        3        1       8"
	ausg4[3] = " 2.1    Imperative & projektorientierte Programmierung    2        2       7"
	ausg4[4] = " 4.1    Nichtsequentielle & verteilte Programmierung      2        4       9"
	ausgTexte[4] = ausg4
	
	ausg5[0] = " LIEBLINGSGETRÄNK"
	ausg5[1] = "------------------"
	ausg5[2] = " Beruhigungstee"											// WIRKLICH??? :D
	ausgTexte[5] = ausg5
	
	ausg6[0] = " Anzahl_Minigames"
	ausg6[1] = "------------------"
	ausg6[2] = " 9"
	ausgTexte[6] = ausg6
	
	ausg7[0] = " V-NAME                         SWS"
	ausg7[1] = "-----------------------------------------"
	ausg7[2] = " Funktionale Programmierung     8"
	ausgTexte[7] = ausg7
	
	ausg8[0] = " Gesamtanzahl_SWS"
	ausg8[1] = "-----------------"
	ausg8[2] = " 64"
	ausgTexte[8] = ausg8
	
	ausg9[0] = " Ort               Anzahl-Veranstaltungen"
	ausg9[1] = "------------------------------------------"
	ausg9[2] = " FU (Dahlem)       6"
	ausg9[3] = " digital (BBB)     3"
	ausg9[4] = " STEPS (Mitte)     5"
	ausgTexte[9] = ausg9
	
	ausg10[0] = " V-NAME                                    SEMESTER        SWS"
	ausg10[1] = "---------------------------------------------------------------"
	ausg10[2] = " Betriebssystemwerkzeuge                   1. Semester     2"
	ausg10[3] = " Grundlagen der technischen Informatik     1. Semester     6"
	ausg10[4] = " Rechnerarchitektur                        2. Semester     4"
	ausg10[5] = " Fachdidaktik Informatik                   3. Semester     4"
	ausg10[6] = " Rechnernetze                              4. Semester     2"
	ausgTexte[10] = ausg10
}

func erzeugeFehlerausgabe(ausgabe textboxen.Textbox) {
	ausgabe.RahmenAn(true)
	ausgabe.SetzeRahmenFarbe(255,0,0)
	ausgabe.SetzePosition(320,527)
	ausgabe.SetzeHöhe(43)
	ausgabe.SetzeZeilenAbstand(3)
	ausgabe.SetzeFarbe(255,0,0)
	ausgabe.SchreibeText("FALSCHE EINGABE! --> Überprüfe die Anfrage / korrigiere die Schreibweise!!!\n(Kommandos GROSS, sonst klein, strings mit '...' und Simikolon nicht vergessen!)")
	ausgabe.Zeichne()
}

func erzeugeAusgabe(n int) {
	Stiftfarbe(0,255,0)
	SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",20)
	for i:=0; i<len(ausgTexte[n]); i++ {
		SchreibeFont(328,uint16(383+i*22),ausgTexte[n][i])
	}
}

func ohneLeerzeichen(s string) string {
	var erg string
	for _,zeichen := range s {
		if zeichen != 32 {
			erg = erg + string(zeichen)
		}
	}
	return erg
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
	var bubbletext textboxen.Textbox = textboxen.New(220,150,305,120)
	var infotext textboxen.Textbox = textboxen.New(600,90,540,120)
	var ted texteditoren.Texteditor
	
//------------------Grafik-Elemente--------------------------------
	
	LadeBild(150,90,path2 + "Bilder/bubble2_flipped_400.bmp")
	LadeBildMitColorKey(50,400,path2 + "Bilder/SQLGame/Herk_200.bmp",255,255,255)
		
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
	datainfo.SchreibeText("> raeume (raumnr, raumname, ort)\n> dozenten (doznr, dozname, lieblingsgetraenk)\n> veranstaltungen (vnr, vname, doznr, semester, sws)\n> mini-games (gamenr, gamename, vnr, maxpunktzahl)")
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
	bubblehead.SetzeSchriftgröße(22)
	bubblehead.SchreibeText("Willkommen zum SQL-Quest!")
	bubblehead.Zeichne()
	
	next.SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf")
	next.ZeichneButton()
	
//----------------Eingabe-Verarbeitung----------------------------

	//ted = texteditoren.New(315,440,830,218,20,true)
	
//----------------Spiel-Steuerung-----------------------------------

	for i:=1; i<len(eingTexte); i++ {
		fmt.Println("aktuelles i:",i)
	
		for {

			taste, status, mausX, mausY := MausLesen1()
			if taste==1 && status==1 {
				if next.TesteXYPosInButton(mausX,mausY) {
					fmt.Println("Weiter angeklickt!")
					Stiftfarbe(255,255,255)
					Vollrechteck(220,115,305,170)
					bubbletext.SchreibeText(texte[aktuellerText])
					bubbletext.Zeichne()
					bubblehead.SchreibeText("Level "+fmt.Sprint(i))
					bubblehead.Zeichne()
					next.ZeichneButton()
					//next.DeaktiviereButton()
					if aktuellerText < len(texte)-1 {
						aktuellerText++
					}
					break
				}
			}
		}
						
		//----------------Eingabe-Verarbeitung--------------
		ted = texteditoren.New(315,595,830,63,20,true)
					
		for {
			//var eing string = ohneLeerzeichen(ted.GibString())
			//if  eing == eingTexte[i] || eing == eingTexte2[i] {
			if ohneLeerzeichen(ted.GibString()) == eingTexte[i] {
				//Stiftfarbe(255,255,255)
				//Vollrechteck(443,238,84,39)
				SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf",30)
				Stiftfarbe(255,255,255)
				Vollrechteck(220,235,315,40)
				Vollrechteck(100,500,100,100)
				LadeBildMitColorKey(50,400,path2 + "Bilder/SQLGame/herk_200.bmp",255,255,255)
				Stiftfarbe(0,255,0)
				SchreibeFont(230,240,"RICHTIG!!!  :-)")
				next.ZeichneButton()
				ausgabe.RahmenAn(true)
				ausgabe.SetzeRahmenFarbe(0,0,0)
				ausgabe.HintergrundAn(true)
				ausgabe.SetzeHintergrundFarbe(255,255,255)
				ausgabe.SchreibeText("")				
				ausgabe.Zeichne()
				Stiftfarbe(255,255,255)
				Vollrechteck(311,375,838,200)
				erzeugeAusgabe(i)
				//next.AktiviereButton()
				break
			} else {
				Stiftfarbe(255,255,255)
				Transparenz(50)
				Vollrechteck(443,238,84,39)
				Transparenz(0)
				Vollrechteck(320,375,820,200)
				Vollrechteck(100,500,100,100)
				LadeBildMitColorKey(50,400,path2 + "Bilder/SQLGame/herk_angry_red_200.bmp",255,255,255)
				erzeugeFehlerausgabe(ausgabe)
				SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf",30)
				Stiftfarbe(255,0,0)
				SchreibeFont(225,240,"NOCHMAL!  :-(")
				ted = texteditoren.New(315,595,830,63,20,true)
			}					
		}
	}
		

	TastaturLesen1()
	
}
