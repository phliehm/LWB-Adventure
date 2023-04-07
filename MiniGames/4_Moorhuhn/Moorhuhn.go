package main
// Autor: B. Schneider
// Datum: 20.03.2023
// Zweck: SWP - Minigame Moorhuhn
//--------------------------------------------------------------------

import ( 	. "gfx"
			"time"
			"fmt"
			"sync"
			"../../Klassen/objekte"
			"../../Klassen/raeume"
			"../../Klassen/texte"
			"math/rand"
			)

const breite = 800 		// von Gott vorgegeben
const hoehe  = 600  	// von Gott vorgegeben
	
func main () {
	var mutex sync.Mutex					// erstellt Mutex
	var punkte int16 = 0					// Spiel-Punktzahl
	var stop bool = false					// für OK-Objekt
	var hubi bool = false					// für Hubi-Abwehr
	var pause bool = false
	var akt	bool = true						// Prüft, ob Grafik aktualisiert werden muss
	obj := make([]objekte.Objekt,0)			// Array für die Objekte der Welt
	
	
	random := rand.New( rand.NewSource( time.Now().UnixNano() ) )	// Initialisiere Random-Objekt mit der Systemzeit
	
	
	pauseObjekt := objekte.New(breite,hoehe,hoehe/4	,1)			// Erstellt das Objekt PAUSE 
	maus 		:= objekte.New(0,0,			30		,0)			// Erstellt das Objekt MAUSZEIGER mit Größe (30)
	okayObjekt 	:= objekte.New(breite,hoehe,hoehe	,20)		// OK-Objekt
	
	Fenster (breite,hoehe)									// Öffnet GFX-Fenster
	Fenstertitel("StEPS-Wars")								// Gibt Fenster-Titel 
	
	SetzeFont ("../../Schriftarten/Freshman.ttf", hoehe/20 ) 	// Setzt Schriftart
	
	
	// Das Hauptprogramm startet die View-Komponente als nebenläufigen Prozess!
	go view_komponente(&obj, maus, pauseObjekt, okayObjekt, &pause, &stop, &akt, &punkte, &mutex)
	
	// Objekte werden nach und nach in der Welt platziert
	go erstelleObjekte(&obj, maus, &pause, &stop, &hubi, &akt, random, &mutex)
	
	// Nebenläufig wird die Kontroll-Komponente für die Maus gestartet.
	go maussteuerung(&obj, maus, okayObjekt, &pause, &stop, &hubi, &akt, &punkte)
	
	
	// Die Kontroll-Komponente 2 ist die 'Mainloop' im Hauptprogramm	
	// Wir fragen hier nur die Tastatur ab.
	
A:	for {
		taste,gedrueckt,_:= TastaturLesen1()
		if gedrueckt == 1 {
			switch taste {
				case 'q': 													// mit 'q' wird das Programm beendet!
				break A
				case 'p': 
				pause = !pause												// Pause-Modus !!

				case 'h':													// für Lvl. 5 zum Vertreiben von Hubi
				if hubi {
					hubi = false
					obj = make([]objekte.Objekt,0)
					akt = true
					SpieleSound("../../Sounds/Sparkle.wav")
					punkte += 10
				} else {
					SpieleSound("../../Sounds/Baeb.wav")
					punkte -= 5
				}
				case 'b':
				SpieleSound("../../Sounds/Roar.wav")
			}
		}
	}
	// return punkte
}

