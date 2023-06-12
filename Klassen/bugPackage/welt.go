/* Philipp Liehm
 * Juni 2023
 * LWB-Adventure: BugAttack
 */

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
		gfx.LadeBild(0,0,"Bilder/BugAttack/Amoebius_klein.bmp")
		punkteTB.SchreibeText(manual + "Level: " + fmt.Sprint(level) + "  |  "+ "Fähigkeiten: "  )		// MUSS UMBENANNT WERDEN!!! Nur manual
		punkteTB.Zeichne()
		
		bugArraySchloss.Lock()
		for index,_ := range bugArray {
			if bugArray[index]==nil {continue}
			bugArray[index].zeichneBug()
		}
		bugArraySchloss.Unlock()
		cursorZeichnen()
		zeichneAlleLadebalken()
		zeichneZeit()
		zeichnePunkte()
		gfx.UpdateAn()
		time.Sleep(1e7)
	}
	time.Sleep(1e9)
}

// Wie ZeichneWelt nur mit einigen Änderungen, nur als Animation für den Startbildschirm genutzt
func ZeichneWeltIntro() {
	level = 1		// muss auf 1 gesetzt werden, sonst crashed das Spiel, nur hier für die Intro Funktion nötig
	defer wg.Done()
	BugAttackTB := textboxen.New(200,150,700,500)
	BugAttackTB.SetzeZentriert()
	BugAttackTB.SchreibeText(
		"BUG ATTACK")
	BugAttackTB.SetzeFont("Schriftarten/ComputerTypewriter.ttf")
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
		
		gfx.Stiftfarbe(0,255,0)
		gfx.SetzeFont("Schriftarten/COOPBL.TTF",30)
		gfx.SchreibeFont(100,600,"Drücke eine Taste, dann verschwinden auch die Bugs, ... vorerst!")
		
		gfx.UpdateAn()
		time.Sleep(1e7)
	}
	level=0		// Setze level wieder auf 0 damit es ganz normal los gehen kan
}

