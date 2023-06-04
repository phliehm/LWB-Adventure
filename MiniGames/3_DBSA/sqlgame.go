//Annalena Cyriacus
//Juni 2023
//LWB-Adventure: Minigame "SQL-Quest"

package sqlgame

import (
	. "gfx"
	//"fmt"
	"../../Klassen/buttons"
	"../../Klassen/textboxen"
)

var path string = "../../"
var path2 string = "./" //MiniGames/2_ALP2/"

//func SQLgame() (float32,uint32) {
func SQLgame() {
	
//------------------Variablen--------------------------------------
	var weiter buttons.Button = buttons.New(400,350,150,45,0,255,0,true,"weiter")
	var datainfo textboxen.Textbox = textboxen.New(350,650,800,50)
	
//------------------Grafik-Elemente--------------------------------
	Fenster(1200,700)
	LadeBild(150,50,path2 + "Bilder/bubble2_flipped_450.bmp")
	LadeBildMitColorKey(50,400,path2 + "Bilder/Herk_200.bmp",255,255,255)
	
	weiter.ZeichneButton()
	weiter.SetzeFont(path2 + "Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf")
	datainfo.RahmenAn(true)
	datainfo.HintergrundAn(true)
	datainfo.SetzeRahmenFarbe(0,0,0)
	datainfo.SetzeHintergrundFarbe(0,0,0)
	datainfo.SetzeFarbe(255,255,255)
	datainfo.SchreibeText("Tabelle 1			Tabelle 2			Tabelle 3")
	datainfo.Zeichne()
	


	TastaturLesen1()
	
}
