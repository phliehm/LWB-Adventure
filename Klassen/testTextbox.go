package main

import (
	"./textboxen"
	"gfx"
	//"strings"
	"fmt"
	"time"
	
)

func main() {
	// Öffne gfx-Fenster
	gfx.Fenster(1000,800)
	gfx.SetzeFont("../Schriftarten/Ubuntu-B.ttf",20)
	// Neues TexboxObjekt
	var tb1 textboxen.Textbox
	tb1 = textboxen.New(100,100,200,200)
	tb1.SetzeSchriftgröße(30)
	tb1.SetzeZentriert()
	tb1.SchreibeText("Das ist ein Test! Crazy! Und hier kommt noch mehr! Und noch mehr!")
	tb1.Zeichne()
	time.Sleep(2e9)
	
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	tb1.SetzePosition(300,300)
	tb1.SetzeBreite(400)
	tb1.SetzeHöhe(400)
	tb1.SetzeZeilenAbstand(20)
	tb1.SetzeFarbe(255,0,0)
	fmt.Println("X: ",tb1.GibX(),"  Y: ", tb1.GibY())
	fmt.Println("Breite: ",tb1.GibBreite(),"  Höhe: ", tb1.GibHöhe())
	fmt.Println("Zeilenabstand: ",tb1.GibZeilenAbstand())
	fmt.Println("Font: ", tb1.GibFont(), "   Schriftgröße: ", tb1.GibSchriftgröße())
	tb1.Zeichne()
	time.Sleep(2e9)	
	
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	tb1.SetzeZentriert()
	tb1.RahmenAn(true)
	tb1.Zeichne()
	time.Sleep(1e9)
	
	
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	tb1.SetzeLinksbündig()
	tb1.SetzeRahmenFarbe(0,0,255)
	tb1.RahmenAn(false)
	tb1.Zeichne()
	time.Sleep(1e9)
	
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	tb1.HintergrundAn(true)
	tb1.SetzeHintergrundFarbe(255,255,0)
	tb1.Zeichne()
	
	
	
	gfx.TastaturLesen1()
	
}
	