// Hilfsfunktion um die Zahlen (Code) / Welt zu zeichnen
func zeichneArray() {
	var s,z uint16
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	
	sr,sg,sb = 0,200,0
	for z=0;z<weltH;z++ {
		for s=0;s<weltB;s++ {
			male_Zahl(s*zB,y_offset*zH+z*zH,welt[z][s])	// y_offset weil die Zahlen erst weiter unten beginnen
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
func TastaturEingabe(c chan bool) {
	var step uint16 = 1				// Variable für Schrittweite des Cursors wenn SHIFT gedrückt wird
	for {
		select {
			case <-c:
				return
			default:
				//gfx.Stiftfarbe(0,255,0)
				taste, gedrueckt, tiefe = gfx.TastaturLesen1()
				if tiefe==1 {
					step=10			// 10-fache Schrittweite
				}else {step=1}
				if gedrueckt == 1 {
					//fmt.Println(taste)
					switch taste {
						case 273:		// hoch
									cursor_y -= step*zH
									// Wenn der Cursor über dem Weltrand oder über dem FensterRand ist, setze auf andere Seite
									// Die 1000 sorgt hier nur für eine (zu) große Zahl, wie sie bei kleiner 0 auftreten würde	
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
						case 'x': 	benutzeAutoAim() // autoAim
						case 'k': 	killNBugs(5)	// Hier kann mein einstellen wie viele Bugs mit 'k' getötet werden sollen
						case 'q':	// Geht zum Endbildschirm
									beendeSpiel()
									
							//gfx.FensterAus()
							//return
					
						default:
							continue				
					}
				//fmt.Println(taste,tiefe,cursor_x,cursor_y)	
				}
				time.Sleep(1e6)
		}
	}
}

// Wenn man autoAim verwendet gibt diese Funktion den nächsten Bug im Array
func getNextAliveBug() (uint16, uint16) {
	// Wenn autoAim noch nicht an ist, mache nichts
	//if !autoAim {return 0,0+y_offset*zH}
	bugArraySchloss.Lock()
	for _,b := range bugArray {
		if b==nil {continue}
		if !b.stirbt {
			bugArraySchloss.Unlock()
			return b.x+3*zB,b.y+3*zH
		}
	}
	bugArraySchloss.Unlock()
	return 0,0+y_offset*zH
}


// Prüft ob ein Bug getroffen wurde und zerstört diesen Bug oder lässt neuen wachsen, oder macht nichts
func bugGetroffen() {
	bugArraySchloss.Lock()
	for _,b:= range bugArray {
		if b==nil{continue}	
		if b.stirbt {
			continue
		}
		// Wenn der Cursor in der Mitte des Bugs ist
		if (cursor_x-3*zB == b.x && cursor_y-3*zH==b.y)  {
			gfx.SpieleSound("Sounds/Retro Sounds/Explosions/Long/sfx_exp_long1.wav")
			b.stirbt=true
			bugArraySchloss.Unlock()
			return
			// Wenn der Cursor auf dem Rest des Körpers des Bugs ist --> erzeuge neue Bugs
		}else if (cursor_x > b.x+zB && cursor_x<b.x+5*zB) && (cursor_y> b.y+zH && cursor_y<b.y+5*zH) {
			//fmt.Println("Oh nein!! Der Bug ist provoziert")
			gfx.SpieleSound("Sounds/Retro Sounds/General Sounds/Negative Sounds/sfx_sounds_damage1.wav")
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

// Zählt die Punkte im Array
func zählePunkte() {
	defer wg.Done()
	for level <=maxLevel && SpielBeendet == false && lvlLäuft == true{
		levelSchloss.Lock()				// nur ich darf auf das "level" zugreifen
		var abzug uint16
		var z,s uint16 
		for z=0;z<weltH;z++ {			// zeilen der Welt
			for s=0;s<weltB;s++ {		// Spalten der Welt
				if welt[z][s] ==2 {abzug+=10}	// Wenn Feld gegessen (schwarz, 2), ziehe Puntke ab
			}
		}
		abzug+=lvlZeit*50				// Ziehe zusätzlich Punkte für die vergangene Zeit ab
		if abzug > maxPunkteProLevel{
			punkteArray[level-1] = 0		// negative Punkte vermeiden
		}else {punkteArray[level-1] = maxPunkteProLevel-abzug}		// Schreibe die aktuellen Punkte in den PunkteArray
		levelSchloss.Unlock()			// "level" ist wieder freigegeben
		time.Sleep(1e8)
	}
}

// Misst die Zeit in einem Level
func lvlTimer() {
	defer wg.Done()
	for lvlLäuft {
		lvlZeit++
		time.Sleep(1e9)
	}
}

// Beendet ein Level mit 0 Punkten und geht dann zum Endbildschirm
func beendeSpiel() {
	if howManyBugs() > 0{
		killNBugsCD = 10		// Erlaubt alle verbleibenden Bugs zu töten
		killNBugs(100)		// tötet alle Bugs
		fmt.Println("Spiel wird beendet")
		punkteArray[level-1] = 0	// Setzt Punkte für das aktuelle Level auf 0
		//return
	}
	SpielBeendet = true

	return
}


// Alle Bugs mit einem Tastendruck töten
func killNBugs(n uint16) {
	var getötet uint16
	if killNBugsCD!=10 || SpielBeendet == true {return}
	gfx.SpieleSound("Sounds/Retro Sounds/Explosions/Long/sfx_exp_long3.wav")
	bugArraySchloss.Lock()
	for _,b:= range bugArray {
		if getötet>=n {break}
		if b!=nil {
			getötet++
			b.stirbt = true
		}
	}
	killNBugsCD = 0
	bugArraySchloss.Unlock()
}

// 
func benutzeAutoAim() {
	if autoAimCD !=10  {return}
	cursor_x, cursor_y = getNextAliveBug() // autoAim
	autoAimCD = 0
}

// Zeichnet die verbleibende Zeit im Level 
func zeichneZeit() {
	//if lvlLäuft == false {return}
	if maxZeit-lvlZeit<1 && lvlLäuft == true{			// Wenn Zeit abgelaufen
		killNBugsCD = 10		// Setze CD hoch damit man alle Bugs töten kann 
		killNBugs(100)		// Töte alle Bugs
		punkteArray[level-1] = 0
		lvlLäuft = false
		beendeSpiel()
	}
	
	gfx.Stiftfarbe(255,255,255)
	gfx.SetzeFont("Schriftarten/ltypeb.ttf",20)
	gfx.SchreibeFont(850,50,"Zeit: ")
	gfx.SchreibeFont(920,50,fmt.Sprint(maxZeit-lvlZeit))
}

// Zeichnet die verbleibenden Punkte 
func zeichnePunkte() {
	gfx.Stiftfarbe(255,255,255)
	gfx.SetzeFont("Schriftarten/ltypeb.ttf",20)
	gfx.SchreibeFont(1000,50,"Punkte: "+fmt.Sprint(punkteArray[level-1]))
	
}
