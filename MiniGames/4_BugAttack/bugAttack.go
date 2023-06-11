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
	c := make(chan bool) 		// Zum beenden der Go-Routine
	rand.Seed(time.Now().UnixNano())		// Seed für Zufallszahlen
	//-----------------initialisiere gfx-Fenster-----------------------	
	if ! gfx.FensterOffen() {
		gfx.Fenster(1200,700)
	}
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	gfx.Stiftfarbe(0,255,0)
	go bugPackage.TastaturEingabe(c)
	
	
	bugPackage.BugAttackIntro()
	bugPackage.Startbildschirm() 
	/*	
	bugPackage.LevelTutorial()
	bugPackage.Level1()
	bugPackage.LevelTutorial()
	bugPackage.Level2()
	bugPackage.Level3()
	*/
	//fmt.Println("ENDE")
	
	// Starte nacheinander die Funktionen der Level
	for _,f:= range bugPackage.LevelArray {
		if bugPackage.SpielBeendet == true {break} 	// Spiel soll beendet werden? Keine weiteren Level
		f()
	}
	
	bugPackage.Endbildschirm()
	gfx.StoppeAlleSounds()
	endN,endP := bugPackage.GibErgebnis()
	c<-true						// Beendet Tastatureingabe
	return endN,endP
}
