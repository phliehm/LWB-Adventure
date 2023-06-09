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
	"../../Klassen/vierecke"
)

var path string = ""
var path2 string = "./" //MiniGames/2_ALP2/"

var punkte, fehler uint32
//var note float32

var texte []string = make([]string,11)
var aktuellerText int = 1

var ausgTexte [][]string = make([][]string,11)
var ausg1 []string = make ([]string,7)
var ausg2 []string = make ([]string,7)
var ausg3 []string = make ([]string,8)
var ausg4 []string = make ([]string,5)
var ausg5 []string = make ([]string,3)
var ausg6 []string = make ([]string,3)
var ausg7 []string = make ([]string,3)
var ausg8 []string = make ([]string,3)
var ausg9 []string = make ([]string,7)
var ausg10 []string = make ([]string,5)

var eingTexte []string = make([]string,11)
var eingTexte2 []string = make([]string,11)

var hilfeTexte [][]string = make([][]string,11)
var hilfen1 []string = make ([]string,4)
var hilfen2 []string = make ([]string,4)
var hilfen3 []string = make ([]string,4)
var hilfen4 []string = make ([]string,4)
var hilfen5 []string = make ([]string,4)
var hilfen6 []string = make ([]string,4)
var hilfen7 []string = make ([]string,4)
var hilfen8 []string = make ([]string,4)
var hilfen9 []string = make ([]string,4)
var hilfen10 []string = make ([]string,4)

//----------------------Hilfsfunktionen----------------------------

func bubbleTexte() {
	texte[0] = "Um die DBSA-Prüfung zu bestehen,\nmusst Du Dein SQL-Wissen nutzen,\num alle Geheimnisse der LWB-Welt\nzu enthüllen!"
	texte[1] = "Die LWB-Adventure-World ist zwar nicht ganz so verwirrend wie die STEPS-World, aber was für Räume gibt es denn hier eigentlich?"
	texte[2] = "Welche Lehrveranstaltungen finden im 4. Semester statt?"
	texte[3] = "Wie heißen die Dozenten im LWB-Adventure?"
	texte[4] = "Welche Lehrveranstaltungen haben etwas mit 'Programmierung' zu tun?"
	texte[5] = "Ob Du es auch schaffst herauszubekommen, was mein ganz persönliches Lieblingsgetränk ist?"
	texte[6] = "Wieviele Mini-Games gibt es in der LWB-Adventure-World? (Überschrift der Ausgabe: AnzahlMiniGames)"
	texte[8] = "Wie heißt die Veranstaltung mit den meisten SWS?"
	texte[7] = "Wieviele SWS müssen in der LWB insgesamt absolviert werden? (Überschrift der Ausgabe: GesamtanzahlSWS)"
	//texte[10] = "Lasse zu jedem Raum die Anzahl der dort stattfindenden Veranstaltungen anzeigen! (Ausgabe aufsteigend, Überschrift der Anzahl-Ausgabe: AnzahlVeranstaltungen)"
	texte[9] = "Mal schauen, ob Du die sechs Kommandos richtig eingibst, mit denen Dir die Namen, Semester und SWS aller Veranstaltungen von Winnie the K nach SWS-Anzahl absteigend sortiert angezeigt werden!"
	texte[10] = "Lasse die Anzahl der Veranstaltungen pro Standort anzeigen! (ohne JOIN!, Ausgabe aufsteigend, Überschrift der Ausgabe: AnzahlVeranstaltungen)"
}

