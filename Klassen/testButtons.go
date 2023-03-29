// ADT button - Test

// Martin Seiß    29.3.2023

package main


import "./buttons"
import "gfx"

/*
 * 
	AendereBeschriftung(text string) 

*/

func main() {
	
	var wtkStatus bool = true			// Status of Winnie the K 
	var wtkx,wtky uint16 = 400,100

	var b1 buttons.Button =	buttons.New(100,100,220,40,255,255,0,true, "Aktiviere Button 2")
	var b2 buttons.Button =	buttons.New(100,200,100,40,255,255,0,false, "Button 2")
	var endebutton buttons.Button =	buttons.New(100,450,100,40,255,255,0,true, "Ende")
	var wtk buttons.Button = buttons.New(wtkx,wtky,380,375,255,255,0,true, "")
	
	gfx.Fenster(800,500)
	gfx.SetzeFont ("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",20)

	// Zeichne Buttons
	b1.ZeichneButton()
	b2.ZeichneButton()
	gfx.LadeBild(wtkx,wtky,"../Bilder/WtheK.bmp")
	gfx.Rechteck(wtkx,wtky,380,375)	
	endebutton.ZeichneButton()
	

	// Mausabfrage
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
		if taste==1 && status==1 {
			if b1.TesteXYPosInButton(mausX,mausY) {
				if b2.GibAktivitaetButton() {
					b2.DeaktiviereButton()
					b1.AendereBeschriftung("Aktiviere Button 2")
				} else {
					b2.AktiviereButton()
					b1.AendereBeschriftung("Deaktiviere Button 2")
				}
				b2.ZeichneButton()
				b1.ZeichneButton()
			}
			if b2.TesteXYPosInButton(mausX,mausY) {
				gfx.SchreibeFont(100,300,"Getroffen!!!")
			} 
			if wtk.TesteXYPosInButton(mausX,mausY) {
				if wtkStatus {
					gfx.LadeBild(400,100,"../Bilder/WtheK_black_sad.bmp")
					gfx.Rechteck(wtkx,wtky,380,375)	
					wtkStatus = false
				} else {
					gfx.LadeBild(400,100,"../Bilder/WtheK.bmp")
					gfx.Rechteck(wtkx,wtky,380,375)	
					wtkStatus = true
				}
			}
			if endebutton.TesteXYPosInButton(mausX,mausY) {
				break
			} 
		}
	}
	
	// gfx.TastaturLesen1()

}

