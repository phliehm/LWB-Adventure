package moorhuhn
// Autor: B. Schneider
// Datum: 20.03.2023
// Zweck: SWP - Minigame Moorhuhn
//--------------------------------------------------------------------

import ( 	. "gfx"
			"time"
			"fmt"
			"sync"
			"../../Klassen/objekte"
			"../../Klassen/texte"
			"math/rand"
			)

	
func Moorhuhn () (note float32, punktExp uint32) {
	var mutex sync.Mutex					// erstellt Mutex
	var punkte int16 = 0					// Spiel-Punktzahl
	var diff int16 = 0						// Punkte-Veränderung
	var stop bool = false					// für OK-Objekt
	var hubi bool = false					// für Hubi-Abwehr
	var pause bool = false					// falls Pause
	var ende bool = false					// falls Ende
	var akt	bool = true						// Prüft, ob Grafik aktualisiert werden muss
	obj := make([]objekte.Objekt,0)			// Array für die Objekte der Welt
	
	
	random := rand.New( rand.NewSource( time.Now().UnixNano() ) )	// Initialisiere Random-Objekt mit der Systemzeit
	
	maus 		:= objekte.New(0, 0, 0, 0)			// Erstellt das Objekt MAUSZEIGER mit Größe (30)
	pauseObjekt := objekte.New(0, 0, 0, 1)			// Erstellt das Objekt PAUSE 
	okayObjekt 	:= objekte.New(0, 0, 0, 20)			// OK-Objekt
	
	Fenstertitel("StEPS-Wars")								// Gibt Fenster-Titel 
	
	// Das Hauptprogramm startet die View-Komponente als nebenläufigen Prozess!
	go view_komponente(&obj, maus, pauseObjekt, okayObjekt, &pause, &stop, &akt, &ende, &punkte, &diff, &mutex)
	
	// Objekte werden nach und nach in der Welt platziert
	go erstelleObjekte(&obj, maus, &pause, &stop, &hubi, &akt, &ende, random, &punkte, &punktExp, &note, &mutex)
	
	// Nebenläufig wird die Kontroll-Komponente für die Maus gestartet.
	go maussteuerung(&obj, maus, pauseObjekt, okayObjekt, &pause, &stop, &hubi, &akt, &ende, &punkte, &diff)
	
	
	// Die Kontroll-Komponente 2 ist die 'Mainloop' im Hauptprogramm	
	// Wir fragen hier nur die Tastatur ab.
	
A:	for {
		taste,gedrueckt,_:= TastaturLesen1()
		if gedrueckt == 1 {
			switch taste {
				case 'q': 											// mit 'q' wird das Programm beendet!
				if ende { 
					break A
				} else if pause { 
					ende = true
					break A
				}
				case 'p': pause = !pause							// Pause-Modus !!
				case 'h':											// für Lvl. 5 zum Vertreiben von Hubi
				if hubi {
					hubi = false
					obj = make([]objekte.Objekt,0)
					akt = true
					SpieleSound("./Sounds/Sparkle.wav")
					punkte += 10
				} else {
					SpieleSound("./Sounds/Baeb.wav")
					punkte -= 5
				}
				case 'b': SpieleSound("./Sounds/Roar.wav")
			}
		}
	}
	// fmt.Println("Vielen Dank für's Spielen!")
	time.Sleep( time.Duration(2e8) )
	
// -----------------------------------------------------------------	
	
	if punkte > 0 {
		punktExp = uint32(punkte)
	} else {
		punktExp = 0
	}

	switch {
		case punktExp == 0:		note = 6.0
		case punktExp < 50:		note = 4.7
		case punktExp < 100:	note = 4.3
		case punktExp < 150:	note = 4.0
		case punktExp < 200:	note = 3.7
		case punktExp < 250:	note = 3.3
		case punktExp < 300:	note = 3.0
		case punktExp < 350:	note = 2.7
		case punktExp < 400:	note = 2.3
		case punktExp < 450:	note = 2.0
		case punktExp < 500:	note = 1.7
		case punktExp < 550:	note = 1.3
		case punktExp > 550:	note = 1.0
	}
	
	mutex.Lock()
	LadeBild (0,0, "./Bilder/Moorhuhn/Seminarraum.bmp")
	Transparenz(120)															
	Vollrechteck(100,50,1000,600)
	Transparenz(0)
	LadeBild (180,200, "./Bilder/Moorhuhn/StEPS-Logo-2.bmp")
	SetzeFont ("./Schriftarten/Freshman.ttf", 50 )
	
	Stiftfarbe(70,20,20)
	for ind,str := range texte.MoorOut {
		SchreibeFont (208, uint16(69+ind*55) ,str )
	}
	SchreibeFont (880, 344 ,fmt.Sprint(punkte) )
	SchreibeFont (880, 399 ,fmt.Sprint(note) )

	Stiftfarbe(255,124,212)
	for ind,str := range texte.MoorOut {
		SchreibeFont (212, uint16(71+ind*55) ,str )
	}
	SchreibeFont (884, 346 ,fmt.Sprint(punkte) )
	SchreibeFont (884, 401 ,fmt.Sprint(note) )
	
	Archivieren()
	mutex.Unlock()
	
	Restaurieren(0,0,1200,700)
	// -----------------------------------------------------------------
	
	TastaturLesen1()
	return
}

