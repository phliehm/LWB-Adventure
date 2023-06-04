//Annalena Cyriacus
//Juni 2023
//LWB-Adventure: Minigame "SQL-Quest"

package sqlgame

import (
	. "gfx"
	//"fmt"
	"../../Klassen/buttons"
	"../../Klassen/textboxen"
	"../../Klassen/texteditoren"
)

var path string = ""
var path2 string = "./" //MiniGames/2_ALP2/"

//----------------------Hilfsfunktionen----------------------------

func angeklickt(button buttons.Button) bool {
	taste, status, mausX, mausY := MausLesen1()
	if taste==1 && status==1 {
		if button.TesteXYPosInButton(mausX,mausY) {
			return true
		}
	}
	return false
}

//func SQLgame() (float32,uint32) {
func SQLgame() {
	
//-----------------initialisiere gfx-Fenster-----------------------	
	Fenster(1200,700)
	SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	
//------------------Variablen--------------------------------------
	var weiter buttons.Button = buttons.New(445,240,80,35,0,255,0,true,"weiter")
	var datainfo textboxen.Textbox = textboxen.New(600,270,540,70)
	var ausgabe textboxen.Textbox = textboxen.New(320,375,820,45)
	SetzeFont(path2 + "Schriftarten/Ubuntu-B.ttf",50)
	var bubblehead textboxen.Textbox = textboxen.New(220,115,305,155)
	SetzeFont(path2 + "Schriftarten/Ubuntu-R.ttf",50)
	var bubbletext textboxen.Textbox = textboxen.New(220,115,305,155)
	var infotext textboxen.Textbox = textboxen.New(600,95,540,120)
	var ted texteditoren.Texteditor
	
//------------------Grafik-Elemente--------------------------------
	
	LadeBild(150,90,path2 + "Bilder/bubble2_flipped_400.bmp")
	LadeBildMitColorKey(50,400,path2 + "Bilder/Herk_200.bmp",255,255,255)
	
	SetzeFont(path2 + "Schriftarten/brlnsdb.ttf",60)
	Stiftfarbe(0,255,0)
	SchreibeFont(140,10,"SQL-Quest:")
	SetzeFont(path2 + "Schriftarten/brlnsr.ttf",50)
	SchreibeFont(460,20,"Explore the world of LWB!")
	
	weiter.ZeichneButton()
	weiter.SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf")
	
	//infotext.SetzeFarbe(0,0,0)
	infotext.SetzeZeilenAbstand(5)
	infotext.SetzeSchriftgröße(18)
	infotext.SchreibeText("Herk hat eine Datenbank erstellt, die alle Geheimnisse der LWB-Welt enthält. Würdig darauf zuzugreifen sind nur jene, die SQL beherrschen.\nDrum schärfe Deinen Verstand und gib die richtigen Anfragen ein, um die Geheimnisse zu ergründen!\n\nFolgende Tabellen sind in seiner Datenbank enthalten:")
	infotext.Zeichne()
	
	datainfo.RahmenAn(true)
	datainfo.HintergrundAn(true)
	datainfo.SetzeRahmenFarbe(0,0,0)
	datainfo.SetzeHintergrundFarbe(0,0,0)
	datainfo.SetzeFarbe(255,255,255)
	datainfo.SetzeZeilenAbstand(5)
	datainfo.SchreibeText("Tabelle 1\nTabelle 2\nTabelle 3")
	datainfo.Zeichne()
	
	ausgabe.RahmenAn(true)
	ausgabe.SetzeRahmenFarbe(0,0,0)
	ausgabe.SetzeFarbe(0,255,0)
	ausgabe.SetzeZeilenAbstand(5)
	ausgabe.SchreibeText("Hier erscheint das Ergebnis Deiner SQL-Anfrage, sobald Du sie eingegeben und mit ENTER ausgeführt hast!")
	ausgabe.Zeichne()
	
	//bubbletext.RahmenAn(true)
	//bubbletext.SetzeRahmenFarbe(0,0,0)
	bubbletext.SetzeZeilenAbstand(3)
	bubbletext.SetzeSchriftgröße(18)
	bubbletext.SchreibeText("\n\nUm die DBSA-Prüfung zu bestehen,\nmusst Du Dein SQL-Wissen nutzen,\num alle Geheimnisse der LWB-Welt\nzu enthüllen!")
	bubbletext.Zeichne()
	bubblehead.SchreibeText("Willkommen zum SQL-Quest!")
	bubblehead.Zeichne()
	
	
	
//----------------Eingabe-Verarbeitung----------------------------
	ted = texteditoren.New(315,440,830,218,20,true)
	
	if angeklickt(weiter) {
		bubbletext.SchreibeText("neuerText")
	}
	
	ted.GibString()
	//for {
		//switch ted.GibString() {
			//case 
	
	
	TastaturLesen1()
	
}