func eingabeTexte() {
	eingTexte[1] = "SELECT*FROMraeume;"
	eingTexte[2] = "SELECT*FROMveranstaltungenWHEREsemester=4;"
	eingTexte[3] = "SELECTdoznameFROMdozenten;"
	eingTexte[4] = "SELECT*FROMveranstaltungenWHEREvnameLIKE'%Programmierung';"
	eingTexte[5] = "SELECTlieblingsgetraenkFROMdozentenWHEREdozname='Herk';"
	eingTexte[6] = "SELECTCOUNT(*)ASAnzahlMiniGamesFROMmini-games;"
	eingTexte[7] = "SELECTSUM(sws)ASGesamtanzahlSWSFROMveranstaltungen;"
	eingTexte[8] = "SELECTvnameFROMveranstaltungenWHEREsws=(SELECTMAX(sws)FROMveranstaltungen);"
	//eingTexte[10] = "SELECTsemester,COUNT(*)ASAnzahlVeranstaltungenFROMveranstaltungenGROUPBYsemesterORDERBYsemester";
	eingTexte[9] = "SELECTvname,semester,swsFROMveranstaltungenNATURALJOINdozentenWHEREdozname='WinnietheK'ORDERBYswsDESC;"
	eingTexte[10] = "SELECTort,COUNT(*)ASAnzahlVeranstaltungenFROMraeume,veranstaltungenWHEREraumnr=semesterGROUPBYortORDERBYCOUNT(*);"
	
}

func eingabeTexte2() {
	//eingTexte2[1] = ""
	//eingTexte2[2] = ""
	//eingTexte2[3] = ""
	eingTexte2[4] = "SELECT*FROMveranstaltungenWHEREvnameLIKE'%Programmierung%';"
	//eingTexte2[5] = ""
	//eingTexte2[6] = ""
	//eingTexte2[7] = ""
	//eingTexte2[8] = ""
	//eingTexte2[9] = ""
	//eingTexte2[10] = ""
}