func erstelleObjekte(obj *[]objekte.Objekt, maus objekte.Objekt, pause,stop,hubi,akt,ende *bool, rand *rand.Rand, 
				punkte *int16, punktExp *uint32, note *float32, mutex *sync.Mutex) {		// füllt Objekte ins Array
	
	
	count3 := objekte.New(0,0,0,13)
	count2 := objekte.New(0,0,0,14)
	count1 := objekte.New(0,0,0,15)

	time.Sleep( time.Duration(1e9) )
	
	
	// -------------------------------------------------------------------------------------------------------------------------------------------------------------
	
	Zwischentext(&texte.MoorEinl, mutex, stop)		// Einleitungs-Text
	
	level1 := objekte.New(0,0,0	,7)
	Levelanzeige(level1, mutex)						// ----------------- LEVEL 1 -------------------- Große Zielscheiben
	
	Countdown(count3,count2,count1, mutex, akt)		// lässt den Bildschirm-Countdown ablaufen
	
	
	for i:=0;i<15;i++ {
		time.Sleep( time.Duration(	rand.Uint32()/2 ) )
		*obj = append(*obj, objekte.New( uint16(rand.Intn(1000))+100, uint16(rand.Intn(500))+100, uint16(rand.Intn(225)+225), 12) )
		*akt = true
	}
	
	time.Sleep( time.Duration(4e9) )
	
	*obj = make([]objekte.Objekt,0)					// leere den Objekte-Slice (Performance!)
	
	// ----------------------------------------------------------------------------------------------
	
	Zwischentext(&texte.MoorLvl2, mutex, stop)		// Level 2-Text
		
	level2 := objekte.New(0,0,0	,8)
	Levelanzeige(level2, mutex)						// ----------------- LEVEL 2 -------------------- Kleine Zielscheiben
	
	Countdown(count3,count2,count1, mutex, akt)		// lässt den Bildschirm-Countdown ablaufen
	
	for i:=0;i<15;i++ {
		time.Sleep( time.Duration( rand.Uint32()/2 ) )
		*obj = append(*obj, objekte.New( uint16(rand.Intn(1000))+100, uint16(rand.Intn(500))+100, uint16(rand.Intn(150)+50), 12) )
		*akt = true
	}
	time.Sleep( time.Duration(4e9) )
	
	*obj = make([]objekte.Objekt,0)			// leere den Objekte-Slice (Performance!)
	
	// ----------------------------------------------------------------------------------------------
	
	Zwischentext(&texte.MoorLvl3, mutex, stop)		// Level 3-Text
	
	level3 := objekte.New(0,0,0	,9)  	
	Levelanzeige(level3, mutex)						// ----------------- LEVEL 3 -------------------- Kaffee und Pizza
	
	Countdown(count3,count2,count1, mutex, akt)		// lässt den Bildschirm-Countdown ablaufen
	
	
	for i:=0;i<15;i++ {
		*obj = append(*obj, objekte.New( uint16(rand.Intn(800))+75,uint16(rand.Intn(400))+75, uint16(rand.Intn(300)+50), uint8(rand.Intn(2)*2+3)) )
		*akt = true
		time.Sleep( time.Duration(	rand.Uint32()/2 ) )
	}
	
	time.Sleep( time.Duration(4e9) )
	
	*obj = make([]objekte.Objekt,0)			// leere den Objekte-Slice (Performance!)
	
	// ----------------------------------------------------------------------------------------------
	
	
	Zwischentext(&texte.MoorLvl4, mutex, stop)		// Level 4-Text
	
	level4 := objekte.New(0,0,0	,10)
	Levelanzeige(level4, mutex)						// ----------------- LEVEL 4 --------------------
	
	Countdown(count3,count2,count1, mutex, akt)		// lässt den Bildschirm-Countdown ablaufen
		
	*obj = append(*obj, objekte.New(75,125,75,uint8(rand.Intn(2)+18)),objekte.New(75,200,75,uint8(rand.Intn(2)+18)),objekte.New(75,275,75,uint8(rand.Intn(2)+18)),		// L
						objekte.New(75,350,75,uint8(rand.Intn(2)+18)),objekte.New(75,425,75,uint8(rand.Intn(2)+18)),objekte.New(75,500,75,uint8(rand.Intn(2)+18)),
						objekte.New(75,575,75,uint8(rand.Intn(2)+18)),objekte.New(150,575,75,uint8(rand.Intn(2)+18)),objekte.New(225,575,75,uint8(rand.Intn(2)+18)),
						  
						objekte.New(323,125,75,uint8(rand.Intn(2)+18)),objekte.New(330,200,75,uint8(rand.Intn(2)+18)),objekte.New(345,275,75,uint8(rand.Intn(2)+18)),	// W
						objekte.New(368,350,75,uint8(rand.Intn(2)+18)),objekte.New(390,425,75,uint8(rand.Intn(2)+18)),objekte.New(413,500,75,uint8(rand.Intn(2)+18)),
						objekte.New(443,575,75,uint8(rand.Intn(2)+18)),objekte.New(488,500,75,uint8(rand.Intn(2)+18)),objekte.New(525,425,75,uint8(rand.Intn(2)+18)),
						objekte.New(563,500,75,uint8(rand.Intn(2)+18)),objekte.New(608,575,75,uint8(rand.Intn(2)+18)),objekte.New(638,500,75,uint8(rand.Intn(2)+18)),
						objekte.New(660,425,75,uint8(rand.Intn(2)+18)),objekte.New(683,350,75,uint8(rand.Intn(2)+18)),objekte.New(705,275,75,uint8(rand.Intn(2)+18)),
						objekte.New(720,200,75,uint8(rand.Intn(2)+18)),objekte.New(728,125,75,uint8(rand.Intn(2)+18)),
						objekte.New(900,125,75,uint8(rand.Intn(2)+18)),objekte.New(900,200,75,uint8(rand.Intn(2)+18)),objekte.New(900,275,75,uint8(rand.Intn(2)+18)),	// B
						objekte.New(900,350,75,uint8(rand.Intn(2)+18)),objekte.New(900,425,75,uint8(rand.Intn(2)+18)),objekte.New(900,500,75,uint8(rand.Intn(2)+18)),
						objekte.New(900,575,75,uint8(rand.Intn(2)+18)),
						objekte.New(990,155,75,uint8(rand.Intn(2)+18)),objekte.New(1050,223,75,uint8(rand.Intn(2)+18)),objekte.New(990,290,75,uint8(rand.Intn(2)+18)),
						objekte.New(1020,373,75,uint8(rand.Intn(2)+18)),objekte.New(1088,440,75,uint8(rand.Intn(2)+18)),objekte.New(1065,523,75,uint8(rand.Intn(2)+18)),
						objekte.New(990,560,75,uint8(rand.Intn(2)+18)) )
	
	*akt = true
	time.Sleep( time.Duration(1.4e10) )
	
	*obj = make([]objekte.Objekt,0)					// leere den Objekte-Slice (Performance!)

	// ----------------------------------------------------------------------------------------------
	
	
	Zwischentext(&texte.MoorLvl51, mutex, stop)		// Level 5-Text - 1
	maus.SetzeTyp(16)
	Zwischentext(&texte.MoorLvl52, mutex, stop)		// Level 5-Text - 2
	maus.SetzeTyp(0)
	level5 := objekte.New(0,0,0	,11)
	Levelanzeige(level5, mutex)						// ----------------- LEVEL 5 -------------------- Kaffee und Pizza
	
	Countdown(count3,count2,count1, mutex, akt)		// lässt den Bildschirm-Countdown ablaufen
	
	jetzt := time.Now().UnixNano()
	zeit := int64(0)
	
	for zeit < 4e10 {	
		
		if i:= rand.Intn(20); i < 8 {
			*obj = append(*obj, objekte.New( uint16(rand.Intn(800))+100,uint16(rand.Intn(400))+100, uint16(rand.Intn(200)+100), uint8(rand.Intn(2)+18)) )
			*akt = true
		} else if i == 10 {
			*obj = append(*obj, objekte.New( uint16(rand.Intn(800))+200,uint16(rand.Intn(300))+200, 0, 16) )
			*hubi = true
			*akt = true
		}
		
		time.Sleep( time.Duration(2e8) )
		zeit = time.Now().UnixNano() - jetzt
	}
	
	time.Sleep( time.Duration(2e9) )
	*obj = make([]objekte.Objekt,0)					// leere den Objekte-Slice (Performance!)
	
	// ----------------------------------------------------------------------------------------------
	
	maus.SetzeTyp(16)
	time.Sleep( time.Duration(2e9) )
	
	Zwischentext(&texte.MoorScore, mutex, stop)		// Score-Text	
	
	// ----------------------------------------------------------------------------------------------
	
	maus.SetzeTyp(17)
	
	time.Sleep( time.Duration(1e9) )
	SpieleSound("./Sounds/Applaus.wav")
	
	*ende = true
	return
}