func erstelleObjekte(obj *[]objekte.Objekt, maus objekte.Objekt, pause,stop,hubi,akt *bool, rand *rand.Rand, mutex *sync.Mutex) {		// füllt Objekte ins Array
	level1 := objekte.New(breite,hoehe,hoehe	,7)
	level2 := objekte.New(breite,hoehe,hoehe	,8)
	level3 := objekte.New(breite,hoehe,hoehe	,9)  
	level4 := objekte.New(breite,hoehe,hoehe	,10)
	level5 := objekte.New(breite,hoehe,hoehe	,15)
	count3 := objekte.New(breite,hoehe,hoehe	,12)
	count2 := objekte.New(breite,hoehe,hoehe	,13)
	count1 := objekte.New(breite,hoehe,hoehe	,14)

	time.Sleep( time.Duration(1e9) )
	
	// -------------------------------------------------------------------------------------------------------------------------------------------------------------
	
	Zwischentext(&texte.MoorEinl, mutex, stop)		// Einleitungs-Text
	
	Levelanzeige(level1, mutex)						// ----------------- LEVEL 1 -------------------- Große Zielscheiben
	
	Countdown(count3,count2,count1, mutex, akt)		// lässt den Bildschirm-Countdown ablaufen
	
	
	for i:=0;i<15;i++ {
		time.Sleep( time.Duration(	rand.Uint32()/2 ) )
		*obj = append(*obj, objekte.New( uint16(rand.Intn(600))+50, uint16(rand.Intn(400))+50, uint16(rand.Intn(150)+150), 11) )
		*akt = true
	}
	
	time.Sleep( time.Duration(4e9) )
	
	*obj = make([]objekte.Objekt,0)					// leere den Objekte-Slice (Performance!)
	
	// ----------------------------------------------------------------------------------------------
	
	Zwischentext(&texte.MoorLvl2, mutex, stop)		// Level 2-Text
		
	Levelanzeige(level2, mutex)						// ----------------- LEVEL 2 -------------------- Kleine Zielscheiben
	
	Countdown(count3,count2,count1, mutex, akt)		// lässt den Bildschirm-Countdown ablaufen
	
	for i:=0;i<15;i++ {
		time.Sleep( time.Duration( rand.Uint32()/2 ) )
		*obj = append(*obj, objekte.New( uint16(rand.Intn(650))+50, uint16(rand.Intn(450))+50, uint16(rand.Intn(100)+30), 11) )
		*akt = true
	}
	time.Sleep( time.Duration(4e9) )
	
	*obj = make([]objekte.Objekt,0)			// leere den Objekte-Slice (Performance!)
	
	// ----------------------------------------------------------------------------------------------
	
	Zwischentext(&texte.MoorLvl3, mutex, stop)		// Level 3-Text
		
	Levelanzeige(level3, mutex)						// ----------------- LEVEL 3 -------------------- Kaffee und Pizza
	
	Countdown(count3,count2,count1, mutex, akt)		// lässt den Bildschirm-Countdown ablaufen
	
	
	for i:=0;i<15;i++ {
		*obj = append(*obj, objekte.New( uint16(rand.Intn(650))+50,uint16(rand.Intn(450))+50, uint16(rand.Intn(200)+50), uint8(rand.Intn(2)*2+3)) )
		*akt = true
		time.Sleep( time.Duration(	rand.Uint32()/2 ) )
	}
	
	time.Sleep( time.Duration(4e9) )
	
	*obj = make([]objekte.Objekt,0)			// leere den Objekte-Slice (Performance!)
	
	// ----------------------------------------------------------------------------------------------
	
	Zwischentext(&texte.MoorLvl4, mutex, stop)		// Level 4-Text
	
	Levelanzeige(level4, mutex)						// ----------------- LEVEL 4 --------------------
	
	Countdown(count3,count2,count1, mutex, akt)		// lässt den Bildschirm-Countdown ablaufen
		
	*obj = append(*obj, objekte.New(50,150,50,uint8(rand.Intn(2)+18)),objekte.New(50,200,50,uint8(rand.Intn(2)+18)),objekte.New(50,250,50,uint8(rand.Intn(2)+18)),		// L
						objekte.New(50,300,50,uint8(rand.Intn(2)+18)),objekte.New(50,350,50,uint8(rand.Intn(2)+18)),objekte.New(50,400,50,uint8(rand.Intn(2)+18)),
						objekte.New(50,450,50,uint8(rand.Intn(2)+18)),objekte.New(100,450,50,uint8(rand.Intn(2)+18)),objekte.New(150,450,50,uint8(rand.Intn(2)+18)),
						  
						objekte.New(215,150,50,uint8(rand.Intn(2)+18)),objekte.New(220,200,50,uint8(rand.Intn(2)+18)),objekte.New(230,250,50,uint8(rand.Intn(2)+18)),	// W
						objekte.New(245,300,50,uint8(rand.Intn(2)+18)),objekte.New(260,350,50,uint8(rand.Intn(2)+18)),objekte.New(275,400,50,uint8(rand.Intn(2)+18)),
						objekte.New(295,450,50,uint8(rand.Intn(2)+18)),objekte.New(325,400,50,uint8(rand.Intn(2)+18)),objekte.New(350,350,50,uint8(rand.Intn(2)+18)),
						objekte.New(375,400,50,uint8(rand.Intn(2)+18)),objekte.New(405,450,50,uint8(rand.Intn(2)+18)),objekte.New(425,400,50,uint8(rand.Intn(2)+18)),
						objekte.New(440,350,50,uint8(rand.Intn(2)+18)),objekte.New(455,300,50,uint8(rand.Intn(2)+18)),objekte.New(470,250,50,uint8(rand.Intn(2)+18)),
						objekte.New(480,200,50,uint8(rand.Intn(2)+18)),objekte.New(485,150,50,uint8(rand.Intn(2)+18)),
						objekte.New(600,150,50,uint8(rand.Intn(2)+18)),objekte.New(600,200,50,uint8(rand.Intn(2)+18)),objekte.New(600,250,50,uint8(rand.Intn(2)+18)),	// B
						objekte.New(600,300,50,uint8(rand.Intn(2)+18)),objekte.New(600,350,50,uint8(rand.Intn(2)+18)),objekte.New(600,400,50,uint8(rand.Intn(2)+18)),
						objekte.New(600,450,50,uint8(rand.Intn(2)+18)),
						objekte.New(660,170,50,uint8(rand.Intn(2)+18)),objekte.New(700,215,50,uint8(rand.Intn(2)+18)),objekte.New(660,260,50,uint8(rand.Intn(2)+18)),
						objekte.New(680,315,50,uint8(rand.Intn(2)+18)),objekte.New(725,360,50,uint8(rand.Intn(2)+18)),objekte.New(710,415,50,uint8(rand.Intn(2)+18)),
						objekte.New(660,440,50,uint8(rand.Intn(2)+18)) )
	
	*akt = true
	time.Sleep( time.Duration(1.4e10) )
	
	*obj = make([]objekte.Objekt,0)					// leere den Objekte-Slice (Performance!)

	// ----------------------------------------------------------------------------------------------
	
	Zwischentext(&texte.MoorLvl51, mutex, stop)		// Level 5-Text - 1
	
	maus.SetzeTyp(16)
	Zwischentext(&texte.MoorLvl52, mutex, stop)		// Level 5-Text - 2
	maus.SetzeTyp(0)
	Levelanzeige(level5, mutex)						// ----------------- LEVEL 5 -------------------- Kaffee und Pizza
	
	Countdown(count3,count2,count1, mutex, akt)		// lässt den Bildschirm-Countdown ablaufen
	
	jetzt := time.Now().UnixNano()
	zeit := int64(0)
	
	for zeit < 4e10 {	
		
		if i:= rand.Intn(20); i < 8 {
			*obj = append(*obj, objekte.New( uint16(rand.Intn(600))+50,uint16(rand.Intn(400))+50, uint16(rand.Intn(150)+100), uint8(rand.Intn(2)+18)) )
			*akt = true
		} else if i == 11 {
			*obj = append(*obj, objekte.New( uint16(rand.Intn(710)),uint16(rand.Intn(510)), 0, 16) )
			*hubi = true
			*akt = true
		}
		
		time.Sleep( time.Duration(2e8) )
		zeit = time.Now().UnixNano() - jetzt
	}
	
	time.Sleep( time.Duration(2e9) )
	*obj = make([]objekte.Objekt,0)					// leere den Objekte-Slice (Performance!)
	
	// ----------------------------------------------------------------------------------------------

	Zwischentext(&texte.MoorOut, mutex, stop)		// Ende-Text	
	
	// ----------------------------------------------------------------------------------------------
}

