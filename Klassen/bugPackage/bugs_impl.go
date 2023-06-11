/* Philipp Liehm
 * Juni 2023
 * LWB-Adventure: BugAttack
 * 
 * Implementierung der Klasse bug
 * 
 */

package bugPackage

import (
		"gfx"
		"fmt"
		"time"
		"math/rand"
		"math"
		)

// Feld zum zeichnen eines Bugs. Die Koordinaten sind relativ zur linken oberen Ecke 
var bug1Shape [21][2] uint16 = [21][2]uint16{{0,0},{6*zB,0},
												{1*zB,1*zH},{5*zB,1*zH},
												{2*zB,2*zH},{3*zB,2*zH},{4*zB,2*zH},
												{0,3*zH},{1*zB,3*zH},{2*zB,3*zH},{3*zB,3*zH},{4*zB,3*zH},{5*zB,3*zH},{6*zB,3*zH},
												{2*zB,4*zH},{3*zB,4*zH},{4*zB,4*zH},
												{1*zB,5*zH},{5*zB,5*zH},
												{0,6*zH},{6*zB,6*zH}}

// Datentyp bug, Gegener im Spiel BugAttack
type bug struct {
	x,y uint16		// x,y geben die linke obere Ecke des Bugs an
	alive bool 		// Bug lebt noch
	dying uint16	// Für die Animation zum Sterben
	a uint16		// AnimatiosSchritt
	stirbt bool 	// true wenn Bug getroffen wurde und stirbt
	typ uint8 		// Art des Bugs
	speed int 		// Geschwindigkeit des Bugs, also wie weit bewegt er sich pro Schritt
	nervosität int // wie oft bewegt sich der Bug
	r,g,b uint8		// Farbe des Bugs
}

// Vor.: Es muss eine Welt aus dem BugPackage geben
// Eff.: Instanziiert neuen Bug mit den gegeben Parametern
func NewBug(x,y uint16) *bug{	// x,y sind hier Koordinaten auf dem Raster
	var b *bug				// 
	b = new(bug)
	b.alive=true			// belebt Bug
	b.speed = 5				// Standard Geschwindigkeit
	b.nervosität = 1		// Standard Nervosität
	b.x = x*zB				// x-Position in der Welt
	b.y = y*zH+y_offset*zH	// y-Position in der Welt, y_offset ist nötig weil die Welt nicht in der linken oberen Ecke beginnt	
	b.r,b.g,b.b = 255,0,0	// Standardfarbe ist Rot
	return b
}

