/* Philipp Liehm
 * Juni 2023
 * LWB-Adventure: BugAttack
 * 
 * ADT zum Erstellen von Ladebalken in BugAttack
 */
 
package bugPackage

import (
		"gfx"
		"time"
		)




// ADT ladebalken zur Verwendung in BugAttack
type ladebalken struct {
	wertAdresse *uint16
	x uint16
	y uint16
	r,g,b uint8
	taste string
	cdlänge	uint16
	sound string
}

// Vor.: Es existiert eine globale Variable and der gegebenen Adresse
// Erg.: ein neuer Ladebalken ist geliefert
func newLadebalken(wertAdresse *uint16,x,y uint16,r,g,b uint8, taste string, cdl uint16,sound string) *ladebalken {
	l := new(ladebalken)
	l.x,l.y,l.r,l.g,l.b,l.taste = x,y,r,g,b,taste			// Ein ladebalken hat eine Position, Farbe, zugehörige Taste
	l.wertAdresse = wertAdresse								// Verknüpft mit dem Wert einer globalen Variable
	l.cdlänge = cdl											// Länge des Cooldowns
	l.sound = sound											// Sound-Datei wenn Fähigkeit des Ladebalkens verfügbar
	return l
}


// Vor.: keine
// Eff.: ladebalken ist gezeichnet
func (l *ladebalken) zeichne() {
	gfx.Stiftfarbe(l.r,l.g,l.b)
	gfx.SetzeFont("Schriftarten/ltypeb.ttf",20)
	gfx.SchreibeFont(l.x,l.y,l.taste)					// Name der Taste zu dem der Balken gehört ist geliefert
	gfx.Vollrechteck(l.x+20,l.y,10* *l.wertAdresse,20)	// Balken, soweit ausgefüllt wie er fortgeschritten ist
	gfx.Rechteck(l.x+20,l.y,100,20)						// Rahmen um den Balken
}

// Vor.: keine
// Eff.: Der Cooldown des Ladebalken hat sich aktualisiert
func (l *ladebalken) cooldown() {
	*l.wertAdresse = 0		// Setze Cooldown auf 0
	for lvlLäuft {			// solange ich noch im Level bin
		if *l.wertAdresse != 0 {		// Wenn der CD nicht 0 ist mache weiter in der Schleife
			time.Sleep(1e8)		
			continue		
		}			
		for i:=uint16(0);i<=10;i++ {			// Lade die Fähigkeit 10 mal
			if lvlLäuft == false {	// Falls das Level beendet ist, beende diese Methode
				return
				}	
			//fmt.Println("CD: ",l.cdlänge," Wert: ",*l.wertAdresse)
					
			*l.wertAdresse = i		// Überschreibe den Wert an der Adresse
			if *l.wertAdresse == 10	{		// wenn voll geladen, spiele Sound um anzuzeigen, dass Fähigkeit verfügbar ist
				gfx.SpieleSound(l.sound)
			}
			time.Sleep(time.Duration(l.cdlänge)*1e8)	// warte bis zum nächsten Ladeschritt
		}
	}
		time.Sleep(1e8)
}	

// Zeichnet alle Ladebalken im alleLadebalken-Slice
func zeichneAlleLadebalken() {
	for _,l:=range alleLadebalken {
		if l!=nil {l.zeichne()}
	}
}

// Löscht Ladebalken eines Levels
func entferneAlleLadebalken() {
	for i,_:=range alleLadebalken {
		alleLadebalken[i] = nil
	}
}

