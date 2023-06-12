package fachjargon

import (
	"gfx"
	"../../Klassen/fachjargonPackage"
	"../../Klassen/audioloops"
	"time"
	"math/rand"
	
)



const breite uint16 = 1200
const höhe uint16 = 700


func FachJargon() (float32,uint32){
	c :=make(chan bool)						// Kanal zum Beenden der Tastatur-Go-Routine
	rand.Seed(time.Now().UnixNano())		// Seed für Zufallszahlen
	if ! gfx.FensterOffen() {
			gfx.Fenster(breite,höhe)
		}

	go audioloops.LoopInfinite("Sounds/Music/30s_Surf.wav",40000)
	fachjargonPackage.StartBildschirm()
	go fachjargonPackage.TastaturEingabe(c)		// Startet das Lesen der Tastatur
	
	
	fachjargonPackage.IntroBildschirm()		// Zeigt Intro
	for _,l := range fachjargonPackage.LevelArray {
		if fachjargonPackage.SpielBeendet == true {break}
		l()
	}
	
	fachjargonPackage.Endbildschirm()
	audioloops.StoppeAudio()	
	
	
	fachjargonPackage.SpielBeendet = false // damit man es nochmal spielen kann
	c<-true
	return fachjargonPackage.EndN,fachjargonPackage.EndP

}

