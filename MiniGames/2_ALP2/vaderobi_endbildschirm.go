//Annalena Cyriacus
//Juni 2023
//LWB-Adventure: Vaderobi-Minigame (ALP2)

package main

import (
	. "gfx"
	"fmt"
	"../../Klassen/vierecke"
	//"../../Klassen/buttons"
)

var path string = "../../"
var path2 string = "./" //MiniGames/2_ALP2/"

var exit vierecke.Viereck = vierecke.New(1080,30,1080,145,1170,145,1170,30)

func nichtzeugnisBildschirm() {
	
	//var spielBeenden buttons.Button = buttons.New()

	Stiftfarbe(255,255,255)
	Vollrechteck(0,0,1200,700)
	
	//LadeBild(250,50,"./Bilder/Tür5.bmp")
	
	//SpieleSound(path + "Sounds/the_force.wav")
	
	LadeBild(150,100,path + "Bilder/sprechblase_flipped_400.bmp")
	LadeBildMitColorKey(100,350,path + "Bilder/Darth_200.bmp",255,255,255)
	LadeBild(600,80,path + "Bilder/MainGame/zertifikat.bmp")
	LadeBild(940,510,path + "Bilder/certified_100.bmp")
	LadeBild(1080,30,path + "Bilder/Zurück-Symbol.bmp")
	exit.SetzeFarbe(0,0,0)
	exit.Zeichnen()
	exit.AktiviereKlickbar()
	
	Stiftfarbe(0,0,0)
	SetzeFont(path + "Schriftarten/brlnsr.ttf",42)
	SchreibeFont(50,10,"Herzlichen Glückwunsch zum erfolgreich absolvierten LWB-Adventure!!!")
	Stiftfarbe(0,0,0)
	SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	SchreibeFont(295,145,"Du hast den")
	SchreibeFont(310,265,"erreicht!")
	SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	SchreibeFont(230,175,"Notendurchschnitt")
	SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
	SchreibeFont(325,200,"1.0")
	
	SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	for i:=1; i<7; i++ {
		SchreibeFont(650,150+uint16((i-1)*68),"Level "+ fmt.Sprint(i) + ":   xx Punkte")
		SchreibeFont(650,175+uint16((i-1)*68),"           Note x.x")
	}
	SchreibeFont(640,130+uint16(6*70),"----------------------")
	SchreibeFont(650,160+uint16(6*70),"Gesamt:    xx Punkte")
	
}

func sqlEndeBildschirm() {
	
	//SpieleSound(path + "Sounds/the_force.wav")
	
	LadeBild(150,100,path + "Bilder/sprechblase_flipped_400.bmp")
	LadeBildMitColorKey(80,370,path + "Bilder/SQLGame/Herk_200.bmp",255,255,255)
	//LadeBildMitColorKey(50,400,path2 + "Bilder/SQLGame/Herk_200.bmp",255,255,255)
	LadeBild(620,80,path + "Bilder/paper_500.bmp")
	LadeBild(960,520,path + "Bilder/certified_100.bmp")
	LadeBild(1080,80,path + "Bilder/Zurück-Symbol.bmp")
	exit.SetKoordinaten(1080,80,1080,195,1170,195,1170,80)
	exit.SetzeFarbe(0,0,0)
	exit.Zeichnen()
	exit.AktiviereKlickbar()
	
	SetzeFont(path + "Schriftarten/brlnsdb.ttf",60)
	Stiftfarbe(0,255,0)
	SchreibeFont(70,10,"SQL-Quest:")
	SetzeFont(path + "Schriftarten/brlnsr.ttf",50)
	SchreibeFont(390,20,"Explore the LWB-Adventure-World!")
	
	Stiftfarbe(0,0,0)
	SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	SchreibeFont(295,140,"Du hast die")
	SchreibeFont(310,260,"erreicht!")
	SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	SchreibeFont(285,170,"Gesamtnote")
	SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
	SchreibeFont(325,195,"1.0")
	
	SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	for i:=1; i<11; i++ {
		SchreibeFont(710,150+uint16((i-1)*40),"Level "+ fmt.Sprint(i) + ":   xx Punkte")
	}
	
	SchreibeFont(700,550,"----------------------")
	SchreibeFont(710,580,"Gesamt:    xx Punkte")
	//SchreibeFont(700,130+uint16(6*70),"----------------------")
	//SchreibeFont(710,160+uint16(6*70),"Gesamt:    xx Punkte")
	
}

func vaderobiEndeBildschirm() {
	
	SpieleSound(path + "Sounds/the_force.wav")
	
	LadeBild(150,100,path + "Bilder/sprechblase_flipped_400.bmp")
	LadeBildMitColorKey(100,350,path + "Bilder/Darth_200.bmp",255,255,255)
	LadeBild(620,80,path + "Bilder/paper_500.bmp")
	LadeBild(960,520,path + "Bilder/certified_100.bmp")
	LadeBild(1080,30,path + "Bilder/Zurück-Symbol.bmp")
	exit.SetzeFarbe(0,0,0)
	exit.Zeichnen()
	exit.AktiviereKlickbar()
	
	Stiftfarbe(0,255,0)
	SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
	SchreibeFont(330,10,"Super - ALP - Escape")
	Stiftfarbe(0,0,0)
	SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	SchreibeFont(295,140,"Du hast die")
	SchreibeFont(310,260,"erreicht!")
	SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	SchreibeFont(285,170,"Gesamtnote")
	SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
	SchreibeFont(325,195,"1.0")
	
	SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	for i:=1; i<7; i++ {
		SchreibeFont(710,150+uint16((i-1)*68),"Level "+ fmt.Sprint(i) + ":   xx Punkte")
		SchreibeFont(710,175+uint16((i-1)*68),"           Note x.x")
	}
	SchreibeFont(700,130+uint16(6*70),"----------------------")
	SchreibeFont(710,160+uint16(6*70),"Gesamt:    xx Punkte")
	
}

func main() {
	Fenster(1200,700)
	
	//vaderobiEndeBildschirm()
	sqlEndeBildschirm()
	//nichtzeugnisBildschirm()
	
	TastaturLesen1()
}