func hilfeText() {
	
	hilfen1[0] = ""
	hilfen1[1] = "Hilfe 1: Nutze SELECT ... FROM ...!"
	hilfen1[2] = "Hilfe 2: Wenn die Ausgabe alle Attribute eines Eintrags enthalten soll, nutze * !"
	hilfen1[3] = "Leider war auch der 3. Versuch falsch!\nDas wäre die richtige Lösung gewesen:\n\nSELECT * FROM raeume;"
	hilfeTexte[1] = hilfen1
	
	hilfen2[0] = ""
	hilfen2[1] = "Hilfe 1: Nutze ... WHERE ... = ...!"
	hilfen2[2] = "Hilfe 2: Wenn die Ausgabe alle Attribute eines Eintrags enthalten soll, nutze * !"
	hilfen2[3] = "Leider war auch der 3. Versuch falsch!\nDas wäre die richtige Lösung gewesen:\n\nSELECT * FROM veranstaltungen WHERE semester = 4;"
	hilfeTexte[2] = hilfen2
	
	hilfen3[0] = ""
	hilfen3[1] = "Hilfe 1: Nutze nur SELECT ... FROM ...!"
	hilfen3[2] = "Hilfe 2: Wenn die Ausgabe nur ein bestimmtes Attribut enthalten soll, nutze statt * den Attributnamen!"
	hilfen3[3] = "Leider war auch der 3. Versuch falsch!\nDas wäre die richtige Lösung gewesen:\n\nSELECT dozname FROM dozenten;"
	hilfeTexte[3] = hilfen3
	
	hilfen4[0] = ""
	hilfen4[1] = "Hilfe 1: Nutze ... WHERE ... LIKE ...!"
	hilfen4[2] = "Hilfe 2: Wie wäre es mit ... '%Programmierung' ?"
	hilfen4[3] = "Leider war auch der 3. Versuch falsch!\nDas wäre die richtige Lösung gewesen:\n\nSELECT * FROM veranstaltungen WHERE vname LIKE '%Programmierung';"
	hilfeTexte[4] = hilfen4
	
	hilfen5[0] = ""
	hilfen5[1] = "Hilfe 1: Nutze ... WHERE ... = ...!"
	hilfen5[2] = "Hilfe 2: Denke daran, dass das Attribut dozname ein String ist, also '...' genutzt werden muss!"
	hilfen5[3] = "Leider war auch der 3. Versuch falsch!\nDas wäre die richtige Lösung gewesen:\n\nSELECT lieblingsgetraenk FROM dozenten WHERE dozname = 'Herk';"
	hilfeTexte[5] = hilfen5
	
	hilfen6[0] = ""
	hilfen6[1] = "Hilfe 1: Nutze SELECT COUNT(*) ... !"
	hilfen6[2] = "Hilfe 2: Um die Ausgabe-Überschrift festzulegen, brauchst Du zwischen COUNT(*) und FROM das Kommando AS ... !"
	hilfen6[3] = "Leider war auch der 3. Versuch falsch!\nDas wäre die richtige Lösung gewesen:\n\nSELECT COUNT(*) AS AnzahlMiniGames FROM mini-games;"
	hilfeTexte[6] = hilfen6
	
	hilfen7[0] = ""
	hilfen7[1] = "Hilfe 1: Nutze SELECT SUM(...) ... !"
	hilfen7[2] = "Hilfe 2: Auch hier brauchst Du  das Kommando AS ... !"
	hilfen7[3] = "Leider war auch der 3. Versuch falsch!\nDas wäre die richtige Lösung gewesen:\n\nSELECT SUM(sws) AS GesamtanzahlSWS FROM veranstaltungen;"
	hilfeTexte[7] = hilfen7
	
	hilfen8[0] = ""
	hilfen8[1] = "Hilfe 1: Nutze ... WHERE ... = (SELECT ...)!"
	hilfen8[2] = "Hilfe 2: In der zweiten SELECT-Anweisung brauchst Du das Kommando MAX(...)"
	hilfen8[3] = "Leider war auch der 3. Versuch falsch!\nDas wäre die richtige Lösung gewesen:\n\nSELECT vname FROM veranstaltungen WHERE sws = (SELECT MAX(sws) FROM veranstaltungen);"
	hilfeTexte[8] = hilfen8
	
	hilfen9[0] = ""
	hilfen9[1] = "Hilfe 1: Drei der sechs benötigten Kommandos sind NATURAL JOIN, WHERE und ORDER BY!"
	hilfen9[2] = "Hilfe 2: Gib genau an, welche Atrribute in der Ausgabe enthalten sein sollen und denke am Ende an die absteigende Sortierung!"
	hilfen9[3] = "Leider war auch der 3. Versuch falsch!\nDas wäre die richtige Lösung gewesen:\n\nSELECT vname, semester, sws FROM veranstaltungen NATURAL JOIN dozenten WHERE dozname = 'WinnietheK' ORDER BY sws DESC;"
	hilfeTexte[9] = hilfen9
	
	hilfen10[0] = ""
	hilfen10[1] = "Hilfe 1: Nutze u.a. COUNT(*), GROUP BY und ORDER BY!"
	hilfen10[2] = "Hilfe 2: Nutze das Kartesische Produkt und die Tatsache, dass raumnr und semester gleich sind! Achte außerdem auf die Attribute, die in der Ausgabe enthalten sein sollen!"
	hilfen10[3] = "Leider war auch der 3. Versuch falsch!\nDas wäre die richtige Lösung gewesen:\n\nSELECT ort, COUNT(*) AS AnzahlVeranstaltungen FROM raeume, veranstaltungen WHERE raumnr = semester GROUP BY ort ORDER BY COUNT(*);"
	hilfeTexte[10] = hilfen10
	
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
	
	ausg10[0] = " Ort               Anzahl-Veranstaltungen"
	ausg10[1] = "------------------------------------------"
	ausg10[2] = " FU (Dahlem)       6"
	ausg10[3] = " digital (BBB)     3"
	ausg10[4] = " STEPS (Mitte)     5"
	ausgTexte[10] = ausg10
	
	ausg9[0] = " V-NAME                                    SEMESTER        SWS"
	ausg9[1] = "---------------------------------------------------------------"
	ausg9[2] = " Betriebssystemwerkzeuge                   1. Semester     2"
	ausg9[3] = " Grundlagen der technischen Informatik     1. Semester     6"
	ausg9[4] = " Rechnerarchitektur                        2. Semester     4"
	ausg9[5] = " Fachdidaktik Informatik                   3. Semester     4"
	ausg9[6] = " Rechnernetze                              4. Semester     2"
	ausgTexte[9] = ausg9
}

