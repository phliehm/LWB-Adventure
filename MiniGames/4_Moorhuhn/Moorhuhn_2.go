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
			"math/rand"
			)

const breite = 800 		// von Gott vorgegeben
const hoehe  = 600  	// von Gott vorgegeben
	
func main () {
	var mutex sync.Mutex
	var punkte uint16 = 10					// Spiel-Punktzahl
	var pause bool = false
	var akt	bool = true						// Prüft, ob Grafik aktualisiert werden muss
	obj := make([]objekte.Objekt,0)			// Array für die Objekte der Welt
	
	
	fmt.Println(time.Now().UnixNano() )
	random := rand.New( rand.NewSource( time.Now().UnixNano() ) )
	
	
    
	pauseObjekt := objekte.New(breite,hoehe,hoehe/4	,1)			// Erstellt das Objekt PAUSE 
	maus 		:= objekte.New(0,0,			30		,0)			// Erstellt das Objekt MAUSZEIGER mit Größe (40)

	
	Fenster (breite,hoehe)									// Öffnet GFXW-Fenster
	Fenstertitel("StEPS-Wars")								// Gibt Fenster-Titel 
	// SetzeFont ("collegeb.ttf", hoehe/4)					// Setzt Schriftart
	// SetzeFont ("Prisma.ttf", hoehe/4)
	SetzeFont ("../../Schriftarten/Freshman.ttf", hoehe/20 ) 
	
	
	// Das Hauptprogramm startet die View-Komponente als nebenläufigen Prozess!
	go view_komponente(&obj, maus, pauseObjekt, &pause, &akt, &punkte, &mutex)
	
	// Objekte werden nach und nach in der Welt platziert
	go erstelleObjekte(&obj, &pause, &akt, random, &mutex)
	
	// Nebenläufig wird die Kontroll-Komponente 1 für die Maus gestartet.
	go maussteuerung (&obj, maus, &pause, &akt, &punkte)
	
	
	// Die Kontroll-Komponente 2 ist die 'Mainloop' im Hauptprogramm	
	// Wir fragen hier nur die Tastatur ab.
	// In anderen Sprachen würde man hier alle Events abfragen, also
	// ob es eine Tasteneingabe oder Mauseingabe war. Dann gäbe es
	// keine Control-Komponente 1.
	
A:	for {
		taste,gedrueckt,_:= TastaturLesen1 ()
		if gedrueckt == 1 {
			switch taste {
				case 'q': 													// mit 'q' wird das Programm beendet!
				break A
				case 'p': 
				pause = !pause												// Pause-Modus !!
				case ' ':
				case 'y':
				case 'x':
				
			}
		}
	}
	

}

func erstelleObjekte(obj *[]objekte.Objekt, pause,akt *bool, rand *rand.Rand, mutex *sync.Mutex) {		// füllt Objekte ins Array
	level1 := objekte.New(breite,hoehe,hoehe/4	,7)
	level2 := objekte.New(breite,hoehe,hoehe/4	,8)
	level3 := objekte.New(breite,hoehe,hoehe/4	,9)
	level4 := objekte.New(breite,hoehe,hoehe/4	,10)
	
	time.Sleep( time.Duration(2e9) )
	
	
	// ----------------- LEVEL 1 --------------------
	mutex.Lock()
	raeume.Moorhuhn(breite)
	level1.Zeichnen()
	Archivieren()
	mutex.Unlock()
	
	time.Sleep( time.Duration(2e9) )
	*akt = true
	
	for i:=0;i<15;i++ {
		time.Sleep( time.Duration(	rand.Uint32()/2 ) )
		*obj = append(*obj, objekte.New( uint16(rand.Intn(700)),uint16(rand.Intn(500)),		uint16(rand.Intn(150)+150)	, 11) )
		*akt = true
	}
	
	// leere den Objekte-Slice (Performance!)
	*obj = make([]objekte.Objekt,0)
	fmt.Println(*obj)
	
	time.Sleep( time.Duration(2e9) )
	
	// ----------------- LEVEL 2 --------------------
	mutex.Lock()
	raeume.Moorhuhn(breite)
	level2.Zeichnen()
	Archivieren()
	mutex.Unlock()
	
	time.Sleep( time.Duration(2e9) )
	*akt = true
	
	for i:=0;i<15;i++ {
		time.Sleep( time.Duration(	rand.Uint32()/2 ) )
		*obj = append(*obj, objekte.New( uint16(rand.Intn(700)),uint16(rand.Intn(500)),		uint16(rand.Intn(100)+30)	, 11) )
		*akt = true
	}
	
	// leere den Objekte-Slice
	*obj = make([]objekte.Objekt,0)
	fmt.Println(*obj)
	
	time.Sleep( time.Duration(2e9) )
	
	// ----------------- LEVEL 3 --------------------
	mutex.Lock()
	raeume.Moorhuhn(breite)
	level3.Zeichnen()
	Archivieren()
	mutex.Unlock()
	
	time.Sleep( time.Duration(2e9) )
	*akt = true
	
	for i:=0;i<20;i++ {
		*obj = append(*obj, objekte.New( uint16(rand.Intn(700)),uint16(rand.Intn(500)),		uint16(rand.Intn(250)+30)	,uint8(rand.Intn(2)*2+3)) )
		*akt = true
		time.Sleep( time.Duration(	rand.Uint32()/2 ) )
	}
	
	// ----------------- LEVEL 4 --------------------
	mutex.Lock()
	raeume.Moorhuhn(breite)
	level4.Zeichnen()
	Archivieren()
	mutex.Unlock()
	
	time.Sleep( time.Duration(2e9) )
	*akt = true
	
	
	*obj = append(*obj, objekte.New(0,150,		350	,3) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(200,0,		300	,5) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(400,300,	250	,3) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(600,200,	200	,5) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(400,100,	150	,3) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(650,80,	100	,5) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(600,50,	50	,3) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(500,40,	50	,5) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	
}

// Es folgt die VIEW-Komponente --- Kein Bestandteil der Welt, also unabhängig---------
// Einzig diese Funktion verwendet gfxw-Befehle!!
func view_komponente (obj *[]objekte.Objekt, maus,pauseObjekt objekte.Objekt, pause,akt *bool, punkte *uint16, mutex *sync.Mutex) {   	
	var t1 int64 = time.Now().UnixNano() 		//Startzeit
	var anz,anzahl int                  		// zur Bestimmung der Frames pro Sekunde
	var verzögerung = 90
	
	for { //Endlos ...
		mutex.Lock()
		UpdateAus () 							// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
		
		Stiftfarbe(255,255,255)
		Cls()									// Cleart vollständigen Screen
		
		if *akt { 
			ObjAktualisieren(obj)
			*akt = false
		} else {
			Restaurieren(0,0,breite,hoehe)		// Restauriert das alte Hintergrundbild
		}
	
		
		maus.Zeichnen()								// Zeichnet Maus
		//mutex.Unlock()
		
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
			if anzahl < 100 { verzögerung--}		//Selbstregulierung der 
			if anzahl > 100 { verzögerung++}		//Frame-Rate :-)
		}
		
		//mutex.Lock()
		UpdateAn () // Nun wird der gezeichnete Frame sichtbar gemacht!
		mutex.Unlock()
		
		time.Sleep(time.Duration(verzögerung * 1e5)) // Immer ca. 100 FPS !!
		
	}
}