func Zwischentext(textArr *[]string, mutex *sync.Mutex, stop *bool) {
	mutex.Lock()
	raeume.Moorhuhn(breite)
	Transparenz(120)															
	Vollrechteck(breite/10,hoehe/10,breite*4/5,hoehe*4/5)
	Transparenz(0)
	SetzeFont ("../../Schriftarten/Freshman.ttf", hoehe/16 )
	Stiftfarbe(124,212,255)
	for ind,str := range *textArr {
		SchreibeFont (breite*4/25, uint16(hoehe*4/25+ind*hoehe/15) ,str )
	}
	Archivieren()
	mutex.Unlock()
	
	*stop = true
	for *stop { time.Sleep( time.Duration(1e9) ) }
}

func Levelanzeige(level objekte.Objekt, mutex *sync.Mutex) {
	mutex.Lock()
	raeume.Moorhuhn(breite)
	level.Zeichnen()
	Archivieren()
	mutex.Unlock()
	
	time.Sleep( time.Duration(1.5e9) )
}

func Countdown(count3,count2,count1 objekte.Objekt, mutex *sync.Mutex, akt *bool) {
	mutex.Lock()
	raeume.Moorhuhn(breite)
	count3.Zeichnen()
	Archivieren()
	mutex.Unlock()
	
	time.Sleep( time.Duration(1e9) )
	
	mutex.Lock()
	raeume.Moorhuhn(breite)
	count2.Zeichnen()
	Archivieren()
	mutex.Unlock()
	
	time.Sleep( time.Duration(1e9) )
	
	mutex.Lock()
	raeume.Moorhuhn(breite)
	count1.Zeichnen()
	Archivieren()
	mutex.Unlock()
	
	time.Sleep( time.Duration(1e9) )
	*akt = true
}

