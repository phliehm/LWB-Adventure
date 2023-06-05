package bugPackage

import (
		"gfx"
		"fmt"
		"time"
		"math/rand"
		"../../Klassen/textboxen"
		)

// zeichnet die Welt
func ZeichneWelt() {
	// Die Wait-Group stellt sicher, dass vor dem Beenden des Levels die Animation fertig ist 
	defer wg.Done()
	//  Nur solange es noch Bugs gibt
	for howManyBugs()>0 {
		gfx.UpdateAus()
		gfx.Cls()
		// setze globale Stiftfarbe
		sg = 255	
		zeichneArray()
		gfx.LadeBild(0,0,"../../Bilder/BugAttack/Amoebius_klein.bmp")
		punkteTB.SchreibeText(manual+"Punkte: "+fmt.Sprint((zählePunkte())))
		punkteTB.Zeichne()
		bugArraySchloss.Lock()
		for index,_ := range bugArray {
			if bugArray[index]==nil {continue}
			bugArray[index].zeichneBug()
		}
		bugArraySchloss.Unlock()
		cursorZeichnen()
		gfx.UpdateAn()
		time.Sleep(1e7)
	}
	time.Sleep(1e9)
}

// Wie ZeichneWelt nur mit einigen Änderungen, nur als Animation für den Startbildschirm genutzt
func ZeichneWeltIntro() {
	defer wg.Done()
	BugAttackTB := textboxen.New(200,150,700,500)
	BugAttackTB.SetzeZentriert()
	BugAttackTB.SchreibeText(
		"BUG ATTACK")
	BugAttackTB.SetzeFont("../../Schriftarten/ComputerTypewriter.ttf")
	BugAttackTB.SetzeFarbe(0,255,0)
	BugAttackTB.SetzeSchriftgröße(200)
	
	for howManyBugs()>0 {				// Sobald es keine Bugs mehr gibt soll die Funktion beendet werden
		gfx.UpdateAus()
		gfx.Stiftfarbe(0,0,0)
		gfx.Cls()

		zeichneArrayIntro()
		
		bugArraySchloss.Lock()
		for index,_ := range bugArray {
			//fmt.Println("zeichne bugArray")
			if bugArray[index]==nil {continue}
			bugArray[index].zeichneBug()
		}
		bugArraySchloss.Unlock()
		
		BugAttackTB.Zeichne()
		
		gfx.UpdateAn()
		time.Sleep(1e7)
	}
}

// Hilfsfunktion um die Zahlen (Code) zu zeichnen
func zeichneArray() {
	var s,z uint16
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	//gfx.Stiftfarbe(0,255,0)
	sr,sg,sb = 0,255,0
	for z=0;z<weltH;z++ {
		for s=0;s<weltB;s++ {
			male_Zahl(s*zB,y_offset*zH+z*zH,welt[z][s])
		}
	}
}

// zeichnet den Code, diesmal aber über das gesamte Fenster
func zeichneArrayIntro() {
	var s,z uint16
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	//gfx.Stiftfarbe(0,255,0)
	sr,sg,sb = 0,100,0
	for z=0;z<weltHIntro;z++ {
		for s=0;s<weltB;s++ {
			male_Zahl(s*zB,z*zH,weltIntro[z][s])
		}
	}
}

// Füllt den Array der Welt mit zufällig 1  oder 0
func beschreibeArray(){
	var s,z uint16
	for z=0;z<weltH;z++ {
		for s=0;s<weltB;s++ {
			welt[z][s] = uint8(rand.Intn(2))
		}
	}
}

// Füllt den Array der Welt mit zufällig 1  oder 0
func beschreibeArrayIntro(){
	var s,z uint16
	for z=0;z<weltHIntro;z++ {
		for s=0;s<weltB;s++ {
			weltIntro[z][s] = uint8(rand.Intn(2))
		}
	}
}

// Füllt den Array mit Schwarz
func beschreibeArraySchwarz(){
	var s,z uint16
	for z=0;z<weltH;z++ {
		for s=0;s<weltB;s++ {
			welt[z][s] = 2
		}
	}
}