func ObjAktualisieren(obj *[]objekte.Objekt) {
	raeume.Moorhuhn(breite)						// Hintergrund des Moorhuhn-Raumes wird gezeichnet
	// raeume.Hauptflur(breite)
	
	for _,ob := range *obj { 					// Zeichnet alleweiteren Objekte ein
		ob.Zeichnen()
	}
	Archivieren()							// Speichert das Hintergrund-Bild
}

// Es folgt die CONTROL-Komponente 1 --- Kein Bestandteil der Welt, also unabhängig -----
func maussteuerung (obj *[]objekte.Objekt, maus objekte.Objekt, pause,akt *bool, punkte *uint16) {
	/*var taste uint8
	var status int8 */
	for {
		taste, status, mausX, mausY := MausLesen1()
		
		maus.SetzeKoordinaten(mausX,mausY)				// Aktualisiert Maus-Koordinaten
		
		if *pause==false && taste==1 && status==1 { 				//LINKE Maustaste gerade gedrückt
			for _,ob := range *obj { 							// Zeichnet alleweiteren Objekte ein
				if get,lang := ob.Getroffen(mausX,mausY,1); get {
					if lang == 0 {
						*punkte -= 5
						 fmt.Println(lang, " 5 MINUS-Punkte")
					} else if lang < 2.5e8 {
						 *punkte += 20
						 fmt.Println(lang, " 20 Punkte")
					 } else if lang < 4.2e8 {
						 *punkte += 15
						 fmt.Println(lang, " 15 Punkte")
					 } else if lang < 5.2e8 {
						 *punkte += 10
						 fmt.Println(lang, " 10 Punkte")
					 } else if lang < 7e8 {
						 *punkte += 5
						 fmt.Println(lang, " 5 Punkte")
					 } else if lang < 1e9 {
						 *punkte += 2
						 fmt.Println(lang, " 2 Punkte")
					 } else {
						 *punkte += 1
						 fmt.Println(lang, " 1 Punkt")
					}	
					*akt = true
				}
			}
		}
		if *pause==false && taste == 3 && status == 1 { 			//RECHTE Maustaste gerade gedrückt
			for _,ob := range *obj { 							// Zeichnet alleweiteren Objekte ein
				if get,lang :=  ob.Getroffen(mausX,mausY,3); get {
					if lang == 0 {
						*punkte -= 5
						 fmt.Println(lang, " 5 MINUS-Punkte")
					} else if lang < 2.5e8 {
						 *punkte += 20
						 fmt.Println(lang, " 20 Punkte")
					 } else if lang < 4.2e8 {
						 *punkte += 15
						 fmt.Println(lang, " 15 Punkte")
					 } else if lang < 5.2e8 {
						 *punkte += 10
						 fmt.Println(lang, " 10 Punkte")
					 } else if lang < 7e8 {
						 *punkte += 5
						 fmt.Println(lang, " 5 Punkte")
					 } else if lang < 1e9 {
						 *punkte += 2
						 fmt.Println(lang, " 2 Punkte")
					 } else {
						 *punkte += 1
						 fmt.Println(lang, " 1 Punkt")
					}	
					*akt = true
				}
			}
			//fmt.Println ("rechts: ", taste, status,mausX,mausY)	// printet Koordinaten des Klicks
		}
	}
}