// Vor.: keine
// Eff.: Zeichnet Bug mit seinen Animationen
func (b *bug) zeichneBug() {
	x:= b.x					// nur zur Abkürzung
	y:= b.y
	// Bug Größe: Breite: 7*9 , Höhe 7*14
	animation1 := func() {			// Animationsschritt 1: Körper
		sr,sg,sb = b.r,b.g,b.b		// Setze globale Farbe auf Bug Farbe
		male_0(x+2*zB,y+2*zH)
		male_0(x+3*zB,y+2*zH)
		male_0(x+4*zB,y+2*zH)
		male_0(x+2*zB,y+3*zH)
		male_0(x+3*zB,y+3*zH)
		male_0(x+4*zB,y+3*zH)
		male_0(x+2*zB,y+4*zH)
		male_0(x+3*zB,y+4*zH)
		male_0(x+4*zB,y+4*zH)
	}
			
	animation2 := func() {			// Animationsschritt 2, Arme/Beine
		sr,sg,sb = b.r,b.g,b.b		
		male_1(x+1*zB,y+1*zH)
		male_1(x+5*zB,y+1*zH)
		
		male_1(x+1*zB,y+3*zH)
		male_1(x+5*zB,y+3*zH)
		
		male_1(x+1*zB,y+5*zH)
		male_1(x+5*zB,y+5*zH)
	}
	animation3 := func() {		// Animationsschritt 3, Hände/Füße
		sr,sg,sb = b.r,b.g,b.b		
		male_1(x,y)
		male_1(x+6*zB,y)

		male_1(x,y+3*zH)
		male_1(x+6*zB,y+3*zH)
		male_1(x,y+6*zH)
		male_1(x+6*zB,y+6*zH)
	}

	implosion:=func(ix,iy uint16) {		// Hilfsfunktion für Animation 4, Implosion. Bug wird nach Abschuss zersetzt wird
		gfx.Punkt(ix+uint16(rand.Intn(int(zB))),iy+uint16(rand.Intn(int(zH))))	// Es werden einfach schwarze Punkte gezeichnet
	}
	animation4 := func() {				// Implosion
		gfx.Stiftfarbe(0,0,0)			// schwarz zersetzt
		for i:=0;i<10;i++ {				// 10 Schritte der Implosion
			for k:=range bug1Shape {	// Es wird über jede Zahl aus der der Bug besteht iteriert
				implosion(x+bug1Shape[k][0],y+bug1Shape[k][1])	// führe Implosion aus
			}
		}
	}
	switch b.a { 						// Abhängig davon bei welchem Animationsschritt der Bug gerade ist, wird dieser gezeichnet
		case 1: 
			animation1()
		case 2: 
			animation1()
			animation2()
		case 3: 
			animation1()
			animation2()
			animation3()
		case 4: 
			animation1()				// Das ist der Vollständige Bug 
			animation2()
			animation3()
			for i:=uint16(0);i<b.dying;i++{		// dying wird größer
				animation4()
			}
		case 5:
			for i:=uint16(0);i<b.dying;i++{		// Code Reperatur, quasi Implosion rückwärts. dying wird kleiner
				animation4()
			}
		
		default: return
	}
	
}


// Vor.: keine
// Eff.: Im Welt-Array ist an der Stelle des Bugs jetzt eine 2
func (b *bug) eatCode() {
		welt[(b.y)/zH-3][(b.x)/zB+3] = 2
}

// Vor.: keine
// Eff.: Bug hat sich zufällig unter Einhaltung der Randbedingungen bewegt 
func (b *bug) startMoving() {
	for b.stirbt == false {						// solange Bug nicht stirbt
		time.Sleep(time.Duration(5e8+rand.Intn(5000/b.nervosität)*1e6)) // Warte zufällige Zeit bevor sich Position ändert
		if b.stirbt == true {return}		// verhindert Bewegung nachdem Bug am sterben ist
		bugArraySchloss.Lock()				// Ich möchte auf den BugArray zugreifen 
		var new_x, new_y uint16				// potentielle neue Koordinaten
		new_x = uint16(int(b.x)+((-b.speed/2 +rand.Intn(1+b.speed)))*int(zB))		// zufällige Bewegung, rand.Intn(3) gibt Zahlen 0,1,2, daher +1
		new_y = uint16(int(b.y)+((-b.speed/2 +rand.Intn(1+b.speed)))*int(zH))		// Multiplikation mit zB oder zH damit sich der Bug auf dem richtigen Raster bewegt
		
		// Jetzt muss geprüft werden, ob die neue Position belegt werden kann
		// Es können nicht immer alle Randbedingung eingehalten werden
		// Insgesamt bewegen sich die Bugs aber voneinander weg wenn sie zu nahe sind
		
		// Kollision mit anderen Bugs
		for _,bu:= range bugArray {				// Prüfe Kollision mit allen Bugs
			if bu==nil {continue}				// Prüfe ob hier wirklich ein Bug im Array ist
			if b!=bu {							// Wenn ich es nicht selber bin
				new_x = versetzeBug(new_x,new_y,bu)		// versetze mich in x-Richtung
			}	
		}
		
		
		// Randbedingungen
		if 	new_x < 0{				// Am linken Rand
			b.x =weltB*zB-7*zB		// Verschiebe den Bug auf die andere Seite, aber mit Abstand zum Rand, 7 ist die Breite eines Bugs
		}else if  new_x > weltB*zB-zB*7{	// Am rechten Rand	
			b.x=0					// Verschiebe den Bug an den linken Rand
		}else {b.x = new_x}			// Mache nichts
		
		if  new_y < y_offset*zH {	// Am unteren Rand
			b.y=weltH*zH+y_offset*zH-7*zH	// Verschiebe an den oberen Rand, aber nur innerhalb der Welt
		}else if new_y > weltH*zH+y_offset*zH- 7*zH{	// Am oberen Rand
			b.y=y_offset*zH			// Verschiebe an den unteren Rand
		}else {b.y = new_y} 		// Mache nichts
		
 		bugArraySchloss.Unlock()	// Gib den BugArray wieder frei
	}
	
}