func initialausgabe(ausgabe textboxen.Textbox) {
	//Stiftfarbe(255,255,255)
	//Vollrechteck(310,370,840,210)
	ausgabe.SetzeHöhe(190)
	ausgabe.RahmenAn(true)
	ausgabe.SetzeRahmenFarbe(0,0,0)
	ausgabe.HintergrundAn(true)
	ausgabe.SetzeHintergrundFarbe(255,255,255)
	ausgabe.SetzeFarbe(0,255,0)
	ausgabe.SetzeZeilenAbstand(5)
	ausgabe.SchreibeText("Hier erscheint das Ergebnis Deiner SQL-Anfrage, sobald Du sie eingegeben und mit ENTER ausgeführt hast! (Hinweis: Kommandos GROSS schreiben, alles andere klein!)")
	ausgabe.Zeichne()
	//ausgabe.RahmenAn(false)
}

func erzeugeFehlerausgabe() {
	var falschausgabe textboxen.Textbox = textboxen.New(320,380,820,43)
	falschausgabe.RahmenAn(true)
	falschausgabe.SetzeRahmenFarbe(255,0,0)
	falschausgabe.SetzePosition(320,527)
	//falschausgabe.SetzeHöhe(43)
	falschausgabe.SetzeZeilenAbstand(3)
	falschausgabe.SetzeFarbe(255,0,0)
	falschausgabe.SchreibeText("FALSCHE EINGABE! --> Überprüfe die Anfrage / korrigiere die Schreibweise!!!\n(Kommandos GROSS, sonst klein, strings mit '...' und Simikolon nicht vergessen!)")
	falschausgabe.Zeichne()
}

func erzeugeAusgabe(n int) {
	Stiftfarbe(0,255,0)
	SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",20)
	for i:=0; i<len(ausgTexte[n]); i++ {
		SchreibeFont(328,uint16(383+i*22),ausgTexte[n][i])
	}
}