// Es folgt die VIEW-Komponente
func view_komponente (obj *[]objekte.Objekt, maus,pauseObjekt,okayObjekt objekte.Objekt, pause, stop ,akt *bool, punkte *int16, mutex *sync.Mutex) {   	
	var t1 int64 = time.Now().UnixNano() 		//Startzeit
	var anz,anzahl int                  		// zur Bestimmung der Frames pro Sekunde
	var verzögerung = 90
	
	for { //Endlos ...
		mutex.Lock()
		UpdateAus () 										// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
		
		Stiftfarbe(255,255,255)
		Cls()												// Cleart vollständigen Screen
		
		if *akt { 
			ObjAktualisieren(obj)
			*akt = false
		} else {
			Restaurieren(0,0,breite,hoehe)					// Restauriert das alte Hintergrundbild
		}
		
		if *stop {
			okayObjekt.Zeichnen()
		}
		
		maus.Zeichnen()										// Zeichnet Maus
		
		SetzeFont ("../../Schriftarten/Freshman.ttf", hoehe/20 )
		Stiftfarbe(76,0,153)  
		SchreibeFont (breite*14/20,0,"Punkte : "+fmt.Sprint (*punkte))	// Schreibe rechts oben Punkte
		Stiftfarbe(100,10,155)
		Schreibe (0,0,"FPS:"+fmt.Sprint (anzahl))					// Schreibe links oben FPS
		if *pause { pauseObjekt.Zeichnen() }
			
		
		if time.Now().UnixNano() - t1 < 1000000000 { //noch in der Sekunde ...
			anz++
		} else {
			t1 = time.Now().UnixNano() // neue Sekunde
			anzahl = anz
			anz=0
			if anzahl < 100 { verzögerung--}				//Selbstregulierung der 
			if anzahl > 100 { verzögerung++}				//Frame-Rate :-)
		}
		
		UpdateAn () 										// Nun wird der gezeichnete Frame sichtbar gemacht!
		mutex.Unlock()
		
		time.Sleep(time.Duration(verzögerung * 1e5)) 		// Immer ca. 100 FPS !!
		
	}
}

