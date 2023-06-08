package bugAttack

import (
		"../../Klassen/bugPackage"
		"time"
		"gfx"
		"math/rand"
		//"../../Klassen/textboxen"
		//"fmt"
	
		)

func BugAttack() (float32,uint32){
	rand.Seed(time.Now().UnixNano())		// Seed für Zufallszahlen
	//-----------------initialisiere gfx-Fenster-----------------------	
	if ! gfx.FensterOffen() {
		gfx.Fenster(1200,700)
	}
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	gfx.Stiftfarbe(0,255,0)
	go bugPackage.CursorPos()
	bugPackage.LevelIntro()
	bugPackage.Startbildschirm() 	
	bugPackage.LevelTutorial()
	bugPackage.Level1()
	bugPackage.Level2()
	bugPackage.Level3()
	//fmt.Println("ENDE")
	bugPackage.EndbildschirmReal()
	endN,endP := bugPackage.GibErgebnis()
	return endN,endP
}
