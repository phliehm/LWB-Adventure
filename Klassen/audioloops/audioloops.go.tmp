package audioloops

/* Author: Philipp Liehm
 * Datum: 24.04.2023
 * 
 * Dieses Paket stellt Funktionen zum Abspielen von Audiodateien zur Verfügung.
 * Damit können Audiodateien wiederholt abgespielt werden. 
 * 
 */

var musikAn bool

import (
	"gfx"
	"time"
	
)
// Vor.: Ein gfx Fenster muss geöffnet sein,
//		 Idealerweise ist die Länge des Loops ein Vielfaches (oder etwas kleiner) der Audiodatei
// Eff.: Die "audiodatei" der Länge "dateiLänge" (in ms) wird so lange abgespielt
//		 bis die Zeit "loopLänge" (in ms) abgelaufen ist
func LoopDuration(audiodatei string, dateiLänge uint, loopLänge uint) () {
	
	// Erstelle einen neuen Timer
	t := time.NewTimer(time.Duration(loopLänge)*time.Millisecond)
	musikAn = true	
	for musikAn {
		select {
		// Timer ist abgelaufen
		case <-t.C:			
			//fmt.Println("Zeit ist abgelaufen:(")
			return
		// Timer läuft noch
		default:
			gfx.SpieleSound(audiodatei)
			// Warte bis Datei zu Ende ist, leider kann man nicht nachprüfen wann die Datei zu Ende ist.
			time.Sleep(time.Duration(dateiLänge)*time.Millisecond)
		}
	}
	musikAn = false
	
}

func StoppeAudio () {
	gfx.StoppeAlleSounds()
	musikAn = false		// Damit Loop nicht wieder neu anfängt
}
