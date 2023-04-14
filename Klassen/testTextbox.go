package main

import (
	"./textboxen"
	"gfx"
	//"strings"
	//"fmt"
	
)

func main() {
	// Öffne gfx-Fenster
	gfx.Fenster(1000,800)
	// Neues TexboxObjekt
	var tb1,tb2,tb3 textboxen.Textbox
	tb1 = textboxen.New(100,100,200,200)
	tb1.SetzeSchriftgröße(30)
	tb1.SetzeZentriert()
	tb1.SchreibeText("Das ist ein Test! Crazy! Und hier kommt noch mehr! Und noch mehr!")
	tb1.Zeichne()
	
	tb2 = textboxen.New(300,100,200,200)
	tb2.SetzeSchriftgröße(10)
	tb2.SchreibeText("Das ist ein Test! Crazy! Und hier kommt noch mehr! Und noch mehr!")
	tb2.SetzeFarbe(0,0,255)
	tb2.SetzeZeilenAbstand(10)
	tb2.Zeichne()
	
	tb3 = textboxen.New(500,100,200,200)
	tb3.SetzeSchriftgröße(20)
	tb3.SchreibeText("Das ist ein Test! Crazy! Und hier kommt noch mehr! Und noch mehr!")
	tb3.SetzeFont("../Schriftarten/Prisma.ttf")
	tb3.Zeichne()
	
	
	gfx.TastaturLesen1()
	
}
	
