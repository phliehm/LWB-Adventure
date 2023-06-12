/* Philipp Liehm
 * Juni 2023
 * LWB-Adventure: BugAttack
 */


package bugAttack

import (
		"../../Klassen/bugPackage"
		"../../Klassen/audioloops"
		"time"
		"gfx"
		"math/rand"
		)

// Main-Spiel Funktion, wird von LWB-Adventure aufgerufen
func BugAttack() (float32,uint32){
	c := make(chan bool) 				// Zum beenden der Go-Routine
	rand.Seed(time.Now().UnixNano())		// Seed für Zufallszahlen
	//-----------------initialisiere gfx-Fenster-----------------------	
	if ! gfx.FensterOffen() {			// Test oder MainGame
		gfx.Fenster(1200,700)
	}
	gfx.Stiftfarbe(0,0,0)				// Schwarzes Fenster machen
	gfx.Vollrechteck(0,0,1200,700)
	gfx.Stiftfarbe(0,255,0)				
	go audioloops.LoopInfinite("Sounds/Music/bugAttack.wav",48000)		// Musik läuft das ganze Spiel und wird immer wieder neu gestartet
	go bugPackage.TastaturEingabe(c)	// Starte Tastatur-Eingabe
	
	
	bugPackage.BugAttackIntro()			// Intro-Bildschirm mit Animation
	bugPackage.Startbildschirm() 		// Einstiegstext
	
	// Starte nacheinander die Funktionen der Level
	for _,f:= range bugPackage.LevelArray {
		if bugPackage.SpielBeendet == true {break} 	// Spiel soll beendet werden? Keine weiteren Level
		f()								// Starte aktuelles Level
	}
	
	bugPackage.Endbildschirm()			// Zertifikat
	
	endN,endP := bugPackage.GibErgebnis()		// Hole die Ergebniswerte
	c<-true						// Beendet Tastatureingabe
	audioloops.StoppeAudio()		// Stoppe die Musik
	bugPackage.SpielBeendet = false 	// Damit man das Spiel nochmal spielen kann
	return endN,endP			// übergebe Ergebnisse an LWB-Adventure
}