//Punkte und Note ausgeben
func punktenoteSchreiben(punktenote textboxen.Textbox, levelpunkte, punkte uint32) {
	punktenote.HintergrundAn(true)
	punktenote.SetzeHintergrundFarbe(255,255,255)
	if punkte < 10 {
		punktenote.SchreibeText("Punkte: "+fmt.Sprint(levelpunkte)+"\n\nNote: ")
	} else {
		punktenote.SchreibeText("Punkte: "+fmt.Sprint(punkte)+"\n\nNote: ")
	}
	
	if notenberechnung(punkte) > 0 {
		punktenote.SetzeHöhe(250)
		punktenote.SchreibeText("Punkte: "+fmt.Sprint(punkte)+"\n\nNote: "+fmt.Sprintf("%2.1f",notenberechnung(punkte)))
	}
	punktenote.Zeichne()
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

func notenberechnung(punkte uint32) float32 {
	if punkte >= 100 { return 1.0
	} else if punkte >= 90 { 
		return 1.3
	} else if punkte >= 80 { 
		return 1.7
	} else if punkte >= 75 {
		return 2.0
	} else if punkte >= 70 {
		return 2.3
	} else if punkte >= 65 {
		return 2.7
	} else if punkte >= 60 {
		return 3.0
	} else if punkte >= 55 {
		return 3.3
	} else if punkte >= 50 {
		return 4.0
	} else { return 0.0 }
}

//func SQLgame() (float32,uint32) {
func SQLgame() (punkte uint32, note float32) {
	
	//punkte = 10
	//fmt.Println("Punkte:",fmt.Sprint(punkte))
	
//-----------------initialisiere gfx-Fenster-----------------------	
	if ! FensterOffen() {
		Fenster(1200,700)
	}
	Stiftfarbe(255,255,255)
	Vollrechteck(0,0,1200,700)
	SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	
	fmt.Println("Führe bubbleTexte() aus!")
	bubbleTexte()
	fmt.Println("Führe eingabeTexte() aus!")
	eingabeTexte()
	fmt.Println("Führe ausgabeTexte() aus!")
	ausgabeTexte()
	fmt.Println("Führe hilfeText() aus!")
	hilfeText()
	
//------------------Variablen--------------------------------------
	
	var next buttons.Button = buttons.New(445,240,80,35,0,255,0,true,"   next")
	var firsted textboxen.Textbox = textboxen.New(320,600,820,48)
	var ausgabe textboxen.Textbox = textboxen.New(320,380,820,190)
	//var falschausgabe textboxen.Textbox = textboxen.New(320,380,820,43)
	var datainfo textboxen.Textbox = textboxen.New(610,265,530,85)
	//var fehlerausgabe textboxen.Textbox = textboxen.New(320,550,820,25)
	SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf",50)
	var bubblehead textboxen.Textbox = textboxen.New(220,115,305,155)
	SetzeFont(path2 + "Schriftarten/Ubuntu-R.ttf",50)
	var bubbletext textboxen.Textbox = textboxen.New(220,150,305,120)
	var infotext textboxen.Textbox = textboxen.New(600,90,540,120)
	var punktenote textboxen.Textbox = textboxen.New(30,110,110,150)
	var durchgefallen textboxen.Textbox = textboxen.New(30,180,110,80)
	var ted texteditoren.Texteditor
	var exit vierecke.Viereck = vierecke.New(1080,90,1080,205,1170,205,1170,90)
	var punktespeicher []uint32
	punktespeicher = make([]uint32,len(texte))
	
//------------------Grafik-Elemente--------------------------------
	
	LadeBild(150,90,path2 + "Bilder/bubble2_flipped_400.bmp")
	LadeBildMitColorKey(50,400,path2 + "Bilder/SQLGame/Herk_200.bmp",255,255,255)
	
	SetzeFont(path2 + "Schriftarten/brlnsdb.ttf",60)
	Stiftfarbe(0,255,0)
	SchreibeFont(70,10,"SQL-Quest:")
	SetzeFont(path2 + "Schriftarten/brlnsr.ttf",50)
	SchreibeFont(390,20,"Explore the LWB-Adventure-World!")
	
	//punktenote.SetzeZeilenAbstand(5)
	punktenote.SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf")
	punktenote.SetzeSchriftgröße(20)
	punktenote.SchreibeText("Punkte: 0\n\nNote:")
	punktenote.Zeichne()
	
	durchgefallen.SetzeZeilenAbstand(5)
	durchgefallen.SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf")
	durchgefallen.SetzeSchriftgröße(20)
	durchgefallen.SetzeFarbe(255,0,0)
	durchgefallen.SchreibeText("nicht bestanden")
	durchgefallen.Zeichne()
	
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
	
	initialausgabe(ausgabe)
	
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

A:	for i:=1; i<len(texte); i++ {
		fmt.Println("aktuelles i:",i)
		
		var levelpunkte uint32 = 10
	
		for {

			taste, status, mausX, mausY := MausLesen1()
			if taste==1 && status==1 {
				if next.TesteXYPosInButton(mausX,mausY) {
					fmt.Println("Weiter angeklickt!")
					
					if i == 1 {
						//Beenden-Hinweis nach dem ersten Klick auf den next-Button
						Stiftfarbe(0,0,0)
						SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",14)
						SchreibeFont(30,280,"Beenden des")
						SchreibeFont(30,300,"Spiels:")
						SchreibeFont(30,320,"Tippe exit")
						SchreibeFont(30,340,"und ENTER!")
					}
					
					Stiftfarbe(255,255,255)
					Vollrechteck(220,115,305,170)
					bubbletext.SchreibeText(texte[aktuellerText])
					bubbletext.Zeichne()
					bubblehead.SchreibeText("Level "+fmt.Sprint(i))
					bubblehead.Zeichne()
					initialausgabe(ausgabe)
					//next.ZeichneButton()
					//Stiftfarbe(255,255,255)							//next-Button "ausgrauen"
					//Transparenz(50)
					//Vollrechteck(443,238,84,39)
					//Transparenz(0)
					
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
			
			if ted.GibString() == "exit" {
				i = 11
				break
			}
			
			for j:=0; j<4; j++ {
			
				//var eing string = ohneLeerzeichen(ted.GibString())
				//if  eing == eingTexte[i] || eing == eingTexte2[i] {
								
				if ohneLeerzeichen(ted.GibString()) == eingTexte[i] || ted.GibString() == "42" {
					//Stiftfarbe(255,255,255)
					//Vollrechteck(443,238,84,39)
					SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf",30)
					Stiftfarbe(255,255,255)
					Vollrechteck(220,235,315,40)
					Vollrechteck(100,500,100,100)
					LadeBildMitColorKey(50,400,path2 + "Bilder/SQLGame/herk_200.bmp",255,255,255)
					Stiftfarbe(0,255,0)
					SchreibeFont(230,240,"RICHTIG!!!  :-)")
					punkte = punkte + levelpunkte
					punktespeicher[i] = levelpunkte
					punktenoteSchreiben(punktenote,levelpunkte,punkte)
					//fmt.Println("Punkte:",fmt.Sprint(punkte))
					next.ZeichneButton()
					
					ausgabe.RahmenAn(true)
					ausgabe.SetzeRahmenFarbe(0,0,0)
					ausgabe.SetzeHöhe(190)
					ausgabe.HintergrundAn(true)
					ausgabe.SetzeHintergrundFarbe(255,255,255)
					ausgabe.SetzeFarbe(0,255,0)
					ausgabe.SchreibeText("")				
					ausgabe.Zeichne()
					Stiftfarbe(255,255,255)
					Vollrechteck(311,375,838,200)
					erzeugeAusgabe(i)
					//next.AktiviereButton()
					continue A
				} else {
					
					switch j {
						case 0:
						levelpunkte = levelpunkte - 1
						case 1:
						levelpunkte = levelpunkte - 2
						case 2:
						levelpunkte = levelpunkte - 3
						case 3:
						levelpunkte = levelpunkte - 4
					}
					//punktenoteSchreiben(punktenote,levelpunkte,punkte)
					
					//fmt.Println("Punkte:",fmt.Sprint(punkte))
					
					if j == 0 {											//Next-Button "ausgrauen"
						Stiftfarbe(255,255,255)
						Transparenz(50)
						Vollrechteck(443,238,84,39)
					}
					
					//ausgabe.Zeichnen()
					Stiftfarbe(255,255,255)
					Transparenz(0)
					Vollrechteck(311,371,838,218)						//alte Ausgabe überdecken
					//Vollrechteck(320,375,820,200)
					Vollrechteck(100,500,100,100)						//Herk-Mund überdecken wegen Transparenz
					LadeBildMitColorKey(50,400,path2 + "Bilder/SQLGame/herk_angry_red_200.bmp",255,255,255)
					
					ausgabe.RahmenAn(false)
					//ausgabe.SetzeRahmenFarbe(0,0,0)
					ausgabe.HintergrundAn(false)
					//ausgabe.SetzeHintergrundFarbe(255,255,255)
					ausgabe.SetzeHöhe(190)
					ausgabe.RahmenAn(true)
					ausgabe.SetzeRahmenFarbe(0,0,0)
					ausgabe.SetzeFarbe(255,0,0)
					ausgabe.SchreibeText(hilfeTexte[i][j])				
					ausgabe.Zeichne()
					
					if j<3 { erzeugeFehlerausgabe() }
					
					SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf",30)
					Stiftfarbe(255,0,0)
					
					if j<3 {
						SchreibeFont(225,240,"NOCHMAL!  :-(")
						ted = texteditoren.New(315,595,830,63,20,true)
					} else {
						Stiftfarbe(255,255,255)
						Vollrechteck(220,235,315,40)
						Stiftfarbe(255,0,0)
						SchreibeFont(225,240,"SCHADE!  :-(")
						next.ZeichneButton()
					}
				}
	
			}
			fmt.Println("Punkte:",fmt.Sprint(punkte))
			break					
		}		
	}
	
	punkte = punkte - 10
	fmt.Println("Punkte:",fmt.Sprint(punkte))
	
	//----------------- Endbildschirm --------------------------------------
	
	Stiftfarbe(255,255,255)
	Cls()
	
	//SpieleSound(path + "Sounds/the_force.wav")
	
	LadeBild(150,100,path + "Bilder/sprechblase_flipped_400.bmp")
	LadeBildMitColorKey(80,370,path + "Bilder/SQLGame/Herk_200.bmp",255,255,255)
	LadeBild(620,80,path2 + "Bilder/paper_500.bmp")
	LadeBild(960,520,path2 + "Bilder/certified_100.bmp")
	LadeBild(1080,90,path2 + "Bilder/Zurück-Symbol.bmp")
	//exit.SetKoordinaten(1080,80,1080,195,1170,195,1170,80)
	exit.SetzeFarbe(0,0,0)
	exit.AktiviereKlickbar()
	
	SetzeFont(path2 + "Schriftarten/brlnsdb.ttf",60)
	Stiftfarbe(0,255,0)
	SchreibeFont(70,10,"SQL-Quest:")
	SetzeFont(path2 + "Schriftarten/brlnsr.ttf",50)
	SchreibeFont(390,20,"Explore the LWB-Adventure-World!")
	
	Stiftfarbe(0,0,0)
	SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	SchreibeFont(295,140,"Du hast die")
	SchreibeFont(310,260,"erreicht!")
	SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	SchreibeFont(285,170,"Gesamtnote")
	SetzeFont(path2 + "Schriftarten/Starjedi.ttf",42)
	SchreibeFont(325,195,"1.0")												//TODO
	
	SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	for i:=1; i<len(texte); i++ {
		if i == 10 {
			if punktespeicher[i] < 10 {
				SchreibeFont(710,150+uint16((i-1)*40),"Level "+ fmt.Sprint(i) + ":    "+fmt.Sprint(punktespeicher[i])+" Punkte")
			} else {
				SchreibeFont(710,150+uint16((i-1)*40),"Level "+ fmt.Sprint(i) + ":   "+fmt.Sprint(punktespeicher[i])+" Punkte")
			}
		} else {
			if punktespeicher[i] < 10 {
				SchreibeFont(710,150+uint16((i-1)*40),"Level "+ fmt.Sprint(i) + ":     "+fmt.Sprint(punktespeicher[i])+" Punkte")
			} else {
				SchreibeFont(710,150+uint16((i-1)*40),"Level "+ fmt.Sprint(i) + ":    "+fmt.Sprint(punktespeicher[i])+" Punkte")
			}
		}
	}
	
	SchreibeFont(700,550,"----------------------")
	if punkte == 100 {
		SchreibeFont(710,580,"Gesamt:   "+fmt.Sprint(punkte)+" Punkte")
	} else {
		SchreibeFont(710,580,"Gesamt:    "+fmt.Sprint(punkte)+" Punkte")
	}
	
	TastaturLesen1()
	
	return punkte, notenberechnung(punkte)	
	
}