func Zwischentext(textArr *[]string, mutex *sync.Mutex, stop *bool) {
	mutex.Lock()
	LadeBild (0,0, "./Bilder/Moorhuhn/Seminarraum.bmp")
	Transparenz(120)
	Stiftfarbe(76,0,153)														
	Vollrechteck(100,50,1000,600)
	Transparenz(0)
	SetzeFont ("./Schriftarten/Freshman.ttf", 50 )
	
	Stiftfarbe(10,20,20)
	for ind,str := range *textArr {
		SchreibeFont (168, uint16(79+ind*55) ,str )
	}
	Stiftfarbe(124,212,255)
	for ind,str := range *textArr {
		SchreibeFont (170, uint16(80+ind*55) ,str )
	}
	Archivieren()
	mutex.Unlock()
	
	*stop = true
	for *stop { time.Sleep( time.Duration(1e8) ) }
}

func Levelanzeige(level objekte.Objekt, mutex *sync.Mutex) {
	mutex.Lock()
	LadeBild (0,0, "./Bilder/Moorhuhn/Seminarraum.bmp")
	level.Zeichnen()
	Archivieren()
	mutex.Unlock()
	
	time.Sleep( time.Duration(1.5e9) )
}

func Countdown(count3,count2,count1 objekte.Objekt, mutex *sync.Mutex, akt *bool) {
	mutex.Lock()
	LadeBild (0,0, "./Bilder/Moorhuhn/Seminarraum.bmp")
	count3.Zeichnen()
	Archivieren()
	mutex.Unlock()
	SpieleNote("2A",0.05,0)
	
	time.Sleep( time.Duration(1e9) )
	
	mutex.Lock()
	LadeBild (0,0, "./Bilder/Moorhuhn/Seminarraum.bmp")
	count2.Zeichnen()
	Archivieren()
	mutex.Unlock()
	SpieleNote("2A",0.05,0)
	
	time.Sleep( time.Duration(1e9) )
	
	mutex.Lock()
	LadeBild (0,0, "./Bilder/Moorhuhn/Seminarraum.bmp")
	count1.Zeichnen()
	Archivieren()
	mutex.Unlock()
	SpieleNote("2A",0.05,0)
	
	time.Sleep( time.Duration(1e9) )
	SpieleNote("3A",0.5,0)
	
	*akt = true
}