// Hilfsfunktion um einen Bug bei Kollision zu versetzen
// Vor.: keine
// Eff.: Der Bug ist in x-Richtung von einem andern Bug wegverschoben, wenn es eine Kollsision gab. Sonst ist nichts passiert
func versetzeBug(x,y uint16, bu *bug) uint16{			// x,y sind die neuen Positionen des Bugs, bu ist ein andere Bug im BugArray
	// Wenn sich zwei Bugs überdecken: Es wird geschaut, ob sich zwei Bugs irgendwo überdecken, auch teilweise					
	if math.Abs(float64(x)-float64(bu.x))<float64(7*zB) && math.Abs(float64(y)-float64(bu.y))<float64(7*zH) {
		return uint16(int(x)-int(bu.x)+int(x))		// versetze den bug nach links oder rechts von dem Bug mit dem die Kollision erfolgte
		
	}
	return x										// Falls es keine Kollision gibt, behalte die x-Position
}

// Vor.: keine
// Eff.: Animationsschritt des Bugs ist geändert
func (b *bug) bugAnimation() {
	b.alive=true				// belebt Bug
	
	for b.alive{				// Ändere Animation nur wenn am leben

		if b.stirbt {			// Wenn Bug getroffen wird stirbt er
			b.a = 4				// Dieser Teil muss wiederholt aufgerufen werden, sonst kann sich ein Bug immer nur am Ende
			break				// einer kompletten Animation zersetzen, das ist zu langsam
		}
		// Animation startet

		b.a=1					// Körper zeichnen
		time.Sleep(4e8)			// warte kurz
		if b.stirbt {			// Wieder Test ob noch am Leben
			b.a = 4
			break
		}
		b.a=2					// Arme und Beine zeichnen
		time.Sleep(4e8)
		if b.stirbt {
			b.a = 4
			break
		}
		b.a=3					// Komplett zeichnen
		time.Sleep(4e8)
		

	}
	
	// Wenn der Bug getroffen wurde geht es hier weiter
	for {
		if b.dying>50 {break}					// Das Attribut dying wird verändert, so dass bei der Implosion verschieden viele Pixel schwarz gezeichnet werden (immer mehr)
		//fmt.Println("Bug is dying:(")
		b.dying++
		time.Sleep(1e7)							// Geschwindigkeit der Implosion
	}
	b.a = 5			// Code Reperatur			
	for {
		if b.dying==0 {break}					// Wenn dying == 0 höre auf mit zeichnen
		b.dying--
		time.Sleep(1e7)
	}
	b.alive = false								// Bug ist nicht mehr am Leben --> Wir dann mit einer anderen Funktion aus dem BugArray gelöscht
	
}


// Hilfsfunktion
// Vor.: Keine
// Erg.: gerade Zahl 
func rundeAufGeradeZahlen(z int) int{
	if z%2 == 0 {return z}
	return z+1
}