// Zeichne den Cursor
func cursorZeichnen() {
		gfx.Stiftfarbe(0,255,0)
		gfx.Vollrechteck(cursor_x+2,cursor_y+2,zB-3,zH-3)
		gfx.Stiftfarbe(0,0,0)
}

// Update Cursor-Position
func CursorPos() {
	var step uint16 = 1
	for {
		//gfx.Stiftfarbe(0,255,0)
		taste, gedrueckt, tiefe = gfx.TastaturLesen1()
		if tiefe==1 {
			step=10
		}else {step=1}
		if gedrueckt == 1 {
			//fmt.Println(taste)
			switch taste {
				case 273:		// hoch
							cursor_y -= step*zH
							// Wenn der Cursor über dem Weltrand oder über dem FensterRand ist, setze auf andere Seite
							if cursor_y<y_offset*zH || cursor_y> 1000*zH{cursor_y = y_offset*zH+weltH*zH-zH}	
							
				case 274:  // runter
							cursor_y += step*zH
							if cursor_y>y_offset*zH+weltH*zH-zH {cursor_y = y_offset*zH}
				case 275:	// rechts
							cursor_x += step*zB
							if cursor_x>weltB*zB-zB {cursor_x = 0}
				case 276:	// links
							cursor_x -= step*zB
							if cursor_x>weltB*zB-zB {cursor_x = weltB*zB-zB}
				case 32 : 	// Schießen
							welt[(cursor_y-y_offset*zH)/zH][cursor_x/zB] = 0
							
							bugGetroffen()
				case 120: 	cursor_x, cursor_y = getNextAliveBug() // autoAim
				case 107: 	killAllBugs()
				case 'q':	// beende Game
					gfx.FensterAus()
					return
			
				default:
					continue				
			}
		//fmt.Println(taste,tiefe,cursor_x,cursor_y)	
		}
		time.Sleep(1e6)
	}
}

// Wenn man autoAim verwendet gibt diese Funktion den nächsten Bug im Array
func getNextAliveBug() (uint16, uint16) {
	// Wenn autoAim noch nicht an ist, mache nichts
	if !autoAim {return 0,0+y_offset*zH}
	bugArraySchloss.Lock()
	for _,b := range bugArray {
		if b!=nil {
			bugArraySchloss.Unlock()
			return b.x+3*zB,b.y+3*zH
		}
	}
	bugArraySchloss.Unlock()
	return 0,0
}

// Zählt die Punkte im Array
func zählePunkte() uint32 {
	var abzug uint32
	var z,s uint16 
	for z=0;z<weltH;z++ {
		for s=0;s<weltB;s++ {
			if welt[z][s] ==2 {abzug+=100}
		}
	}
	return maxPunkteProLevel-abzug
}

// Prüft ob ein Bug getroffen wurde und zerstört diesen Bug oder lässt neuen wachsen, oder macht nichts
func bugGetroffen() {
	bugArraySchloss.Lock()
	for _,b:= range bugArray {
		if b==nil{continue}	
		// Wenn der Cursor in der Mitte des Bugs ist
		if (cursor_x-3*zB == b.x && cursor_y-3*zH==b.y) {
			gfx.SpieleSound("../../Sounds/Retro Sounds/Explosions/Long/sfx_exp_long1.wav")
			b.stirbt=true
			bugArraySchloss.Unlock()
			return
			// Wenn der Cursor auf dem Rest des Körpers des Bugs ist --> erzeuge neue Bugs
		}else if (cursor_x > b.x+zB && cursor_x<b.x+6*zB) && (cursor_y> b.y+zH && cursor_y<b.y+6*zH) {
			//fmt.Println("Oh nein!! Der Bug ist provoziert")
			gfx.SpieleSound("../../Sounds/Retro Sounds/General Sounds/Negative Sounds/sfx_sounds_damage1.wav")
			// Neuen Bug generieren
			bugArraySchloss.Unlock()
			babyBugs(b)
			return
		}
				
	}
	bugArraySchloss.Unlock()
	
	// keinen Bug getroffen, mache Feld schwarz
	welt[cursor_y/zH-y_offset][cursor_x/zB] = 2
}
