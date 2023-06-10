package bugPackage

import (
		"gfx"
		"time"
		)
		
//var path string = "../../"
var path string = ""

var amoebiusArray [13]string = [13]string{path+"Bilder/BugAttack/Amoebius40.bmp",path+"Bilder/BugAttack/Amoebius50.bmp",
	path+"Bilder/BugAttack/Amoebius60.bmp",path+"Bilder/BugAttack/Amoebius70.bmp",path+"Bilder/BugAttack/Amoebius80.bmp",
	path+"Bilder/BugAttack/Amoebius90.bmp",path+"Bilder/BugAttack/Amoebius100.bmp",
	path+"Bilder/BugAttack/Amoebius90.bmp",path+"Bilder/BugAttack/Amoebius80.bmp",
	path+"Bilder/BugAttack/Amoebius70.bmp",path+"Bilder/BugAttack/Amoebius60.bmp",path+"Bilder/BugAttack/Amoebius50.bmp",
	path+"Bilder/BugAttack/Amoebius40.bmp"}
	
var bugAnimationArray [3]string = [3]string{path+"Bilder/BugAttack/Bug1.bmp",path+"Bilder/BugAttack/Bug2.bmp",path+"Bilder/BugAttack/Bug3.bmp"}

func amoebiusAndBugAnimation() {
	var i uint16
	for {
		// Bewege Amoebius nacht rechts
		for i=0;i<1000;i+=130{
			// Iteriere über jedes Bild für die Animation, dabei wir das Bild immer leicht verschoben
			for j,b :=  range amoebiusArray {
				// Wenn ein BeendenSignal kommt, return 
				select {
					case <-quit:
						return
					default: 
						gfx.UpdateAus()
						gfx.Stiftfarbe(0,0,0)
						gfx.Restaurieren(0,0,1200,700)
						gfx.LadeBildMitColorKey(0+i+uint16(j)*10,50,b,0,0,0)
						gfx.LadeBildMitColorKey(0+i+uint16(j)*10+100,50,bugAnimationArray[j%3],0,0,0)
						gfx.UpdateAn()
						time.Sleep(1e8)
				}
			}
		}
	}
}

/*
func bugLevelAnimation() {
	var i uint16
	for {
		for i=0;i<1000;i+=130 {
			for j,b := range bugAnimationArray {
				select {
					case <-quit:
						return
					default:
						gfx.LadeBildMitColorKey(0+i+uint16(j)*40,150,b,0,0,0)
						time.Sleep(4e8)
				}
			}
		}
	}
}
*/