func view_komponente (obj *[]objekte.Objekt, maus,pauseObjekt,okayObjekt objekte.Objekt, pause, stop ,akt, ende *bool, punkte, diff *int16, mutex *sync.Mutex) {   	
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
			Restaurieren(0,0,1200,700)					// Restauriert das alte Hintergrundbild
		}
		
		if *stop { okayObjekt.Zeichnen() }				// zeichnet den OK-Knopf
		if *pause { pauseObjekt.Zeichnen() }			// zeichnet das Pause-Objekt
		
		maus.Zeichnen()									// zeichnet Maus
		
		SetzeFont ("./Schriftarten/Freshman.ttf", 35 )
		//Stiftfarbe(76,0,153)  
		Stiftfarbe(70,20,20)		
		SchreibeFont (455,5,"Punkte : "+fmt.Sprint (*punkte,"           Letzter Treffer : ",*diff))	// Schreibe Punkte
		Stiftfarbe(255,124,212)
		SchreibeFont (458,7,"Punkte : "+fmt.Sprint (*punkte,"           Letzter Treffer : ",*diff))	// Schreibe Punkte
		Stiftfarbe(100,10,155)
		Schreibe (1,1,"FPS:"+fmt.Sprint (anzahl))					// Schreibe links oben FPS
		
			
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
		
		if *ende { return }
		time.Sleep(time.Duration(verzögerung * 1e5)) 		// Immer ca. 100 FPS !!
		
	}
}