// Vor.: keine
// Eff.: Erzeugt einen neuen Bug an der Stelle eines anderen Bugs
func babyBugs(b *bug) {
	
	if howManyBugs() < len(bugArray) {			// Prüft wie viele Bugs es gibt, wenn es schon zu viele gibt, dann können keine mehr geboren werden
		bugArraySchloss.Lock()					// Ich möchte den BugArray beschreiben
		for index,bu:= range bugArray {			
			// Wenn noch Platz ist erzeuge neuen Bug
			if bu==nil {						// Wenn ein leerer Platz im BugArray gefunden wurde
				bugArray[index]=NewBug(b.x/zB,b.y/zH-y_offset) 		// erzeuge neuen Bug an der gleichen Stelle wie der "Mutter/Vater"-Bug
				bugArray[index].nervosität=5+rand.Intn(5)			// Alle Baby-Bugs bekommen eine zufällige Nervosität
				bugArray[index].speed = rundeAufGeradeZahlen(rand.Intn(5)) // Bei ungeraden Zahlen bewegen sich die Bugs nicht genau zufällig nach links/rechts, oben/unten, daher runden
				bugArray[index].b = uint8(25*bugArray[index].nervosität)	// Füge den neuen Bug zum BugArray hinzu
				go bugArray[index].bugAnimation()							// Starte Animation des neuen Bugs
				go bugArray[index].startMoving()							// Starte Bewegung des neuen Bugs
				break
			}
		}
		bugArraySchloss.Unlock()										// BugArray ist wieder frei
	}
	
	
}


/*
 * 
 *		WEITERE FUNKTIONEN DIE ALLE BUGS BETREFFEN
 * 
 */
// Vor.: keine
// Erg.: Anzahl der lebenden Bugs ist geliefert
// Eff.: Bugs haben an ihrem Körpermittelpunkt den Code gefressen (schwarz gemacht)
func howManyBugs() int {
	var anzahl int
	bugArraySchloss.Lock()		// Ich will auf den BugArray zugreifen
	for _,b:= range bugArray {	// Für alle Bugs im BugArray
		if b!=nil {				// nur wenn es einen Bug gibt
			anzahl++			// Erhöhe Anzahl
			b.eatCode()			// Färbe Feld in der der Mitte des Bugs schwarz
			}
	}
	bugArraySchloss.Unlock()	// Ich gebe den BugArray wieder frei
	//time.Sleep(1e6)
	return anzahl				// Gib die Anzahl der lebenden Bugs zurück
}
// Vor.:keine
// Eff.: Alle toten Bugs die noch im BugArray sind werden gelöscht
func cleanBugArray() {
	for {			// Solange bis das Spiel beendet wird
		bugArraySchloss.Lock()			// Ich möchte auf den BugArray zugreifen
		for index,b:= range bugArray {
			// entferne Bug wenn er tot ist
			if b!=nil && b.alive == false { // wenn der existiert aber tot ist
				bugArray[index] = nil		// entferne Bug
			}
		}
		bugArraySchloss.Unlock()
		time.Sleep(1e8)
	}
}




// Vor.: keine
// Eff.: n Bugs sind erstellt wurden, BugArray wurde von Anfang bis n überschrieben 
func createNBugs(n uint16,speed,nervosität int) {
	bugArraySchloss.Lock()
	
	//  Es werden n Bugs erzeugt, speed und nervosität werden zufällig bis 
	//	zur maximalen speed, nervosität gewählt

	for i:=uint16(0);i<n;i++ {
		//fmt.Println(i)
		b := NewBug(uint16(rand.Intn(130)),uint16(rand.Intn(41)))
		b.speed = rand.Intn(speed + 1)
		b.g = uint8(25*b.speed)
		b.nervosität = rand.Intn(nervosität)+1
		b.b = uint8(25*b.nervosität)
		go b.bugAnimation()
		go b.startMoving()
		bugArray[i] = b
		
	}
	bugArraySchloss.Unlock()
}


// Zum Debugging, anzeigen wie viele Bugs es gibt
func ShowBugs(){
	for{
		time.Sleep(1e9)
		fmt.Println(howManyBugs())
	}
}
