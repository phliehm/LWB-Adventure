package main

import (
		"../../Klassen/bugPackage"
		"time"
		"gfx"
		"math/rand"
		//"../../Klassen/textboxen"
		//"fmt"
	
		)

func main() {
	rand.Seed(time.Now().UnixNano())		// Seed für Zufallszahlen
	gfx.Fenster(1200,700)
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	gfx.Stiftfarbe(0,255,0)
	go bugPackage.CursorPos()
	bugPackage.LevelIntro()
	bugPackage.Startbildschirm() 	
	bugPackage.LevelTutorial()
	bugPackage.Level0()
	//bugPackage.Level1()
	bugPackage.Level2()
	bugPackage.EndbildschirmReal()
}
