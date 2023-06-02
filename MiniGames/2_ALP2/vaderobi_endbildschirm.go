//Annalena Cyriacus
//Juni 2023
//LWB-Adventure: Vaderobi-Minigame (ALP2)

package main

import (
	. "gfx"
	"fmt"
)

var path string = "../../"
var path2 string = "./" //MiniGames/2_ALP2/"

func main() {
	Fenster(1200,700)
	
	SpieleSound(path + "Sounds/the_force.wav")
	
	LadeBild(150,100,path + "Bilder/sprechblase_flipped_400.bmp")
	LadeBildMitColorKey(100,350,path + "Bilder/Darth_200.bmp",255,255,255)
	LadeBild(620,80,path + "Bilder/paper_500.bmp")
	LadeBild(960,520,path + "Bilder/certified_100.bmp")
	LadeBild(1080,30,path + "Bilder/Zurück-Symbol.bmp")
	
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

	TastaturLesen1()
}