func ObjAktualisieren(obj *[]objekte.Objekt) {
	raeume.Moorhuhn(breite)									// Hintergrund des Moorhuhn-Raumes wird gezeichnet
	
	for _,ob := range *obj { 								// Zeichnet alleweiteren Objekte ein
		ob.Zeichnen()
	}
	Archivieren()											// Speichert das Hintergrund-Bild
}

// Es folgt die CONTROL-Komponente 1 --- Kein Bestandteil der Welt, also unabhängig -----
func maussteuerung (obj *[]objekte.Objekt, maus,okayObjekt objekte.Objekt, pause,stop,hubi,akt *bool, punkte *int16) {
	/*var taste uint8
	var status int8 */
	for {
		taste, status, mausX, mausY := MausLesen1()
		
		maus.SetzeKoordinaten(mausX,mausY)					// Aktualisiert Maus-Koordinaten
		
		if *stop {
			if taste==1 && status==1 { 						//LINKE Maustaste gerade gedrückt
				if ja,_ := okayObjekt.Getroffen(mausX,mausY,1); ja {
					*stop = false
				}
			}
		} else if *pause {	
		} else if *hubi {									// falls ein Hubi (Lvl 5) aktiv ist
			if status==1 {
				SpieleSound("../../Sounds/Baeb.wav")
				*punkte -= 5
			}
		} else {
			if taste==1 && status==1 { 						//LINKE Maustaste gerade gedrückt
				for _,ob := range *obj { 							// Zeichnet alleweiteren Objekte ein
					if get,lang := ob.Getroffen(mausX,mausY,1); get {
						if lang == 0 {
							*punkte -= 5
							 fmt.Println("Knapp daneben - 5 MINUS-Punkte")
						} else if lang < 3.5e8 {
							 *punkte += 20
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 20 Punkte")
						 } else if lang < 4.2e8 {
							 *punkte += 15
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 15 Punkte")
						 } else if lang < 5.2e8 {
							 *punkte += 10
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 10 Punkte")
						 } else if lang < 7e8 {
							 *punkte += 5
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 5 Punkte")
						 } else if lang < 1e9 {
							 *punkte += 2
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 2 Punkte")
						 } else {
							 *punkte += 1
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 1 Punkt")
						}	
						*akt = true
					}
				}
			}
			if taste == 3 && status == 1 { 			//RECHTE Maustaste gerade gedrückt
				for _,ob := range *obj { 							// Zeichnet alleweiteren Objekte ein
					if get,lang :=  ob.Getroffen(mausX,mausY,3); get {
						if lang == 0 {
							*punkte -= 5
							 fmt.Println("Knapp daneben - 5 MINUS-Punkte")
						} else if lang < 3.5e8 {
							 *punkte += 20
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 20 Punkte")
						} else if lang < 4.2e8 {
							 *punkte += 15
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 15 Punkte")
						} else if lang < 5.2e8 {
							 *punkte += 10
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 10 Punkte")
						} else if lang < 7e8 {
							 *punkte += 5
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 5 Punkte")
						} else if lang < 1e9 {
							 *punkte += 2
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 2 Punkte")
						} else {
							 *punkte += 1
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 1 Punkt")
						}		
						*akt = true
					}
				}
			}
		}
	}
}