func ObjAktualisieren(obj *[]objekte.Objekt) {
	LadeBild (0,0, "./Bilder/Moorhuhn/Seminarraum.bmp")		// Hintergrund des Moorhuhn-Raumes wird gezeichnet
	
	for _,ob := range *obj { 								// Zeichnet alleweiteren Objekte ein
		ob.Zeichnen()
	}
	Archivieren()											// Speichert das Hintergrund-Bild
}

// Es folgt die CONTROL-Komponente 1 --- Kein Bestandteil der Welt, also unabhängig -----
func maussteuerung (obj *[]objekte.Objekt, maus,pauseObjekt,okayObjekt objekte.Objekt, pause,stop,hubi,akt,ende *bool, punkte, diff *int16) {
	/*var taste uint8
	var status int8 */
	for {
		taste, status, mausX, mausY := MausLesen1()
		
		maus.SetzeKoordinaten(mausX,mausY)					// Aktualisiert Maus-Koordinaten
		
		if *ende {
			return
		} else if *pause {		
			if taste==1 && status==1 { 						//LINKE Maustaste gerade gedrückt
				if ja,_ := pauseObjekt.Getroffen(mausX,mausY,1); ja {
					*ende = true
					return
				}
			}
		} else if *stop {
			if taste==1 && status==1 { 						//LINKE Maustaste gerade gedrückt
				if ja,_ := okayObjekt.Getroffen(mausX,mausY,1); ja {
					*stop = false
				}
			}
		} else if *hubi {									// falls ein Hubi (Lvl 5) aktiv ist
			if status==1 {
				SpieleSound("./Sounds/Baeb.wav")
				*punkte -= 5
			}
		} else {
			if taste==1 && status==1 { 						//LINKE Maustaste gerade gedrückt
				for _,ob := range *obj { 							// Zeichnet alleweiteren Objekte ein
					if get,lang := ob.Getroffen(mausX,mausY,1); get {
						if lang == 0 {
							*punkte -= 5
							*diff = -5
							 fmt.Println("Knapp daneben - 5 MINUS-Punkte")
						} else if lang < 3.5e8 {
							 *punkte += 20
							 *diff = 20
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 20 Punkte")
						 } else if lang < 4.2e8 {
							 *punkte += 15
							 *diff = 15
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 15 Punkte")
						 } else if lang < 5.2e8 {
							 *punkte += 10
							 *diff = 10
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 10 Punkte")
						 } else if lang < 7e8 {
							 *punkte += 5
							 *diff = 5
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 5 Punkte")
						 } else if lang < 1e9 {
							 *punkte += 2
							 *diff = 2
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 2 Punkte")
						 } else {
							 *punkte += 1
							 *diff = 1
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
							*diff = -5
							 fmt.Println("Knapp daneben - 5 MINUS-Punkte")
						} else if lang < 3.5e8 {
							 *punkte += 20
							 *diff = 20
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 20 Punkte")
						} else if lang < 4.2e8 {
							 *punkte += 15
							 *diff = 15
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 15 Punkte")
						} else if lang < 5.2e8 {
							 *punkte += 10
							 *diff = 10
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 10 Punkte")
						} else if lang < 7e8 {
							 *punkte += 5
							 *diff = 5
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 5 Punkte")
						} else if lang < 1e9 {
							 *punkte += 2
							 *diff = 2
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 2 Punkte")
						} else {
							 *punkte += 1
							 *diff = 1
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 1 Punkt")
						}		
						*akt = true
					}
				}
			}
		}
	}
	fmt.Println("Beende Maussteuerung!")
}

